package models

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"social-network/utils"
)

//	func InsertUser(user utils.User) (int, error ){
//		insertuserquery := "INSERT INTO users (first_name, last_name, nickname, email, age, gender, password, avatar, AboutMe)  VALUES(?,?,?,?,?,?,?,?,?)"
//		 result,err := Db.Exec(insertuserquery, user.FirstName, user.LastName, user.Nickname, user.Email, user.Age, user.Gender, user.Password, user.Avatar, user.AboutMe);
//		 if err != nil{
//			 fmt.Println(err)
//			 return 0, err
//			}
//			user.ID,_ = result.LastInsertId()
//		return  int(user.ID),nil
//	}
func RegisterUser(user *utils.User) error {
	insertuserquery := "INSERT INTO users (first_name, last_name, nickname, email, age, gender, password, avatar, AboutMe)  VALUES(?,?,?,?,?,?,?,?,?)"
	result, err := Db.Exec(insertuserquery, user.FirstName, user.LastName, user.Nickname, user.Email, user.Age, user.Gender, user.Password, user.Avatar, user.AboutMe)
	if err != nil {
		return err
	}
	user.ID, err = result.LastInsertId()
	return err
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

// func InserOrUpdate(follower, followed string) (string, error) {
// 	privacy, err := IsPrivateProfile(followed)
// 	if err != nil {
// 		return "", err
// 	}

// 	if !privacy {
// 		if err := InsertFollow(follower, followed); err != nil {
// 			if err := Deletfollow(follower, followed); err != nil {
// 				fmt.Println(err)
// 				return "", err
// 			}
// 			fmt.Println(err)
// 			return "unfollow seccessfully", nil
// 		}
// 		return "following seccessfully", nil
// 	}
// 	InsertFollowreq(followed)
// 	return "follow request sent", nil
// }

// func InserOrUpdate(follower, followed string) (string, error) {
// 	privacy, err := IsPrivateProfile(followed)
// 	if err != nil {
// 		return "", err
// 	}
// 	if !privacy {

// 		if err := InsertFollow(follower, followed); err != nil {
// 			if err := Deletfollow(follower, followed); err != nil {
// 				fmt.Println(err)
// 				return "", err
// 			}
// 			fmt.Println(err)
// 			return "unfollow seccessfully", nil
// 		}
// 		return "following seccessfully", nil
// 	}
// 	InsertFollowreq(followed)
// 	return "follow request sent", nil
// }

func InsertFollow(follower int, followed string) error {
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
	fmt.Println(err)
	return err
}

func InsertMsg(msg utils.Message) error {
	query := "INSERT INTO messages (sender_id, reciever_id, content, imagePath, creation_date) VALUES (?,?,?,?,?)"
	_, err := Db.Exec(query, msg.Sender_id, msg.Reciever_id, msg.Content, msg.Filename, time.Now().Unix())
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

func InsertNotification(noti utils.Notification) (int, error) {
	var oldNoti utils.Notification
	var noti_id int
	SelectOneNoti(&oldNoti)
	var err error
	if oldNoti.Id != 0 {
		err = UpdateNoti(oldNoti, noti)
	} else {
		query := "INSERT INTO notifications (user_id, target_id, actor_id, message) VALUES (?,?,?,?)"
		res, err := Db.Exec(query, noti.Sender_id, noti.Target_id, noti.Actor_id, noti.Message)
		if err != nil {
			log.Println("inserting notin error: ", err)
		}
		lastId, _ := res.LastInsertId()
		noti_id = int(lastId)
	}

	return noti_id, err
}

func SaveInvitation(Groupe_id, sender_id, resever_id int) error {
	Quirie := "INSERT INTO invitation (sender_id,recever_id,groupe_id) VALUES(?,?,?)"
	_, err := Db.Exec(Quirie, sender_id, resever_id, Groupe_id)
	return err
}

func InsertGroupe(title, description string, creator_id int) (int, error) {
	query := "INSERT INTO groups (name, description, group_oner) VALUES (?, ?, ?)"
	res, err := Db.Exec(query, title, description, creator_id)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func InsserMemmberInGroupe(Groupe_id, User_id int, role string) error {
	Quirie := "INSERT INTO group_members (group_id,user_id,role) VALUES (?,?,?)"
	_, err := Db.Exec(Quirie, Groupe_id, User_id, role)
	return err
}

func InsserEventInDatabase(event utils.Event) (int, error) {
	Quirie := "INSERT INTO events (group_id,title,description,event_time,created_by) VALUES (?,?,?,?,?)"
	res, err := Db.Exec(Quirie, event.GroupID, event.Title, event.Description, event.EventTime, event.CreatedBy)
	if err != nil {
		log.Println("inserting event err", err)
		return 0, err
	}
	lastid, _ := res.LastInsertId()
	return int(lastid), err
}

func InsserResponceInDatabase(responce utils.EventResponse) error {
	Quirie := "INSERT INTO event_responses (user_id,event_id,response) VALUES (?,?,?)"

	_, err := Db.Exec(Quirie, responce.UserID, responce.EventID, responce.Response)
	fmt.Println(err)
	return err
}

// this function is not used yet it supposed to insert people who can see the private posts into DB.
func AddPrivateViewers(postID int, viewerIDs []int) error {
	query := `INSERT INTO friends (post_id, friend_id) VALUES (?, ?)`

	stmt, err := Db.Prepare(query)
	if err != nil {
		return err
	}
	for _, viewerID := range viewerIDs {
		_, err := stmt.Exec(postID, viewerID)
		if err != nil {
			return err
		}
	}
	return nil
}

// reactions functions

func AddOrUpdateReaction(userID, postID int, reactionType string) error {
	var existingID int
	checkQuery := "SELECT id FROM reactions WHERE user_id = ? AND post_id = ?"
	err := Db.QueryRow(checkQuery, userID, postID).Scan(&existingID)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	if err == sql.ErrNoRows {
		insertQuery := `
			INSERT INTO reactions (user_id, post_id, reaction_type)
			VALUES (?, ?, ?)
		`
		_, err := Db.Exec(insertQuery, userID, postID, reactionType)
		return err
	} else {
		// Reaction exists → update it
		updateQuery := `
			UPDATE reactions
			SET reaction_type = ?, date = CURRENT_TIMESTAMP
			WHERE id = ?
		`
		_, err := Db.Exec(updateQuery, reactionType, existingID)
		return err
	}
}

func InsertComment(comment *utils.Comment) (int, error) {
	query := `
		INSERT INTO comments (user_id, post_id, comment, imagePath, date) 
		VALUES (?, ?, ?, ?, strftime('%s', 'now'))
	`

	result, err := Db.Exec(query, comment.UserId, comment.PostId, comment.Content, comment.ImagePath)
	if err != nil {
		fmt.Println("Error inserting comment:", err)
		return 0, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(lastId), nil
}

// InsertReaction adds a new reaction to the database
func InsertReaction(reaction *utils.Reaction) (int, error) {
	query := `
		INSERT INTO reactions (user_id, post_id, reaction_type, date) 
		VALUES (?, ?, ?, strftime('%s', 'now'))
	`

	result, err := Db.Exec(query, reaction.UserId, reaction.PostId, reaction.ReactionType)
	if err != nil {
		fmt.Println("Error inserting reaction:", err)
		return 0, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(lastId), nil
}

