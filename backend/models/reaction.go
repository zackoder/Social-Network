package models

import (
	"fmt"
	"social-network/utils"
)

// InsertReaction adds a new reaction to the database
func InsertReaction(reaction *utils.Reaction) (int, error) {
	query := `
		INSERT INTO reactions (user_id, post_id, reaction_type, date) 
		VALUES (?, ?, ?, strftime('%s', 'now'))
	`

	result, err := Db.Exec(query, reaction.UserId, reaction.PostId, reaction.ReactionType)
	if err != nil {
		fmt.Println("Error inserting reaction:", err)
		return 0, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(lastId), nil
}

// GetReactionsByPostId retrieves all reactions for a specific post
func GetReactionsByPostId(postId int) ([]utils.Reaction, error) {
	query := `
		SELECT r.id, r.post_id, r.user_id, r.reaction_type, 
		       strftime('%s', r.date) as unix_date
		FROM reactions r
		WHERE r.post_id = ?
	`

	rows, err := Db.Query(query, postId)
	if err != nil {
		fmt.Println("Error querying reactions:", err)
		return nil, err
	}
	defer rows.Close()

	var reactions []utils.Reaction
	for rows.Next() {
		var reaction utils.Reaction
		err := rows.Scan(
			&reaction.Id,
			&reaction.PostId,
			&reaction.UserId,
			&reaction.ReactionType,
			&reaction.Date,
		)
		if err != nil {
			fmt.Println("Error scanning reaction row:", err)
			continue
		}
		reactions = append(reactions, reaction)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return reactions, nil
}

// GetUserReactionForPost gets a specific user's reaction to a post, if any
func GetUserReactionForPost(userId, postId int) (*utils.Reaction, error) {
	query := `
		SELECT id, post_id, user_id, reaction_type, 
		       strftime('%s', date) as unix_date
		FROM reactions
		WHERE user_id = ? AND post_id = ?
	`

	var reaction utils.Reaction
	err := Db.QueryRow(query, userId, postId).Scan(
		&reaction.Id,
		&reaction.PostId,
		&reaction.UserId,
		&reaction.ReactionType,
		&reaction.Date,
	)

	if err != nil {
		return nil, err
	}

	return &reaction, nil
}

// UpdateReaction changes the type of a user's reaction on a post
func UpdateReaction(userId, postId int, newReactionType string) error {
	query := `
		UPDATE reactions 
		SET reaction_type = ?, date = strftime('%s', 'now')
		WHERE user_id = ? AND post_id = ?
	`

	_, err := Db.Exec(query, newReactionType, userId, postId)
	return err
}

// DeleteReaction removes a user's reaction from a post
func DeleteReaction(userId, postId int) error {
	query := "DELETE FROM reactions WHERE user_id = ? AND post_id = ?"
	_, err := Db.Exec(query, userId, postId)
	return err
}

// CountReactionsByType counts reactions by type for a given list of reactions
func CountReactionsByType(reactions []utils.Reaction) map[string]int {
	counts := make(map[string]int)

	for _, reaction := range reactions {
		counts[reaction.ReactionType]++
	}

	return counts
}
