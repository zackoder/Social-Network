package controllers

import (
	"fmt"
	"net/http"

	"social-network/models"
	"social-network/utils"
)

func GetUsers(w http.ResponseWriter, r *http.Request,userID int) {
	if r.Method != "GET" {
		utils.WriteJSON(w, map[string]string{"error":"Method Not Allowd"}, http.StatusMethodNotAllowed)
	}
	fmt.Println("im used")
	fmt.Println(userID)
	// userID = 11
	users, err := models.GetUserFriends(userID)
	if err != nil {
		utils.WriteJSON(w,map[string]string{"error":"Internal Server Error"},http.StatusInternalServerError)
		fmt.Println("can't fetch users from DataBase")
		return
	}
	utils.WriteJSON(w, users,http.StatusAccepted)
}