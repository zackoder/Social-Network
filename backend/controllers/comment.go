package controllers

import (
	"encoding/json"
	"net/http"
	"social-network/models"
	"social-network/utils"
	"strconv"
)

// AddComment handles adding a new comment to a post
func AddComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriteJSON(w, map[string]string{"error": "Method not allowed"}, http.StatusMethodNotAllowed)
		return
	}

	// Get user session to identify the commenter
	cookie, err := r.Cookie("token")
	if err != nil {
		utils.WriteJSON(w, map[string]string{"error": "Authentication required"}, http.StatusUnauthorized)
		return
	}

	userId, err := models.Get_session(cookie.Value)
	if err != nil {
		utils.WriteJSON(w, map[string]string{"error": "Invalid session"}, http.StatusUnauthorized)
		return
	}

	// Parse the comment data
	var comment utils.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		utils.WriteJSON(w, map[string]string{"error": "Invalid request data"}, http.StatusBadRequest)
		return
	}

	// Set the user ID from the session
	comment.UserId = userId

	// Check if post exists
	postExists, err := models.PostExists(comment.PostId)
	if err != nil || !postExists {
		utils.WriteJSON(w, map[string]string{"error": "Post not found"}, http.StatusNotFound)
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
	user, err := models.GetUserById(userId)
	if err == nil {
		comment.UserName = user.FirstName + " " + user.LastName
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
