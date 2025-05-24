package models

import "fmt"

// PostExists checks if a post exists in the database
func PostExists(postId int) (bool, error) {
	query := "SELECT EXISTS(SELECT 1 FROM posts WHERE id = ?)"
	var exists bool
	err := Db.QueryRow(query, postId).Scan(&exists)
	if err != nil {
		fmt.Println("Error checking if post exists:", err)
		return false, err
	}
	return exists, nil
}
