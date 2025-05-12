package models

import (
	"database/sql"
	"fmt"
)

// CanUserAccessPost checks if a user can access a post based on its privacy settings
func CanUserAccessPost(userId int, postId int) (bool, error) {
	// Query to get the post's privacy and poster id
	var privacy string
	var posterId int
	query := `SELECT privacy, poster_id FROM posts WHERE id = ?`
	err := Db.QueryRow(query, postId).Scan(&privacy, &posterId)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, fmt.Errorf("Post not found")
		}
		return false, err
	}

	// If user is the post owner, they can always access it
	if userId == posterId {
		return true, nil
	}

	// Check privacy settings
	switch privacy {
	case "public":
		// Anyone can access public posts
		return true, nil

	case "almostPrivet":
		// Check if the user is a follower of the post creator
		isFollower, err := IsUserFollowing(userId, posterId)
		if err != nil {
			return false, err
		}
		return isFollower, nil

	case "private":
		// Check if the user is specifically allowed to see this post
		// This would require checking a table that stores which friends can see private posts
		// For now, we'll check if they're mentioned in the friends list for this post
		var count int
		query = `SELECT COUNT(*) FROM post_friends WHERE post_id = ? AND user_id = ?`
		err := Db.QueryRow(query, postId, userId).Scan(&count)
		if err != nil {
			return false, err
		}
		return count > 0, nil

	default:
		// Unknown privacy setting, default to no access
		return false, nil
	}
}

// IsUserFollowing checks if userA is following userB
func IsUserFollowing(followerID, followedID int) (bool, error) {
	var count int
	query := `SELECT COUNT(*) FROM followers WHERE follower_id = ? AND followed_id = ?`
	err := Db.QueryRow(query, followerID, followedID).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
