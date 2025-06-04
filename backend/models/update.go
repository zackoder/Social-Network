package models

import (
	"strconv"

	"social-network/utils"
)

func UpdateProfile(id int) string {
	privacy, _ := IsPrivateProfile(strconv.Itoa(id))
	updateprofile := "UPDATE users SET privacy = ? WHERE id = ?"
	if !privacy {
		Db.Exec(updateprofile, "private", id)
		return "private"
	} else if privacy {
		Db.Exec(updateprofile, "public", id)
		return "public"
	}
	return ""
}

func UpdateNoti(oldNoti, noti utils.Notification) error {
	updatenoti := `
		UPDATE notifications 
		SET user_id = ?, actor_id = ? AND target_id = ?
		WHERE id = ?;
	`
	_, err := Db.Exec(updatenoti, noti.Sender_id, noti.Actor_id, noti.Target_id, oldNoti.Id)
	return err
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
