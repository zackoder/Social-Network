package models

import (
	"strconv"

	"social-network/utils"
)

func UpdateProfile(id int) string {
	privacy, _ := IsPrivateProfile(strconv.Itoa(id))
	updateprofile := "UPDATE users SET privacy = ? WHERE id = ?"
	if !privacy {
		Db.Exec(updateprofile, "privet", id)
		return "privet"
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
