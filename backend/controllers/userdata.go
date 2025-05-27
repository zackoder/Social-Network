package controllers

import (
	"fmt"
	"net/http"

	"social-network/models"
	"social-network/utils"
)

func UserData(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.WriteJSON(w, map[string]string{"error": "Not allowed"}, http.StatusMethodNotAllowed)
		return
	}
	cookie, _ := r.Cookie("token")
	fmt.Println(cookie.Value)
	
	var userD utils.UserD
	err := models.Db.QueryRow(
		`SELECT u.id, u.first_name, u.avatar
            FROM users u 
            JOIN sessions s on s.user_id = u.id 
            WHERE token = ?`, cookie.Value).Scan(&userD.Id, &userD.Firstname, &userD.Avatar)
	if err != nil {
		fmt.Println(err)
	}
	if userD.Avatar != "" {
		userD.Avatar = r.Host + userD.Avatar
	}
	utils.WriteJSON(w, userD, 200)
}
