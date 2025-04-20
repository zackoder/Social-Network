package models

import "fmt"

func deletfollow(follower, followed string) error {
	fmt.Println("hello")
	deleteQuery := "DELETE FROM followers WHERE follower_id = ? AND followed_id = ?"
	_, err := Db.Exec(deleteQuery, follower, followed)
	return err
}
