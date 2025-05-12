package controllers

import (
	"fmt"
	"net/http"

	"social-network/models"
	"social-network/utils"
)

func GetUsers(w http.ResponseWriter, r *http.Request,userID int) {
	// cookie, err := r.Cookie("token")
	// if err != nil {
	// 	utils.WriteJSON(w, map[string]string{"error": "Unauthorized"}, http.StatusUnauthorized)
	// 	return
	// }
	// userID, err := models.Get_session(cookie.Value)
	// if err != nil {
	// 	utils.WriteJSON(w, map[string]string{"error": "Session not found"}, http.StatusUnauthorized)
	// 	return
	// }
	users, err := models.GetUserFriends(userID)
	if err != nil {
		fmt.Println("can't fetch users")
		return
	}
	utils.WriteJSON(w, users,http.StatusAccepted)
}


