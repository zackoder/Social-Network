package models

import (
	"fmt"
	"net/http"
	"os"
	"social-network/utils"
)

func QueryPosts(offset int, r *http.Request) []utils.Post {
	var posts []utils.Post
	queryPosts := `SELECT * FROM posts`
	host := r.Host
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
		if post.Image != "" {
			post.Image = host + post.Image
		}
		posts = append(posts, post)
	}
	return posts
}

func GetPrivecy(followed string) (string, error) {
	getPrivacy := "SELECT privacy FROM users WHERE id = ?"
	var privacy string
	err := Db.QueryRow(getPrivacy, followed).Scan(&privacy)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return privacy, nil
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
			os.Exit(10)
		}
		groups = append(groups, group_id)
	}
	return groups
}
