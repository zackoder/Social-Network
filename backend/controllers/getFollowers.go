package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"social-network/models"
	"social-network/utils"
)

func GetFollowers(w http.ResponseWriter, r *http.Request, userID int) {
	if r.Method != "GET" {
		utils.WriteJSON(w, map[string]string{"error": "Method Not Allowed"}, http.StatusMethodNotAllowed)
		return
	}
	userId, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		fmt.Println("we cant convert the id")
		utils.WriteJSON(w, map[string]string{"error": "data not available "}, http.StatusBadRequest)
		return
	}
	followers, err := models.GetFollowers(userId)
	fmt.Println(followers)
	if err != nil {
		utils.WriteJSON(w, map[string]string{"error": "We Can't Fetsh Followers List"}, http.StatusInternalServerError)
		return
	}
	utils.WriteJSON(w, followers, http.StatusAccepted)
}

func Getfollowings(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		utils.WriteJSON(w, map[string]string{"error": "Method Not Allowed"}, http.StatusMethodNotAllowed)
		return
	}
	userId, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		fmt.Println("we cant convert the id")
		utils.WriteJSON(w, map[string]string{"error": "data not available "}, http.StatusBadRequest)
		return
	}
	followings, err := models.GetFollowings(userId)
	if err != nil {
		utils.WriteJSON(w, map[string]string{"error": "We Can't Fetsh Followers List"}, http.StatusInternalServerError)
		return
	}
	utils.WriteJSON(w, followings, http.StatusAccepted)
}
