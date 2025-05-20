package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"social-network/models"
	"social-network/utils"
	"strconv"
)

// UserProfileInfo holds user avatar and name information
type UserProfileInfo struct {
	Avatar   string
	UserName string
}

// GetUserAvatarAndUserName retrieves the avatar URL and full name for a user by ID
func GetUserAvatarAndUserName(userId int) *UserProfileInfo {
	user, err := models.GetUserById(userId)
	if err != nil {
		return &UserProfileInfo{
			Avatar:   "",
			UserName: "Anonymous",
		}
	}

	return &UserProfileInfo{
		Avatar:   user.Avatar,
		UserName: user.FirstName + " " + user.LastName,
	}
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
	// Get the file path for the uploaded image
	filepath, err := utils.UploadImage(r)
	if err != nil {
		utils.WriteJSON(w, map[string]string{"error": err.Error()}, http.StatusInternalServerError)
		fmt.Println("Upload Image error:", err)
		return
	}
	// // Get user session to identify the commenter
	// cookie, err := r.Cookie("token")
	// if err != nil {
	// 	utils.WriteJSON(w, map[string]string{"error": "Authentication required"}, http.StatusUnauthorized)
	// 	return
	// }

	// userId, err := models.Get_session(cookie.Value)
	// if err != nil {
	// 	utils.WriteJSON(w, map[string]string{"error": "Invalid session"}, http.StatusUnauthorized)
	// 	return
	// }

	// Parse the comment data
	var comment utils.Comment
	if err := json.Unmarshal([]byte(postData), &comment); err != nil {
		utils.WriteJSON(w, map[string]string{"error": "Invalid request data"}, http.StatusBadRequest)
		return
	}
	if filepath != "" {
		comment.ImagePath = filepath
	}

	// Set the user ID from the session
	comment.UserId = userID

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

	if filepath != "" {
		comment.ImagePath = r.Host + filepath
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

	comment.UserAvatar = GetUserAvatarAndUserName(userID).Avatar
	comment.UserName = GetUserAvatarAndUserName(userID).UserName

	if comment.UserAvatar == "" {
		comment.UserAvatar = "https://example.com/default-avatar.png" // Default avatar URL
	}
	if comment.UserName == "" {
		comment.UserName = "Anonymous" // Default name if not found
	}

	utils.WriteJSON(w, comment, http.StatusOK)
}

// GetComments retrieves all comments for a specific post
func GetComments(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.WriteJSON(w, map[string]string{"error": "Method not allowed"}, http.StatusMethodNotAllowed)
		return
	}

	// Get post ID from query parameters
	postIdStr := r.URL.Query().Get("postId")
	if postIdStr == "" {
		utils.WriteJSON(w, map[string]string{"error": "Post ID is required"}, http.StatusBadRequest)
		return
	}

	postId, err := strconv.Atoi(postIdStr)
	if err != nil {
		utils.WriteJSON(w, map[string]string{"error": "Invalid post ID"}, http.StatusBadRequest)
		return
	}

	// Get comments from database
	comments, err := models.GetCommentsByPostId(postId)
	if err != nil {
		utils.WriteJSON(w, map[string]string{"error": "Failed to retrieve comments"}, http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, comments, http.StatusOK)
}
