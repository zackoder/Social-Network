package models

import (
	"fmt"

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
	for rows.Next(){
		var post utils.Post
		
		err := rows.Scan(&post.Id, &post.Privacy, &post.Title, &post.Content, &post.Poster, &post.Image, &post.CreatedAt)
		if err != nil{
			fmt.Println("error scaning the rows",err)
		}
		posts = append(posts, post)
	}
	if err = rows.Err(); err != nil {
		fmt.Println("Error during rows iteration:", err)
		return nil
	}
	return posts
}

func GetProfilePrivecy(followed string) (string, error) {
	getPrivacy := "SELECT privacy FROM users WHERE id = ?"
	var privacy string
	err := Db.QueryRow(getPrivacy, followed).Scan(&privacy)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return privacy, nil
}
