package models

import (
	"net/http"

	"social-network/utils"
)

func Deletfollow(follower, followed string) error {
	deleteQuery := "DELETE FROM followers WHERE follower_id = ? AND followed_id = ?"
	_, err := Db.Exec(deleteQuery, follower, followed)
	return err
}

func DeleteSession(userData *utils.User) error {
	query := `DELETE FROM sessions WHERE user_id =  ?;`
	_, err := Db.Exec(query, userData.ID)
	return err
}

func RemoveSessionFromDB(token *http.Cookie) error {
	query := "DELETE  FROM sessions WHERE token = ?"
	_, err := Db.Exec(query, token)
	return err
}
