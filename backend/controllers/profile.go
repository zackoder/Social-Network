package controllers

import (
	"fmt"
	"net/http"
	"social-network/models"
	"social-network/utils"
	"strconv"
)

func HandleFollow(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriteJSON(w, map[string]string{"error": "Not allowed"}, http.StatusMethodNotAllowed)
		return
	}
	var noti utils.Notification
	follower := r.URL.Query().Get("follower")
	followed := r.URL.Query().Get("followed")
	noti.Sender_id, _ = strconv.Atoi(follower)
	noti.Target_id, _ = strconv.Atoi(followed)
	noti.Message = "follow request"
	privacy, err := models.IsPrivateProfile(followed)
	fmt.Println("privacy", privacy, "err", err)
	if err != nil {
		utils.WriteJSON(w, map[string]string{"resp": "user not found"}, 404)
		return
	}

	alreadyFriends, _ := models.FriendsChecker(noti.Sender_id, noti.Target_id)
	if alreadyFriends {
		models.Deletfollow(follower, followed)
		utils.WriteJSON(w, map[string]string{"resp": "unfollowed seccessfoly"}, 200)
		return
	}
	if !privacy {
		err := models.InsertFollow(follower, followed)
		if err != nil {
			utils.WriteJSON(w, map[string]string{"resp": "try to follow net time"}, 404)
			return
		}
		utils.WriteJSON(w, map[string]string{"resp": "followed seccessfoly"}, 200)
	} else {
		utils.WriteJSON(w, map[string]string{"resp": "follow request sent"}, 200)
		BroadcastNotification(noti)
		noti.Actor_id = noti.Sender_id
		models.InsertNotification(noti)
	}
}

func UpdatePrivacy(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	test := models.UpdateProfile(id)
	utils.WriteJSON(w, map[string]string{"test": test}, 200)
}

func UserData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("22222222222222222222222222")
	if r.Method != http.MethodGet {
		utils.WriteJSON(w, map[string]string{"error": "Not allowed"}, http.StatusMethodNotAllowed)
		return
	}
	cookie, _ := r.Cookie("token")
	
	var userD utils.UserD
	err := models.Db.QueryRow(
		`SELECT u.id, u.first_name, u.avatar
            FROM users u 
            JOIN sessions s on s.id = u.id 
            WHERE token = ?`, cookie.Value).Scan(&userD.Id, &userD.Firstname, &userD.Avatar)
	if err != nil {
		fmt.Println(err)
	}
	if userD.Avatar != "" {
		userD.Avatar = r.Host + userD.Avatar
	}
	utils.WriteJSON(w, userD, 200)
}