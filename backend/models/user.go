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
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}

	return &user, nil
}
