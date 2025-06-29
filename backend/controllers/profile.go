package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"social-network/models"
	"social-network/utils"
)

func HandleFollow(w http.ResponseWriter, r *http.Request, follower int) {
	if r.Method != http.MethodPost {
		utils.WriteJSON(w, map[string]string{"error": "Not allowed"}, http.StatusMethodNotAllowed)
		return
	}
	var noti utils.Notification
	followed := r.URL.Query().Get("followed")
	noti.Sender_id = follower
	noti.Target_id, _ = strconv.Atoi(followed)
	noti.Message = "follow request"
	privacy, err := models.IsPrivateProfile(followed)
	fmt.Println("privacy", privacy, "err", err)
	if err != nil {
		utils.WriteJSON(w, map[string]string{"resp": "user not found"}, 404)
		return
	}

	alreadyFriends, _ := models.FriendsCheckerForFollow(noti.Sender_id, noti.Target_id)
	if alreadyFriends {
		models.Deletfollow(follower, followed)
		utils.WriteJSON(w, map[string]string{"resp": "unfollowed seccessfoly"}, 200)
		return
	}
	fmt.Println(privacy)
	if !privacy {
		err := models.InsertFollow(follower, followed)
		if err != nil {
			utils.WriteJSON(w, map[string]string{"resp": "try to follow net time"}, 404)
			return
		}
		utils.WriteJSON(w, map[string]string{"resp": "followed seccessfoly"}, 200)
	} else {
		utils.WriteJSON(w, map[string]string{"resp": "follow request sent"}, 200)
		noti.Actor_id = noti.Sender_id
		noti.Id, _ = models.InsertNotification(noti)
		models.SelectMetaData(&noti)
		Broadcast(noti.Target_id, noti)
	}
}

func UpdatePrivacy(w http.ResponseWriter, r *http.Request, user_id int) {
	privacy := models.UpdateProfile(user_id)
	fmt.Println(privacy)
	utils.WriteJSON(w, map[string]string{"profile_status": privacy}, 200)
}

func UserData(w http.ResponseWriter, r *http.Request, user_id int) {
	if r.Method != http.MethodGet {
		utils.WriteJSON(w, map[string]string{"error": "Not allowed"}, http.StatusMethodNotAllowed)
		return
	}
	var userD utils.UserD
	user, err := models.GetUserById(user_id)
	if err != nil {
		fmt.Println(err)
	}
	userD.Firstname = user.FirstName
	userD.Id = user_id
	userD.Avatar = r.Host + user.Avatar

	utils.WriteJSON(w, userD, 200)
}
