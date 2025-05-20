package models

import (
	"database/sql"
	"fmt"
	"social-network/utils"
)

// GetUserById retrieves a user by their ID
func GetUserById(userId int) (*utils.User, error) {
	query := `
		SELECT id, first_name, last_name, nickname, email, avatar, AboutMe, privacy 
		FROM users 
		WHERE id = ?
	`

	var user utils.User
	err := Db.QueryRow(query, userId).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Nickname,
		&user.Email,
		&user.Avatar,
		&user.AboutMe,
		&user.Privacy,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}

	return &user, nil
}

// AreUsersFriends checks if two users are friends (both follow each other)
func AreUsersFriends(userA int, userB int) (bool, error) {
	// Two users are considered friends if they follow each other
	isAFollowingB, err := IsUserFollowing(userA, userB)
	if err != nil {
		return false, err
	}

	isBFollowingA, err := IsUserFollowing(userB, userA)
	if err != nil {
		return false, err
	}

	// Both users must be following each other to be considered friends
	return isAFollowingB && isBFollowingA, nil
}

// IsProfilePrivate checks if a user's profile is set to private
func IsProfilePrivate(userId int) (bool, error) {
	var privacy string
	query := `SELECT privacy FROM users WHERE id = ?`
	err := Db.QueryRow(query, userId).Scan(&privacy)
	if err != nil {
		return false, err
	}

	return privacy == "private", nil
}

// IsUserFollowing checks if a user is following another user
func IsUserFollowing(followerId int, followedId int) (bool, error) {
	var count int
	query := `SELECT COUNT(*) FROM followers WHERE follower_id = ? AND followed_id = ?`
	err := Db.QueryRow(query, followerId, followedId).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
