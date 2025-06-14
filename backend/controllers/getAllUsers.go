package controllers

import (
	"fmt"
	"net/http"

	"social-network/models"
	"social-network/utils"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request, userid int) {
	if r.Method != "GET" {
		utils.WriteJSON(w, map[string]string{"error": "Method Not Allowed"}, http.StatusMethodNotAllowed)
		return
	}
	users, err := models.GetThemAll(userid)

	if err != nil {
		fmt.Println("we cant get all users",err)
		return
	}
	utils.WriteJSON(w, users, 200)
}
