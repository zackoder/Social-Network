package controllers

import (
	"fmt"
	"log"
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
	var userId int
	var err error

	if (r.URL.Query().Get("id")) != "" {
		userId, err = strconv.Atoi((r.URL.Query().Get("id")))
		if err != nil {
			fmt.Println("we cant convert the id", err)
			utils.WriteJSON(w, map[string]string{"error": "data not available "}, http.StatusBadRequest)
			return
		}
	}

	if userId != 0 {
		MyFollowers, err := models.GetFollowers(userId)
		fmt.Println(MyFollowers)
		if err != nil {
			utils.WriteJSON(w, map[string]string{"error": "We Can't Fetsh Followers List For Posts"}, http.StatusInternalServerError)
			return
		}
		utils.WriteJSON(w, MyFollowers, http.StatusAccepted)
		return
	}

	log.Println("user id", userID)
	UserFollowers, err := models.GetFollowers(userID)
	if err != nil {
		utils.WriteJSON(w, map[string]string{"error": "We Can't Fetsh Followers List For Profile"}, http.StatusInternalServerError)
		return
	}
	utils.WriteJSON(w, UserFollowers, http.StatusAccepted)
}

func GetfollowingsForProfile(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		utils.WriteJSON(w, map[string]string{"error": "Method Not Allowed"}, http.StatusMethodNotAllowed)
		return
	}
	userId, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		fmt.Println("111111we cant convert the id")
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
