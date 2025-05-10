package models

import (
	"fmt"
	"time"

	"social-network/utils"
)

func QueryPosts(offset int, host string) []utils.Post {
	fmt.Println("kantesti f mok")
	var posts []utils.Post
	queryPosts := `SELECT p.id, p.post_privacy, p.title, p.content, p.user_id, u.first_name, p.imagePath, p.createdAt
	FROM posts p
	JOIN users u ON p.user_id = u.id`


	rows, err := Db.Query(queryPosts)
	if err != nil {
		fmt.Println("ana hnaa",err)
		return nil
	}
	defer rows.Close()
	for rows.Next() {
		var post utils.Post
		err := rows.Scan(&post.Id, &post.Privacy, &post.Title, &post.Content, &post.Poster_id, &post.Poster_name , &post.Image, &post.CreatedAt)
		if err != nil {
			fmt.Println("scaning error:", err)
		}
		if post.Image != "" {
			post.Image = host + post.Image
		}
		if post.Image != "" {
			post.Image = host + post.Image
		}
		posts = append(posts, post)
	}
	return posts
}

func GetProfilePost(user_id, offset int) ([]utils.Post,error) {
	fmt.Println("kantesti f mok")
	var posts []utils.Post
	fmt.Printf("Querying posts for user_id=%d with offset=%d\n", user_id, offset)

	query := "SELECT * FROM posts WHERE user_id = ? LIMIT 10 OFFSET ?"
	rows, err := Db.Query(query, user_id, offset)
	if err != nil {
		fmt.Println("Error querying posts:", err)
		return nil,err
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
		return nil,err
	}
	fmt.Println(posts)
	return posts,err
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
	query := `SELECT user_id FROM sessions WHERE token = ?;`
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

func IsFollower(profileOwnerID int, viewerID int) (bool, error) {
	query := `SELECT COUNT(*) FROM followers WHERE followed_id = ? AND follower_id = ?`
	var count int
	err := Db.QueryRow(query, profileOwnerID, viewerID).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func GetFollowers(userID int) ([]int, error) {
	query := `SELECT follower_id FROM followers WHERE followed_id = ?`
	rows, err := Db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	var followerIDs []int
	for rows.Next() {
		var followerID int
		if err := rows.Scan(&followerID); err != nil {
			return nil, err
		}
		followerIDs = append(followerIDs, followerID)
	}
	return followerIDs, nil
}
func GetPublicAndAlmostPrivatePosts(profileOwnerID, viewerID int) ([]utils.Post, error) {
	query := `
	SELECT p.id, p.post_privacy, p.title, p.content, p.user_id, u.first_name, p.imagePath, p.createdAt
	FROM posts p
	JOIN users u ON p.user_id = u.id
	WHERE p.user_id = ?
	  AND (
		p.post_privacy = 'public'
		OR (
			p.post_privacy = 'almostPrivate'
			AND EXISTS (
				SELECT 1 FROM followers f
				WHERE f.followed_id = p.user_id AND f.follower_id = ?
			)
		)
	)
	ORDER BY p.createdAt DESC
	LIMIT ? OFFSET ?
	`

	rows, err := Db.Query(query, profileOwnerID, viewerID, 10, 0)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []utils.Post
	for rows.Next() {
		var post utils.Post
		var createdAt int64

		err := rows.Scan(&post.Id, &post.Privacy, &post.Title, &post.Content,
			&post.Poster_id, &post.Poster_name, &post.Image, &createdAt)
		if err != nil {
			return nil, err
		}

		post.CreatedAt = int(createdAt)
		posts = append(posts, post)
	}

	return posts, nil
}

func GetRegistration(id string) (utils.Regester, error) {
	query := `SELECT * FROM users WHERE id = ?`

	var data utils.Regester
	var Id int

	err := Db.QueryRow(query, id).Scan(
		&Id,             
		&data.NickName,     
		&data.FirstName, 
		&data.LastName, 
		&data.Age,              
		&data.Gender,        
		&data.Email,          
		&data.Avatar,         
		&data.Password,        
		&data.About_Me,        
		&data.Pravecy,        
	)
	if err != nil {
		return data, err
	}
	 data.Password = ""
	return data, nil
}

func GetAllowedPosts(profileOwnerID int, viewerID int) ([]utils.Post, error) {
	query := `
	SELECT DISTINCT p.id, p.post_privacy, p.title, p.content, p.user_id, u.first_name, p.imagePath, p.createdAt
	FROM posts p
	JOIN users u ON p.user_id = u.id
	LEFT JOIN private_post_viewers ppv ON p.id = ppv.post_id
	WHERE p.user_id = ?
	AND (
		p.post_privacy = 'public'
		OR (p.post_privacy = 'almostPrivate')
		OR (p.post_privacy = 'private' AND ppv.viewer_id = ?)
	)
	ORDER BY p.createdAt DESC
	`

	rows, err := Db.Query(query, profileOwnerID, viewerID)
	if err != nil {
		return nil, err
	}
	var posts []utils.Post
	for rows.Next() {
		var post utils.Post
		err := rows.Scan(&post.Id, &post.Privacy, &post.Title, &post.Content, &post.Poster_id, &post.Poster_name, &post.Image, &post.CreatedAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
