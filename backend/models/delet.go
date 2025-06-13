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

func RemoveNoti(res utils.EventResponse) {
	deleteNoti := `
		DELET FROM notifications WHERE 
	`
	_ = deleteNoti
}
