package models

import (
	"fmt"
	"net/http"
	"time"

	"social-network/utils"
)

func QueryPosts(offset int, r *http.Request) []utils.Post {
	host := r.Host
	var posts []utils.Post
	// cookie, _ := r.Cookie("token")
	if 5 >= 4 {

	}
	// id := 5
	queryPosts := `
		SELECT p.id, p.title, p.content, p.imagePath, p.createdAt, p.user_id, p.post_privacy, u.first_name 
		FROM posts p
		JOIN users u ON u.id = p.user_id
		`

	rows, err := Db.Query(queryPosts)
	if err != nil {
		fmt.Println("ana hnaa", err)
		return nil
	}
	defer rows.Close()
	for rows.Next() {
		var post utils.Post
		err := rows.Scan(&post.Id, &post.Title, &post.Content, &post.Image, &post.CreatedAt, &post.Poster_id, &post.Privacy, &post.Poster_name)
		if err != nil {
			fmt.Println("scaning error:", err)
		}
		if post.Image != "" {
			post.Image = host + post.Image
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

		err := rows.Scan(&post.Id, &post.Privacy, &post.Title, &post.Content, &post.Poster_id, &post.Image, &post.CreatedAt)
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
	fmt.Println(privacy)
	return privacy == "privet", nil
}

func CheckPostPrivacy(post string) (string, error) {
	query := "SELECT post_privacy FROM posts WHERE id = ?"
	var privacy string
	err := Db.QueryRow(query, post).Scan(&privacy)
	if err != nil {
		fmt.Println("is privet post", err)
		return "", err
	}
	return privacy, nil
}

///////////////////////////login///////////////////////////////////////////

func ValidCredential(userData *utils.User) error {
	query := `SELECT id, password FROM users WHERE nickname = ? OR email = ?;`
	err := Db.QueryRow(query, userData.Email, userData.Email).Scan(&userData.ID, &userData.Password)
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

func GetClientGroups(user_id int) []int {
	var groups []int
	selectGroups := "SELECT group_id FROM group_members WHERE user_id = ?"
	rows, err := Db.Query(selectGroups, user_id)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer rows.Close()
	for rows.Next() {
		var group_id int
		if err := rows.Scan(&group_id); err != nil {
			fmt.Println(err)
		}
		groups = append(groups, group_id)
	}
	return groups
}

func FriendsChecker(Sender_id, Reciever_id int) (bool, error) {
	query := "SELECT EXISTs(SELECT 1 FROM followers WHERE follower_id = ? AND followed_id = ? OR follower_id = ? AND followed_id = ?)"
	var friends bool
	err := Db.QueryRow(query, Sender_id, Reciever_id, Reciever_id, Sender_id).Scan(&friends)
	return friends, err
}
