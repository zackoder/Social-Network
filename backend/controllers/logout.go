package controllers

import (
 
	"net/http"

	utils "social-network/utils"
	models "social-network/models"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request, userId int) {
	if r.Method != http.MethodPost {
		utils.WriteJSON(w, map[string]string{"error":"Method Not Allowed"} , http.StatusMethodNotAllowed)
		return
	}
	
	models.DeleteSession(userId)
	utils.ClearSession(w)
}




