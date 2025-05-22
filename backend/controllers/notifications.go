package controllers

import (
	"log"
	"net/http"

	"social-network/models"
	"social-network/utils"
)

func GetNotifications(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		utils.WriteJSON(w, map[string]string{"error": "Unauthorized"}, http.StatusUnauthorized)
		return
	}
	user_id, _ := models.Get_session(cookie.Value)
	notifications, err := models.SelectNotifications(user_id)
	if err != nil {
		log.Println(err)
	}
	utils.WriteJSON(w, notifications, 200)
}
