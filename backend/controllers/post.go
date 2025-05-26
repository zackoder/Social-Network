package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"social-network/models"
	"social-network/utils"
)

func AddPost(w http.ResponseWriter, r *http.Request, user_id int) {
	if r.Method != http.MethodPost {
		utils.WriteJSON(w, map[string]string{"error": "Method Not allowd"}, http.StatusMethodNotAllowed)
		return
	}

	var post utils.Post
	post.Poster_id = user_id

	host := r.Host

	postData := r.FormValue("postData")

	err := json.Unmarshal([]byte(postData), &post)
	if err != nil {
		utils.WriteJSON(w, map[string]string{"error": "internal server error\nparsing post"}, http.StatusInternalServerError)
		fmt.Println("unmarshal err:", err)
		return
	}

	if strings.TrimSpace(post.Title) == "" || strings.TrimSpace(post.Content) == "" {
		utils.WriteJSON(w, map[string]string{"error": "title or content is empty"}, http.StatusBadRequest)
		return
	}

	filepath, err := utils.UploadImage(r)
	if err != nil {
		utils.WriteJSON(w, map[string]string{"error": err.Error()}, http.StatusInternalServerError)
		fmt.Println("Upload Image error:", err)
		return
	}

	post.Image = filepath

	if filepath == "" {
		post.Image = "/uploads/defaulte.jpg"
	}

	user, _ := models.GetUserById(user_id)
	post.Poster_name = user.FirstName
	post.Avatar = host + user.Avatar

	post.Id, err = models.InsertPost(post)
	if err != nil {
		utils.RemoveIMG(filepath)
		utils.WriteJSON(w, map[string]string{"error": "internal server error\ninserting post"}, http.StatusInternalServerError)
		return
	}

	if len(post.Friendes) != 0 {
		models.InsertFriends(post.Id, post.Friendes)
		post.Friendes = []int{}
	}
	post.Image = host + post.Image
	utils.WriteJSON(w, post, 200)
}

/*
func GetProfilePosts(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		utils.WriteJSON(w, map[string]string{"error": "Unauthorized"}, http.StatusUnauthorized)
		return
	}

	profileOwnerIDStr := r.URL.Query().Get("id")
	fmt.Println(profileOwnerIDStr)
	if profileOwnerIDStr == "" {
		utils.WriteJSON(w, map[string]string{"error": "Profile ID is missing"}, http.StatusBadRequest)
		return
	}

	viewerID, err := models.Get_session(cookie.Value)
	if err != nil {
		utils.WriteJSON(w, map[string]string{"error": "Session not found"}, http.StatusUnauthorized)
		return
	}
	fmt.Println(viewerID)
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
		// if the  profile is public we show all posts exept the privet ones and the almostPrivet posts
		// we check them one by one we fetch them in case the visiter is a follower .
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

	// here we fetch privet posts only for people that are allowed to see them by
	// hecking the the user id across the privet post viewrs that stors the post
	// with people allowed to see it
	posts, err := models.GetAllowedPosts(profileOwnerID, viewerID)
	if err != nil {
		utils.WriteJSON(w, map[string]string{"error": "Failed to fetch posts"}, http.StatusInternalServerError)
		return
	}
	utils.WriteJSON(w, posts, 200)
}

*/

func GetProfilePosts(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		return
	}

	profilOwnerId := r.URL.Query().Get("id")

	userId, err := models.Get_session(cookie.Value)
	if err != nil {
		utils.WriteJSON(w, map[string]string{"error": "user id not found "}, http.StatusNotFound)

		return
	}
	useridstr := strconv.Itoa(userId)
	if profilOwnerId == useridstr {
		ProfilePosts, err := models.GetProfilePost(userId, 0)
		if err != nil {
			// tanshofo shno ndiro fl error
		}
		utils.WriteJSON(w, ProfilePosts, 200)
	} else if profilOwnerId != useridstr {
		profilPrivacy, err := models.IsPrivateProfile(profilOwnerId)
		if err != nil {
			utils.WriteJSON(w, map[string]string{"error": "not found"}, http.StatusNotFound)
		}
		if !profilPrivacy {
			profileOwnerId, err := strconv.Atoi(profilOwnerId)
			if err != nil {
				fmt.Println("we cant convert")
				return
			}
			userPostsForDisplay, err := models.GetProfilePost(profileOwnerId, 0)
			if err != nil {
				// tanshofo shno ndiro fl error
			}
			utils.WriteJSON(w, userPostsForDisplay, 200)

		} else if profilPrivacy {
			// checkPostPrivacy,err := models.CheckPostPrivacy()
		}
	}
}
