package models

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"social-network/utils"
)

func QueryPosts(offset int, r *http.Request) []utils.Post {
	host := r.Host
	var posts []utils.Post
	queryPosts := `SELECT p.id, p.post_privacy, p.title, p.content, p.user_id, u.first_name, p.imagePath, p.createdAt
	FROM posts p
	JOIN users u ON p.user_id = u.id`
	// cookie, _ := r.Cookie("token")
	if 5 >= 4 {
	}
	// id := 5
	rows, err := Db.Query(queryPosts)
	if err != nil {
		fmt.Println("ana hnaa", err)
		return nil
	}
	defer rows.Close()
	for rows.Next() {
		var post utils.Post
		err := rows.Scan(&post.Id, &post.Privacy, &post.Title, &post.Content, &post.Poster_id, &post.Poster_name, &post.Image, &post.CreatedAt)
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

func GetProfilePost(user_id, offset int) ([]utils.Post, error) {
	var posts []utils.Post
	fmt.Printf("Querying posts for user_id=%d with offset=%d\n", user_id, offset)

	query := "SELECT * FROM posts WHERE user_id = ? LIMIT 10 OFFSET ?"
	rows, err := Db.Query(query, user_id, offset)
	if err != nil {
		fmt.Println("Error querying posts:", err)
		return nil, err
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
		return nil, err
	}
	return posts, err
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
	err := Db.QueryRow(query, post).Scan(&privacy)
	if err != nil {
		fmt.Println("is private post", err)
		return "", err

	}
	return privacy, nil
}

///////////////////////////login///////////////////////////////////////////

func ValidCredential(userData *utils.User) error {
	fmt.Println("i was here")
	query := `SELECT id, password FROM users WHERE nickname = ? OR email = ?;`
	err := Db.QueryRow(query, userData.Email, userData.Email).Scan(&userData.ID, &userData.Password)
	if err != nil {
		return err
	}
	return err
}

func GetActiveSession(userData *utils.User) (bool, error) {
	fmt.Println("i was here ")
	var exists bool

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
		fmt.Println(err)
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

func GetFollowers(userID int) ([]utils.Regester, error) {
	query := `SELECT  f.follower_id, u.first_name FROM followers f 
	JOIN users u 
	ON f.follower_id = u.id 
	WHERE f.followed_id  = ? `
	rows, err := Db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	var followersInfo []utils.Regester

	for rows.Next() {
		var followerinfo utils.Regester
		if err := rows.Scan(&followerinfo.ID, &followerinfo.FirstName); err != nil {
			fmt.Println(followerinfo.ID)
			fmt.Println(followerinfo.FirstName)
			return nil, err
		}
		followersInfo = append(followersInfo, followerinfo)
	}
	return followersInfo, nil
}
func GetFollowings(userID int) ([]utils.Regester,error) {
		query := `SELECT  f.followed_id, u.first_name FROM followers f 
	JOIN users u 
	ON f.followed_id = u.id 
	WHERE f.follower_id  = ? `
	rows, err := Db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	var followersInfo []utils.Regester

	for rows.Next() {
		var followerinfo utils.Regester
		if err := rows.Scan(&followerinfo.ID, &followerinfo.FirstName); err != nil {
			fmt.Println(followerinfo.ID)
			fmt.Println(followerinfo.FirstName)
			return nil, err
		}
		followersInfo = append(followersInfo, followerinfo)
	}
	return followersInfo, nil

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

func GetFollowStatus(profileowner, userId int) bool {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM notifications WHERE message = 'follow request' AND actor_id = ? AND target_id = ?) `
	err := Db.QueryRow(query, userId, profileowner).Scan(&exists)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		} else {
			fmt.Println(err)
		}
	}
	return exists
}

func GetProfileStatus(profileowner, userId int) (string, error) {
	if profileowner == userId {
		private, err := IsPrivateProfile(strconv.Itoa(userId))
		if private {
			return "private", err
		} else {
			return "public", err
		}
	} else {
		follower, err := IsFollower(profileowner, userId)
		if follower {
			return "unfollow", err
		} else if GetFollowStatus(profileowner, userId) {
			return "follow sent", nil
		} else {
			return "follow", err
		}
	}
}

func GetPuclicPosts(userID int) ([]utils.Post, error) {
	var publicPosts []utils.Post
	query := `
	SELECT p.id, p.post_privacy, p.title, p.content, p.user_id, u.first_name, p.imagePath, p.createdAt
		FROM posts p
		JOIN users u ON p.user_id = u.id
		WHERE p.user_id = ?
		AND (
			p.post_privacy = 'public' 
		)
		ORDER BY p.createdAt 
		LIMIT ? OFFSET ?
	`

	rows, err := Db.Query(query, userID, 10, 0)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var post utils.Post
		// var createdAt int64
		err := rows.Scan(&post.Id, &post.Privacy, &post.Title, &post.Content,
			&post.Poster_id, &post.Poster_name, &post.Image, &post.CreatedAt)
		if err != nil {
			return nil, err
		}

		publicPosts = append(publicPosts, post)
	}

	return publicPosts, nil
}

func GetAllowedPosts(profileOwnerID int, viewerID int) ([]utils.Post, error) {
	query := `
	SELECT DISTINCT p.id, p.post_privacy, p.title, p.content, p.user_id, u.first_name, p.imagePath, p.createdAt
	FROM posts p
	JOIN users u ON p.user_id = u.id
	LEFT JOIN friends ppv ON p.id = ppv.post_id
	WHERE p.user_id = ?
	AND (
		p.post_privacy = 'public'
		OR (p.post_privacy = 'almostPrivate')
		OR (p.post_privacy = 'private' AND ppv.friend_id = ?)
		)
		ORDER BY p.createdAt DESC
		`

	rows, err := Db.Query(query, profileOwnerID, viewerID)
	if err != nil {
		// fmt.Println(err)
		return nil, err
	}
	var posts []utils.Post
	for rows.Next() {
		var post utils.Post
		err := rows.Scan(&post.Id, &post.Privacy, &post.Title, &post.Content, &post.Poster_id, &post.Poster_name, &post.Image, &post.CreatedAt)
		if err != nil {
			fmt.Println("here im ", err)
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func GetAllowedPostsForFeed(profileOwnerIDs []int, viewerID int) ([]utils.Post, error) {
	if len(profileOwnerIDs) == 0 {
		return []utils.Post{}, nil // No user IDs means no posts
	}

	// Create placeholders (?, ?, ?) and args for user IDs
	placeholders := make([]string, len(profileOwnerIDs))
	args := make([]interface{}, len(profileOwnerIDs)+1)
	for i, id := range profileOwnerIDs {
		placeholders[i] = "?"
		args[i] = id
	}
	args[len(profileOwnerIDs)] = viewerID // for ppv.friend_id = ?

	userFilter := strings.Join(placeholders, ", ")

	query := fmt.Sprintf(`
		SELECT DISTINCT p.id, p.post_privacy, p.title, p.content, p.user_id, u.first_name, p.imagePath, p.createdAt
		FROM posts p
		JOIN users u ON p.user_id = u.id
		LEFT JOIN friends ppv ON p.id = ppv.post_id
		WHERE p.user_id IN (%s)
		AND (
			p.post_privacy = 'public'
			OR (p.post_privacy = 'almostPrivate')
			OR (p.post_privacy = 'private' AND ppv.friend_id = ?)
		)
		ORDER BY p.createdAt DESC
	`, userFilter)

	rows, err := Db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []utils.Post
	for rows.Next() {
		var post utils.Post
		err := rows.Scan(
			&post.Id,
			&post.Privacy,
			&post.Title,
			&post.Content,
			&post.Poster_id,
			&post.Poster_name,
			&post.Image,
			&post.CreatedAt,
		)
		if err != nil {
			fmt.Println("scan error:", err)
			return nil, err
		}
		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func GetUserFriends(userId int, host string) ([]utils.Regester, error) {
	var users []utils.Regester

	query := `
	SELECT DISTINCT u.id, u.first_name, u.last_name, u.avatar
	FROM users u
	INNER JOIN (
		SELECT followed_id AS friend_id FROM followers WHERE follower_id = ?
		UNION
		SELECT follower_id AS friend_id FROM followers WHERE followed_id = ?
	) f ON u.id = f.friend_id
	LEFT JOIN messages m
		ON (u.id = m.sender_id OR u.id = m.reciever_id)
		AND (m.sender_id = ? OR m.reciever_id = ?)
	WHERE u.id <> ?
	GROUP BY u.id
	ORDER BY
		MAX(m.creation_date) DESC,
		u.first_name ASC;
	`

	rows, err := Db.Query(query, userId, userId, userId, userId, userId)
	if err != nil {
		fmt.Println("DB query error:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var u utils.Regester
		err := rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Avatar)
		if err != nil {
			return nil, err
		}
		if u.Avatar != "" {
			u.Avatar = host + u.Avatar
		}

		users = append(users, u)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func FriendsChecker(Sender_id, Reciever_id int) (bool, error) {
	query := "SELECT EXISTS(SELECT 1 FROM followers WHERE follower_id = ? AND followed_id = ? OR follower_id = ? AND followed_id = ?)"
	var friends bool
	err := Db.QueryRow(query, Sender_id, Reciever_id, Reciever_id, Sender_id).Scan(&friends)
	return friends, err
}

func CheckSender(group_id, sender_id int) bool {
	query := "SELECT EXISTS(SELECT 1 FROM group_members WHERE group_id = ? AND user_id = ?)"
	var exists bool
	Db.QueryRow(query, group_id, sender_id).Scan(&exists)
	return exists
}

func IsMember(groupID, userID int) bool {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM group_members WHERE group_id = ? AND user_id = ?)"
	err := Db.QueryRow(query, groupID, userID).Scan(&exists)
	if err != nil {
		fmt.Println("Query error:", err)
	}
	return exists
}

func SearchGroupsInDatabase(tocken string) ([]utils.Groupe, error) {
	var Groups []utils.Groupe
	quirie := `SELECT * FROM groups WHERE title LIKE %?%`
	rows, err := Db.Query(quirie, tocken)
	if err != nil {
		fmt.Println("Error querying Groups", err)
		return nil, err
	}
	for rows.Next() {
		var Groupe utils.Groupe

		err = rows.Scan(&Groupe.CreatorId, &Groupe.Title, &Groupe.Description)
		if err != nil {
			fmt.Println("error scaning the rows", err)
			continue
		}
		Groups = append(Groups, Groupe)
	}
	return Groups, nil
}

func GetGroups(user_id int) []string {
	var res []string

	quirie0 := "SELECT group_id FROM group_members WHERE user_id = ?"
	rows, err := Db.Query(quirie0, user_id)
	if err != nil {
		fmt.Println("Error querying group_ids for user:", err)
		return nil
	}
	defer rows.Close()
	var groupIDs []int
	for rows.Next() {
		var group_id int
		if err := rows.Scan(&group_id); err != nil {
			fmt.Println("Error scanning group_id:", err)
			return nil
		}
		groupIDs = append(groupIDs, group_id)
	}
	if err := rows.Err(); err != nil {
		fmt.Println("Error with rows iteration:", err)
		return nil
	}

	if len(groupIDs) == 0 {
		return res
	}

	query := "SELECT name FROM groups WHERE id IN (?)"
	query = fmt.Sprintf(query, strings.Join(strings.Split(fmt.Sprint(groupIDs), " "), ","))
	row, err := Db.Query(query)
	if err != nil {
		fmt.Println("Error querying group names:", err)
		return nil
	}
	defer row.Close()

	for row.Next() {
		var groupName string
		if err := row.Scan(&groupName); err != nil {
			fmt.Println("Error scanning group name:", err)
			return nil
		}
		res = append(res, groupName)
	}

	if err := row.Err(); err != nil {
		fmt.Println("Error with rows iteration:", err)
		return nil
	}

	return res
}

func GetAllGroups() []string {
	res := []string{}

	Quirie := "SELECT name FROM groups"
	rows, err := Db.Query(Quirie)
	if err != nil {
		fmt.Println("Error querying names:", err)
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		var groupName string
		err := rows.Scan(&groupName)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return nil
		}
		res = append(res, groupName)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating over rows:", err)
		return nil
	}

	return res
}

func MyGroupes(user_id int) []string {
	var res []string
	quirie := "SELECT group_id FROM group_members WHERE user_id = ? AND role = 'creator' "
	rows, err := Db.Query(quirie, user_id)
	if err != nil {
		fmt.Println("Error querying group_ids for user:", err)
		return nil
	}
	defer rows.Close()
	var groupIDs []int
	for rows.Next() {
		var group_id int
		if err := rows.Scan(&group_id); err != nil {
			fmt.Println("Error scanning group_id:", err)
			return nil
		}
		groupIDs = append(groupIDs, group_id)
	}
	if err := rows.Err(); err != nil {
		fmt.Println("Error with rows iteration:", err)
		return nil
	}

	if len(groupIDs) == 0 {
		return res
	}
	query := "SELECT name FROM groups WHERE id IN (?)"
	query = fmt.Sprintf(query, strings.Join(strings.Split(fmt.Sprint(groupIDs), " "), ","))
	row, err := Db.Query(query)
	if err != nil {
		fmt.Println("Error querying group names:", err)
		return nil
	}
	defer row.Close()

	for row.Next() {
		var groupName string
		if err := row.Scan(&groupName); err != nil {
			fmt.Println("Error scanning group name:", err)
			return nil
		}
		res = append(res, groupName)
	}

	if err := row.Err(); err != nil {
		fmt.Println("Error with rows iteration:", err)
		return nil
	}

	return res
}

func SelectNotifications(user_id int) ([]utils.Notification, error) {
	var notis []utils.Notification
	quetyNotifications := `SELECT id, user_id, actor_id, target_id, message FROM notifications WHERE target_id = ?`
	rows, err := Db.Query(quetyNotifications, user_id)
	if err != nil {
		return notis, err
	}

	for rows.Next() {
		var noti utils.Notification
		if err := rows.Scan(&noti.Id, &noti.Sender_id, &noti.Actor_id, &noti.Target_id, &noti.Message); err != nil {
			log.Println("scaning notifacations error:", err)
		}
		notis = append(notis, noti)
	}

	defer rows.Close()
	return notis, nil
}

func SelectOneNoti(noti *utils.Notification) {
	queryNoti := "SELECT message, target_id, actor_id, user_id FROM notifications WHERE id = ?"
	err := Db.QueryRow(queryNoti, noti.Id).Scan(&noti.Message, &noti.Target_id, &noti.Actor_id, &noti.Sender_id)
	if err != nil {
		fmt.Println(err)
	}
}

func GetGroupOwner(group utils.Groupe_member) int {
	var owner int
	selectGroupOwner := `SELECT group_oner FROM groups WHERE id = ?`
	err := Db.QueryRow(selectGroupOwner, group.Groupe_id).Scan(&owner)
	if err != nil {
		log.Println("getting group owner id error", err)
	}
	return owner
}

func CheckGroup(group_id int) bool {
	var exists bool
	getGroup := "SELECT EXISTS(SELECT 1 FROM groups WHERE id = ?)"
	err := Db.QueryRow(getGroup, group_id).Scan(&exists)
	log.Println(err)
	return exists
}

// func GetReactionsCount(postID int) (map[string]int, error) {
// 	query := `
// 		SELECT reaction_type, COUNT(*) as count	FROM reactionsWHERE post_id = ?
// 		GROUP BY reaction_type
// 	`

// 	rows, err := Db.Query(query, postID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	summary := make(map[string]int)

// 	for rows.Next() {
// 		var reactionType string
// 		var count int
// 		if err := rows.Scan(&reactionType, &count); err != nil {
// 			return nil, err
// 		}
// 		summary[reactionType] = count
// 	}
// 	return summary, nil
// }

// func GetPublicAndAlmostPrivatePosts(profileOwnerID, viewerID int) ([]utils.Post, error) {
// 	query := `
// 	SELECT p.id, p.post_privacy, p.title, p.content, p.user_id, u.first_name, p.imagePath, p.createdAt
// 	FROM posts p
// 	JOIN users u ON p.user_id = u.id
// 	WHERE p.user_id = ?
// 	AND (
// 		p.post_privacy = 'public'
// 		OR (
// 			p.post_privacy = 'almostPrivate'
// 			AND EXISTS (
// 				SELECT 1 FROM followers f
// 				WHERE f.followed_id = p.user_id AND f.follower_id = ?
// 			)
// 		)
// 	)
// 	ORDER BY p.createdAt DESC
// 	LIMIT ? OFFSET ?
// 	`

// 	rows, err := Db.Query(query, profileOwnerID, viewerID, 10, 0)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var posts []utils.Post
// 	for rows.Next() {
// 		var post utils.Post
// 		// var createdAt int64
// 		err := rows.Scan(&post.Id, &post.Privacy, &post.Title, &post.Content,
// 			&post.Poster_id, &post.Poster_name, &post.Image , &post.CreatedAt)
// 		if err != nil {
// 			return nil, err
// 		}
// 		// post.CreatedAt = int(createdAt)
// 		posts = append(posts, post)
// 	}
// 	return posts, nil
// }

func IsUserRegistered(userData *utils.User) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE email = ?);`
	err := Db.QueryRow(query, userData.Email).Scan(&exists)
	return exists, err
}

func GetNotifications(userId int, limit int, offset int) ([]utils.Notification, error) {
	var notifications []utils.Notification

	query := `SELECT id, user_id, actor_id, target_id, message, is_read, created_at 
			  FROM notifications 
			  WHERE user_id = ? 
			  ORDER BY created_at DESC 
			  LIMIT ? OFFSET ?`

	rows, err := Db.Query(query, userId, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var notification utils.Notification
		var isRead int
		var createdAt string

		err := rows.Scan(
			&notification.Id,
			&notification.Sender_id,
			&notification.Actor_id,
			&notification.Target_id,
			&notification.Message,
			&isRead,
			&createdAt,
		)
		if err != nil {
			return nil, err
		}
		notifications = append(notifications, notification)
	}

	return notifications, nil
}
