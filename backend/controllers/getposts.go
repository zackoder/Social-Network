package controllers

import (
	"net/http"
	"social-network/models"
	"social-network/utils"
	"strconv"
)

func Posts(w http.ResponseWriter, r *http.Request) {
	host := r.Host
	offset := 0
	posts := models.QueryPosts(offset, host)
	utils.WriteJSON(w, posts, 200)
}


func GetProfilePosts(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		utils.WriteJSON(w, map[string]string{"error": "Unauthorized"}, http.StatusUnauthorized)
		return
	}

	profileOwnerIDStr := r.URL.Query().Get("id")
	if profileOwnerIDStr == "" {
		utils.WriteJSON(w, map[string]string{"error": "Profile ID is missing"}, http.StatusBadRequest)
		return
	}

	viewerID, err := models.Get_session(cookie.Value)
	if err != nil {
		utils.WriteJSON(w, map[string]string{"error": "Session not found"}, http.StatusUnauthorized)
		return
	}

	// in casse u wanna see ur profile 
	if strconv.Itoa(viewerID) == profileOwnerIDStr {
		allPosts, err := models.GetProfilePost(viewerID, 0)  
		if err != nil {
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

	profileOwnerID, err := strconv.Atoi(profileOwnerIDStr)
	if err != nil {
		utils.WriteJSON(w, map[string]string{"error": "Invalid profile ID"}, http.StatusBadRequest)
		return
	}

	if !profilePrivacy {
		// if profile is public we show all posts exept the privet ones 
		publicPosts, err := models.GetPublicAndAlmostPrivatePosts(profileOwnerID, viewerID)
		if err != nil {
			utils.WriteJSON(w, map[string]string{"error": "Failed to fetch posts"}, http.StatusInternalServerError)
			return
		}
		utils.WriteJSON(w, publicPosts, 200)
		return
	}

	// if the profile is private we use the func IsFollower to chekck the list of followers if the visiter is amog the follower 
	// we show him the public posts + the olmost privet posts 
	isFollower, err := models.IsFollower(profileOwnerID, viewerID)
	if err != nil {
		utils.WriteJSON(w, map[string]string{"error": "Failed to check follower status"}, http.StatusInternalServerError)
		return
	}

	if !isFollower {
		utils.WriteJSON(w, map[string]string{"error": "This profile is private"}, http.StatusForbidden)
		return
	}

	 // here we fetch privet posts only for people tha are allowed to see them by 
	 // hecking the the user id across the privet post viewrs that stors the post 
	 // with people allowed to see it 
	posts, err := models.GetAllowedPosts(profileOwnerID, viewerID)
	if err != nil {
		utils.WriteJSON(w, map[string]string{"error": "Failed to fetch posts"}, http.StatusInternalServerError)
		return
	}
	utils.WriteJSON(w, posts, 200)
}
