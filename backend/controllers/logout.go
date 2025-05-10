package controllers

import (
 
	"net/http"

	utils "social-network/utils"
	models "social-network/models"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriteJSON(w, map[string]string{"error":"Method Not Allowed"} , http.StatusMethodNotAllowed)
		return
	}
	token, _ := r.Cookie("token")
	models.RemoveSessionFromDB(token)
	utils.ClearSession(w)
}




