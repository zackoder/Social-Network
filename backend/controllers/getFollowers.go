package controllers

import (
	"fmt"
	"net/http"

	"social-network/models"
	"social-network/utils"
)

func GetFollowers(w http.ResponseWriter, r *http.Request, userID int) {
	if r.Method != "GET" {
		utils.WriteJSON(w, map[string]string{"error": "Method Not Allowed"}, http.StatusMethodNotAllowed)
		return
	}
	followers, err := models.GetFollowers(userID)
	fmt.Println(followers)
	if err != nil {
		utils.WriteJSON(w, map[string]string{"error": "We Can't Fetsh Followers List"}, http.StatusInternalServerError)
		return
	}
	utils.WriteJSON(w, followers, http.StatusAccepted)
}
