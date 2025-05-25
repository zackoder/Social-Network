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
	insetpostQuery := "INSERT INTO posts (post_privacy, title, content, user_id, imagePath, createdAt,groupe_id) VALUES (?,?,?,?,?,strftime('%s', 'now'),?)"
	res, err := Db.Exec(insetpostQuery, post.Privacy, post.Title, post.Content, post.Poster_id, post.Image, post.Groupe_id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	lastId, _ := res.LastInsertId()
	return int(lastId), nil
}

func InsertFriends(id int, friendes []int) {
	insertFriend := "INSERT INTO friends (post_id, friend_id) VALUES(?,?)"
	for _, friend := range friendes {
		Db.Exec(insertFriend, id, friend)
	}
}

func InsertFollow(follower, followed string) error {
	inserQuery := "INSERT INTO followers (follower_id, followed_id) VALUES (?,?)"
	_, err := Db.Exec(inserQuery, follower, followed)
	return err
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

func InsertGroupMSG(msg utils.Message) error {
	query := "INSERT INTO groups_chat (group_id, sender_id, content, imagePath) VALUES (?,?,?,?)"
	_, err := Db.Exec(query, msg.Group_id, msg.Sender_id, msg.Content, msg.Filename)
	if err != nil {
		fmt.Println("inserting error:", err)
	} else {
		fmt.Println("the message was inserted")
	}
	return err
}

func InsertNotification(noti utils.Notification) error {
	query := "INSERT INTO notifications (user_id, target_id, actor_id, message) VALUES (?,?,?,?)"
	_, err := Db.Exec(query, noti.Sender_id, noti.Target_id, noti.Actor_id, noti.Message)
	return err
}

func SaveInvitation(Groupe_id, sender_id, resever_id int) error {
	Quirie := "INSER INTO invitation (sender_id,recever_id,groupe_id) VALUES(?,?)"
	_, err := Db.Exec(Quirie, sender_id, resever_id, Groupe_id)
	return err
}

func InsserGroupe(title, description string, creator_id int) (int,error) {
	Query := "INSERT INTO groups (name, description, group_oner) VALUES (?,?,?)"
	res, err := Db.Exec(Query, title, description, creator_id)
	id, err := res.LastInsertId()
	fmt.Println("errggggggggggggggggggggggggggg",err)
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func InsserMemmberInGroupe(Groupe_id, User_id int,role string) error {
	Quirie := "INSERT INTO group_members (group_id,user_id,role) VALUES (?,?,?)"
	_, err := Db.Exec(Quirie, Groupe_id, User_id,role)
	fmt.Println("errrrrrrrrrrrrr",err)
	return err
}

func InsserEventInDatabase(event utils.Event) error {
	Quirie := "INSERT INTO events (group_id,title,description,event_time,created_by) VALUES (?,?,?,?)"
	_, err := Db.Exec(Quirie, event.GroupID, event.Title, event.Description, event.EventTime, event.CreatedBy)
	return err
}

func InsserResponceInDatabase(responce utils.EventResponse) error {
	Quirie := "INSERT INTO event_responses (user_id,event_id,response) VALUES (?,?,?)"
	_, err := Db.Exec(Quirie, responce.UserID, responce.EventID, responce.Response)
	return err
}
