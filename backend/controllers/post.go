package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"social-network/models"
	"social-network/utils"
)

func AddPost(w http.ResponseWriter, r *http.Request, userId int) {
	fmt.Println("userid",userId)
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

	user, _ := models.GetUserById(userId)
	log.Println(user)
	post.Poster_name = user.FirstName
	post.Avatar = host + user.Avatar  

	fmt.Println("post", r.URL.Query().Get("id"))

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
	}
	utils.WriteJSON(w, post, 200)
}

func Posts(w http.ResponseWriter, r *http.Request) {
	offset := 0
	posts := models.QueryPosts(offset, r)
	utils.WriteJSON(w, posts, 200)
}
