package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"social-network/models"
	"social-network/utils"
)

// UserProfileInfo holds user avatar and name information
type UserProfileInfo struct {
	Avatar   string
	UserName string
}

// GetUserAvatarAndUserName retrieves the avatar URL and full name for a user by ID
func GetUserAvatarAndUserName(userId int) *utils.Regester {
	user, err := models.GetUserById(userId)
	if err != nil {
		log.Println(err)
	}
	return user
}

// AddComment handles adding a new comment to a post
func AddComment(w http.ResponseWriter, r *http.Request, userID int) {
	if r.Method != http.MethodPost {
		utils.WriteJSON(w, map[string]string{"error": "Method not allowed"}, http.StatusMethodNotAllowed)
		return
	}
	// Parse the form data
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		utils.WriteJSON(w, map[string]string{"error": "Failed to parse form data"}, http.StatusBadRequest)
		return
	}

	postData := r.FormValue("commentData")
	fmt.Println("this id postdata ", postData)

	filepath, err := utils.UploadImage(r)
	if err != nil {
		utils.WriteJSON(w, map[string]string{"error": err.Error()}, http.StatusInternalServerError)
		fmt.Println("Upload Image error:", err)
		return
	}

	var comment utils.Comment
	if err := json.Unmarshal([]byte(postData), &comment); err != nil {
		utils.WriteJSON(w, map[string]string{"error": "Invalid request data"}, http.StatusBadRequest)
		return
	}

	comment.ImagePath = filepath

	if strings.TrimSpace(comment.Content) == "" && filepath == "" {
		utils.WriteJSON(w, map[string]string{"error": "Empty Message"}, http.StatusBadRequest)
		return
	}
	// Set the user ID from the session
	comment.UserId = userID
	fmt.Println("comment", comment.UserId, comment.Content, comment.PostId)
	// Check if post exists
	postExists, err := models.PostExists(comment.PostId)
	if err != nil || !postExists {
		utils.WriteJSON(w, map[string]string{"error": "Post not found"}, http.StatusNotFound)
		return
	}

	// Check if user has access to interact with this post based on privacy settings
	canAccess, err := models.CanUserAccessPost(userID, comment.PostId)
	if err != nil {
		utils.WriteJSON(w, map[string]string{"error": "Error checking post access: " + err.Error()}, http.StatusInternalServerError)
		return
	}
	if !canAccess {
		utils.WriteJSON(w, map[string]string{"error": "You don't have permission to comment on this post"}, http.StatusForbidden)
		return
	}

	// Insert the comment into the database
	commentID, err := models.InsertComment(&comment)
	if err != nil {
		utils.WriteJSON(w, map[string]string{"error": "Failed to save comment"}, http.StatusInternalServerError)
		return
	}

	// Set the generated ID and return the comment
	comment.Id = commentID

	// Get user details to include in response
	user, err := models.GetUserById(userID)
	if err == nil {
		comment.UserName = user.FirstName + " " + user.LastName
	}

	// Set the date and default values
	comment.Date = utils.GetCurrentDate()
	userdata := GetUserAvatarAndUserName(userID)
	comment.UserAvatar = user.Avatar
	comment.UserName = fmt.Sprintf("%s %s", userdata.FirstName, userdata.LastName)

	utils.WriteJSON(w, comment, http.StatusOK)
}

// GetComments retrieves all comments for a specific post
func GetComments(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.WriteJSON(w, map[string]string{"error": "Method not allowed"}, http.StatusMethodNotAllowed)
		return
	}
	offsetStr := r.URL.Query().Get("offset")
	limitStr := r.URL.Query().Get("limit")
	fmt.Println(offsetStr)

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

	// Get post ID from query parameters
	postIdStr := r.URL.Query().Get("postId")
	if postIdStr == "" {
		utils.WriteJSON(w, map[string]string{"error": "Post ID is required"}, http.StatusBadRequest)
		return
	}
	log.Println("working")
	postId, err := strconv.Atoi(postIdStr)
	if err != nil {
		utils.WriteJSON(w, map[string]string{"error": "Invalid post ID"}, http.StatusBadRequest)
		return
	}

	// Get comments from database
	comments, err := models.GetCommentsByPostId(postId, limit, offset)
	if err != nil {
		utils.WriteJSON(w, map[string]string{"error": "Failed to retrieve comments"}, http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, comments, http.StatusOK)
}
