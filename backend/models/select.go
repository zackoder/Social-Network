package models

import (
	"fmt"
	"time"

	"social-network/utils"
)

func QueryPosts(offset int) []utils.Post {
	var posts []utils.Post
	queryPosts := `SELECT * FROM posts`

	rows, err := Db.Query(queryPosts)
	if err != nil {
		return nil
	}
	defer rows.Close()
	for rows.Next() {
		var post utils.Post
		err := rows.Scan(&post.Id, &post.Privacy, &post.Title, &post.Content, &post.Poster, &post.Image, &post.CreatedAt)
		if err != nil {
			fmt.Println("scaning error:", err)
		}
		posts = append(posts, post)
	}
	return posts
}

func GetProfilePost(user_id, offset int) []utils.Post {
	var posts []utils.Post
	fmt.Printf("Querying posts for user_id=%d with offset=%d\n", user_id, offset)

	query := "SELECT * FROM posts WHERE user_id = ? LIMIT 10 OFFSET ?"
	rows, err := Db.Query(query, user_id, offset)
	if err != nil {
		fmt.Println("Error querying posts:", err)
		return nil
	}
	defer rows.Close()
	for rows.Next() {
		var post utils.Post

		err := rows.Scan(&post.Id, &post.Privacy, &post.Title, &post.Content, &post.Poster, &post.Image, &post.CreatedAt)
		if err != nil {
			fmt.Println("error scaning the rows", err)
		}
		posts = append(posts, post)
	}
	if err = rows.Err(); err != nil {
		fmt.Println("Error during rows iteration:", err)
		return nil
	}
	return posts
}

func IsPrivateProfile(followed string) (bool, error) {
	query := "SELECT privacy FROM users WHERE id = ?"
	var privacy string
	err := Db.QueryRow(query, followed).Scan(&privacy)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	return privacy == "private", nil
}
func CheckPostPrivacy(post string) (string, error) {
	query := "SELECT post_privacy FROM posts WHERE id = ?"
		var privacy string
	err := Db.QueryRow(query,post).Scan(&privacy)
	if err != nil {
		fmt.Println("is privet post",err)
		return "",err
	}
	return privacy,nil
}
///////////////////////////login///////////////////////////////////////////

func ValidCredential(userData *utils.User) error {
	query := `SELECT id, password FROM users WHERE nickname = ? OR email = ?;`
	err := Db.QueryRow(query, userData.Nickname, userData.Email).Scan(&userData.ID, &userData.Password)
	if err != nil {
		return err
	}
	return err
}

func GetActiveSession(userData *utils.User) (bool, error) {
	var exists bool
	currentTime := time.Now()
	fmt.Println(currentTime)
	query := `SELECT EXISTS(SELECT 1 FROM sessions WHERE user_id = ? );`
	err := Db.QueryRow(query, userData.ID).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}


func Get_session(ses string) (int, error) {
	var sessionid int
	query := `SELECT user_id FROM sessions WHERE token = ?`
	err := Db.QueryRow(query, ses).Scan(&sessionid)
	if err != nil {
		return 0, err
	}
	return sessionid, nil
}

