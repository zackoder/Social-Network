package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"social-network/models"
	"social-network/utils"
)

func AddPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriteJSON(w, map[string]string{"error": "Method Not allowd"}, http.StatusMethodNotAllowed)
		return
	}

	host := r.Host
	postData := r.FormValue("postData")
	filepath, err := utils.UploadImage(r)
	if err != nil {
		utils.WriteJSON(w, map[string]string{"error": "internal server error\nUploading image"}, http.StatusInternalServerError)
		fmt.Println("Upload Image error:", err)
		return
	}

	var post utils.Post
	fmt.Println("post data", postData)
	err = json.Unmarshal([]byte(postData), &post)
	if err != nil {
		utils.WriteJSON(w, map[string]string{"error": "internal server error\nparsing post"}, http.StatusInternalServerError)
		fmt.Println("unmarshal err:", err)
		return
	}

	if filepath != "" {
		post.Image = host + filepath[1:]
	}
	post.Id, err = models.InsertPost(post)
	if err != nil {
		utils.WriteJSON(w, map[string]string{"error": "internal server error\ninserting post"}, http.StatusInternalServerError)
		return
	}
	utils.WriteJSON(w, post, 200)
}

func Posts(w http.ResponseWriter, r *http.Request) {
	offset := 0
	posts := models.QueryPosts(offset)
	utils.WriteJSON(w, posts, 200)
}
func GetProfilePosts(w http.ResponseWriter, r * http.Request){
	posts := models.GetProfilePost(1, 0)
	utils.WriteJSON(w,posts,200)
}
