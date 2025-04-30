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

func InsertPost(post utils.Post) (int, error) {
	insetpostQuery := "INSERT INTO posts (post_privacy, title, content, user_id, imagePath, createdAt) VALUES (?,?,?,?,?,strftime('%s', 'now'))"
	res, err := Db.Exec(insetpostQuery, post.Privacy, post.Title, post.Content, 1, post.Image)
	if err != nil {
		return 0, err
	}
	lastId, _ := res.LastInsertId()
	return int(lastId), nil
}

func InsertFriends(id int, friendes []string) {
	insertFriend := "INSERT INTO friends (post_id, friend_id) VALUES(?,?)"
	for _, friend := range friendes {
		Db.Exec(insertFriend, id, friend)
	}
}

func InserOrUpdate(follower, followed string) (string, error) {
	privacy, err := IsPrivateProfile(followed)
	if err != nil {
		return "", err
	}
	if privacy {

		if err := insertFollow(follower, followed); err != nil {
			if err := deletfollow(follower, followed); err != nil {
				fmt.Println(err)
				return "", err
			}
			fmt.Println(err)
			return "unfollow seccessfully", nil
		}
		return "following seccessfully", nil
	}
	InsertFollowreq(followed)
	return "follow request sent", nil
}

func insertFollow(follower, followed string) error {
	inserQuery := "INSERT INTO followers (follower_id, followed_id) VALUES (?,?)"
	_, err := Db.Exec(inserQuery, follower, followed)
	return err
}

func InsertFollowreq(followed string) {

}

func InsertNewGroup(group *utils.NewGroup, user_id int) error {
	insertGroup := "INSERT INTO groups (name, description, group_oner) VALUES (?,?,?)"
	res, err := Db.Exec(insertGroup, group.Title, group.Content, user_id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	lastId, _ := res.LastInsertId()
	group.Id = int(lastId)
	return nil
}

func InsertMumber(group_id, user_id int) error {
	insertMumber := "INSERT INTO group_members (group_id, user_id) VALUES (?,?)"
	if _, err := Db.Exec(insertMumber, group_id, user_id); err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("user added")
	return nil
}

func InsertSession(userData *utils.User) error {
	_, err := Db.Exec("INSERT INTO sessions ( user_id, token) VALUES (?, ?)", userData.ID, userData.SessionId)
	return err
}

func InsertMsg(msg utils.Message) error {
	query := "INSERT INTO messages (sender_id, reciever_id, content, imagePath) VALUES (?,?,?,?)"
	_, err := Db.Exec(query, msg.Sender_id, msg.Reciever_id, msg.Content, msg.Filename)
	if err != nil {
		fmt.Println("inserting error:", err)
	} else {
		fmt.Println("the message was inserted")
	}
	return err
}
