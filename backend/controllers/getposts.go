package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"social-network/models"
	"social-network/utils"
)

func GetProfilePosts(w http.ResponseWriter, r *http.Request, userId int) {
	offsetStr := r.URL.Query().Get("offset")
	limitStr := r.URL.Query().Get("limit")
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		fmt.Println("offset", err)
		return
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		fmt.Println("limiit", err)
		return
	}

	profileOwnerIDStr := r.URL.Query().Get("id")
	if profileOwnerIDStr == "" {
		utils.WriteJSON(w, map[string]string{"error": "Profile ID is missing"}, http.StatusBadRequest)
		return
	}

	// in casse u wanna see ur profile
	if strconv.Itoa(userId) == profileOwnerIDStr {
		allPosts, err := models.GetProfilePost(userId, limit, offset)
		if err != nil && err != sql.ErrNoRows {
			utils.WriteJSON(w, map[string]string{"error": "Failed to fetch posts"}, http.StatusInternalServerError)
			return
		}
		utils.WriteJSON(w, allPosts, 200)
		return
	}

	// in case u wanna visit other prifile
	profilePrivacy, err := models.IsPrivateProfile(profileOwnerIDStr)
	if err != nil {
		utils.WriteJSON(w, map[string]string{"error": "Profile not found"}, http.StatusNotFound)
		return
	}
	fmt.Println("this is the profile privacy", profilePrivacy)

	profileOwnerID, err := strconv.Atoi(profileOwnerIDStr)
	if err != nil {
		utils.WriteJSON(w, map[string]string{"error": "Invalid profile ID"}, http.StatusBadRequest)
		return
	}

	isFollower, err := models.IsFollower(profileOwnerID, userId)
	// if the profile is private we use the func IsFollower to chekck the list of followers if the visiter is amog the follower
	// we show him the public posts + the almost privet posts
	fmt.Println("this is concerning the followers", isFollower)
	if err != nil {
		utils.WriteJSON(w, map[string]string{"error": "Failed to check follower status"}, http.StatusInternalServerError)
		return
	}
	// in case where the the profile is public and the ID request seeing the profile is not a follower
	// we fetch only the public posts
	if !profilePrivacy && !isFollower {
		publicPosts, err := models.GetPuclicPosts(profileOwnerID, limit, offset)
		// fmt.Println("public posts", err, publicPosts)
		if err != nil {
			fmt.Println("")
			utils.WriteJSON(w, map[string]string{"error": "Internal Server Error"}, http.StatusInternalServerError)
			return
		}
		utils.WriteJSON(w, publicPosts, 200)

		return
	}

	// here we fetch all posts  exept the privat ones we includ them only for people that are allowed to see them by
	// checking the the user id across the privet posts viewrs that stors the post
	// with people allowed to see it
	if isFollower {
		posts, err := models.GetAllowedPosts(profileOwnerID, userId, limit, offset)
		fmt.Println("this one is for the privat posts")
		if err != nil {
			utils.WriteJSON(w, map[string]string{"error": "Failed to fetch posts01"}, http.StatusInternalServerError)
			return
		}
		utils.WriteJSON(w, posts, 200)
	} else {
		utils.WriteJSON(w, map[string]string{"message": "this profile is private"}, 200)
	}
}
