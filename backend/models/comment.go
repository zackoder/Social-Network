package models

import (
	"fmt"
	"social-network/utils"
)

// InsertComment adds a new comment to the database
func InsertComment(comment *utils.Comment) (int, error) {
	query := `
		INSERT INTO comments (user_id, post_id, comment, imagePath, date) 
		VALUES (?, ?, ?, ?, strftime('%s', 'now'))
	`

	result, err := Db.Exec(query, comment.UserId, comment.PostId, comment.Content, comment.ImagePath)
	if err != nil {
		fmt.Println("Error inserting comment:", err)
		return 0, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(lastId), nil
}

// GetCommentsByPostId retrieves all comments for a specific post
func GetCommentsByPostId(postId int) ([]utils.Comment, error) {
	query := `
		SELECT c.id, c.post_id, c.user_id, c.comment, c.imagePath, c.date,
		       u.first_name || ' ' || u.last_name as user_name, u.avatar
		FROM comments c
		JOIN users u ON c.user_id = u.id
		WHERE c.post_id = ?
		ORDER BY c.date DESC
	`

	rows, err := Db.Query(query, postId)
	if err != nil {
		fmt.Println("Error querying comments:", err)
		return nil, err
	}
	defer rows.Close()

	var comments []utils.Comment
	for rows.Next() {
		var comment utils.Comment
		err := rows.Scan(
			&comment.Id,
			&comment.PostId,
			&comment.UserId,
			&comment.Content,
			&comment.ImagePath,
			&comment.Date,
			&comment.UserName,
			&comment.UserAvatar,
		)
		if err != nil {
			fmt.Println("Error scanning comment row:", err)
			continue
		}
		comments = append(comments, comment)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}

// DeleteComment removes a comment from the database
func DeleteComment(commentId int) error {
	query := "DELETE FROM comments WHERE id = ?"
	_, err := Db.Exec(query, commentId)
	return err
}

// CanModifyComment checks if a user has permission to modify or delete a comment
// func CanModifyComment(userId, commentId int) (bool, error) {
// 	// User can modify their own comments or if they are the owner of the post
// 	query := `
// 		SELECT EXISTS(
// 			SELECT 1 FROM comments 
// 			WHERE id = ? AND user_id = ?
// 			UNION
// 			SELECT 1 FROM comments c
// 			JOIN posts p ON c.post_id = p.id
// 			WHERE c.id = ? AND p.user_id = ?
// 		) as can_modify
// 	`

// 	var canModify bool
// 	err := Db.QueryRow(query, commentId, userId, commentId, userId).Scan(&canModify)
// 	return canModify, err
// }

// // PostExists checks if a post exists in the database
// func PostExists(postId int) (bool, error) {
// 	query := "SELECT EXISTS(SELECT 1 FROM posts WHERE id = ?)"
// 	var exists bool
// 	err := Db.QueryRow(query, postId).Scan(&exists)
// 	return exists, err
// }
