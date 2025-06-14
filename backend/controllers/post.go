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

func AddPost(w http.ResponseWriter, r *http.Request, userId int) {
	if r.Method != http.MethodPost {
		utils.WriteJSON(w, map[string]string{"error": "Method Not allowd"}, http.StatusMethodNotAllowed)
		return
	}
	var post utils.Post
	post.Poster_id = userId
	host := r.Host
	postData := r.FormValue("postData")

	err := json.Unmarshal([]byte(postData), &post)
	if err != nil {
		utils.WriteJSON(w, map[string]string{"error": "internal server error\nparsing post"}, http.StatusInternalServerError)
		fmt.Println("unmarshal err:", err)
		return
	}

	filepath, err := utils.UploadImage(r)
	if err != nil {
		utils.WriteJSON(w, map[string]string{"error": err.Error()}, http.StatusBadRequest)
		fmt.Println("Upload Image error:", err)
		return
	}
	if filepath == "" || (strings.TrimSpace(post.Title) == "" && strings.TrimSpace(post.Content) == "") {
		utils.WriteJSON(w, map[string]string{"error": "title or content is empty"}, http.StatusBadRequest)
		return
	}

	post.Image = filepath

	user, _ := models.GetUserById(userId)
	// log.Println(user)
	post.Poster_name = user.FirstName
	post.Avatar = host + user.Avatar

	post.Image = filepath

	fmt.Println(post.Poster_id)
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

	if filepath != "" {
		post.Image = host + post.Image

		if filepath != "" {
			post.Image = host + post.Image
		}
	}
	utils.WriteJSON(w, post, 200)
}

func Posts(w http.ResponseWriter, r *http.Request) {
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

	posts := models.QueryPosts(limit, offset, r)
	utils.WriteJSON(w, posts, 200)
}

// func GetProfilePosts(w http.ResponseWriter, r *http.Request) {
// 	cookie, err := r.Cookie("token")
// 	if err != nil {
// 		return
// 	}

// 	profilOwnerId := r.URL.Query().Get("id")

// 	userId, err := models.Get_session(cookie.Value)
// 	if err != nil {
// 		utils.WriteJSON(w, map[string]string{"error": "user id not found "}, http.StatusNotFound)

// 		return
// 	}
// 	useridstr := strconv.Itoa(userId)
// 	if profilOwnerId == useridstr {
// 		ProfilePosts := models.GetProfilePost(userId, 0)
// 		utils.WriteJSON(w, ProfilePosts, 200)
// 	} else if profilOwnerId != useridstr {
// 		profilPrivacy, err := models.IsPrivateProfile(profilOwnerId)
// 		if err != nil {
// 			utils.WriteJSON(w, map[string]string{"error": "not found"}, http.StatusNotFound)
// 		}
// 		if !profilPrivacy {
// 			profileOwnerId, err := strconv.Atoi(profilOwnerId)
// 			if err != nil {
// 				fmt.Println("we cant convert")
// 				return
// 			}
// 			userPostsForDisplay := models.GetProfilePost(profileOwnerId, 0)
// 			utils.WriteJSON(w, userPostsForDisplay, 200)

// 		} else if profilPrivacy {
// 			// checkPostPrivacy,err := models.CheckPostPrivacy()
// 		}
// 	}
// }
