package models

import (
	"fmt"
	"log"
	"net/http"
	"strings"
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

func IsPrivateProfile(followed string) (bool, error) {
	fmt.Println("is private profile", followed)
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

// checking if the notification already sent
func CheckNoti(noti utils.Notification) (bool, error) {
	var exists bool
	checknoti := `
		SELECT
    EXISTS (
        SELECT
            1
        FROM
            notifications
        WHERE
            actor_id = ?
            AND target_id = ?
            AND message = ?
    );
	`
	err := Db.QueryRow(checknoti, noti.Actor_id, noti.Target_id, noti.Message).Scan(&exists)
	return exists, err
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
