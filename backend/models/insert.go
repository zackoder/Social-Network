package models

import (
	"fmt"
	"social-network/utils"
)

func InsertUser(user utils.Regester) error {
	insertuserquery := "INSERT INTO users (first_name, last_name, nickname, email, age, gender, password, avatar, AboutMe)  VALUES(?,?,?,?,?,?,?,?,?)"
	if _, err := Db.Exec(insertuserquery, user.FirstName, user.LastName, user.NickName, user.Email, user.Age, user.Gender, user.Password, user.Avatar, user.About_Me); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func InsertPost(post utils.Post) error {
	insetpostQuery := "INSERT INTO posts (post_privacy, title, content, user_id, imagePath, createdAt) VALUES (?,?,?,?,?,strftime('%s', 'now'))"
	res, err := Db.Exec(insetpostQuery, post.Privacy, post.Title, post.Content, 1, post.Image)
	if err != nil {
		return err
	}
	lastId, _ := res.LastInsertId()
	post.Id = int(lastId)
	return nil
}
