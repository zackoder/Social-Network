package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"social-network/models"
	"social-network/utils"
)
 
 

func GetProfilePosts(w http.ResponseWriter, r *http.Request, userId int) {
	profileOwnerIDStr := r.URL.Query().Get("id")
	fmt.Println(profileOwnerIDStr)
	if profileOwnerIDStr == "" {
		utils.WriteJSON(w, map[string]string{"error": "Profile ID is missing"}, http.StatusBadRequest)
		return
	}

	fmt.Println(userId)
	// in casse u wanna see ur profile
	if strconv.Itoa(userId) == profileOwnerIDStr {
		allPosts, err := models.GetProfilePost(userId, 0)
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
	
	if !profilePrivacy && !isFollower {
		publicPosts, err := models.GetPuclicPosts(profileOwnerID)
		if err != nil {
			utils.WriteJSON(w, map[string]string{"error": "Internal Server Error"}, http.StatusInternalServerError)
		}
		utils.WriteJSON(w, publicPosts, 200)
	}
	/*else if profilePrivacy && isFollower {
	   // if the  profile is public we show all posts exept the privet ones and the almostPrivet posts
	   // we check them one by one we fetch them in case the visiter is a follower .
	   publicAnAlmstPublicPosts, err := models.GetPublicAndAlmostPrivatePosts(profileOwnerID, userId)
	   if err != nil {
		   utils.WriteJSON(w, map[string]string{"error": "Failed to fetch posts"}, http.StatusInternalServerError)
		   return
	   }
	   utils.WriteJSON(w, publicAnAlmstPublicPosts, 200)
	   return
	}
	else if  !isFollower {
	   utils.WriteJSON(w, map[string]string{"error": "This profile is private"}, http.StatusForbidden)
	   return
	}*/
	
	// here we fetch the public and almost privet posts and the ones only for people that are allowed to see them by
	// checking the the user id across the privet post viewrs that stors the post
	// with people allowed to see it
	if /*profilePrivacy && isFollower  || !profilePrivacy &&*/ isFollower   {
		posts, err := models.GetAllowedPosts(profileOwnerID, userId)
		fmt.Println("this one is for the privat posts")
		if err != nil {
			utils.WriteJSON(w, map[string]string{"error": "Failed to fetch posts01"}, http.StatusInternalServerError)
			return
		}
		fmt.Println(posts)
		utils.WriteJSON(w, posts, 200)
	}
}
