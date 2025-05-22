package models

func Deletfollow(follower, followed string) error {
	deleteQuery := "DELETE FROM followers WHERE follower_id = ? AND followed_id = ?"
	_, err := Db.Exec(deleteQuery, follower, followed)
	return err
}

func DeleteSession(userId int) error {
	query := `DELETE FROM sessions WHERE user_id =  ?;`
	_, err := Db.Exec(query, userId)
	return err
}
