package models

import (
	"net/http"

	"social-network/utils"
)

func Deletfollow(follower int, followed string) error {
	deleteQuery := "DELETE FROM followers WHERE follower_id = ? AND followed_id = ?"
	_, err := Db.Exec(deleteQuery, follower, followed)
	return err
}

func DeleteSession(userId int) error {
	query := `DELETE FROM sessions WHERE user_id =  ?;`
	_, err := Db.Exec(query, userId)
	return err
}

func RemoveSessionFromDB(token *http.Cookie) error {
	query := "DELETE  FROM sessions WHERE token = ?"
	_, err := Db.Exec(query, token)
	return err
}

func DeleteNoti(noti_id int) error {
	deleteNoti := "DELETE FROM notifications WHERE id = ?"
	_, err := Db.Exec(deleteNoti, noti_id)
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
