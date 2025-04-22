package models

import (
	"fmt"
	"social-network/utils"
)

func QueryPosts(offset int, host string) []utils.Post {
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
