package controllers

import (
	"encoding/json"
	"net/http"
	"social-network/models"
	"social-network/utils"
	"strconv"
)

// AddReaction handles adding a new reaction to a post
func AddReaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriteJSON(w, map[string]string{"error": "Method not allowed"}, http.StatusMethodNotAllowed)
		return
	}

	// Get user session to identify who is reacting
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

	// Parse the reaction data
	var reaction utils.Reaction
	if err := json.NewDecoder(r.Body).Decode(&reaction); err != nil {
		utils.WriteJSON(w, map[string]string{"error": "Invalid request data"}, http.StatusBadRequest)
		return
	}

	// Set the user ID from the session
	reaction.UserId = userId

	// Check if post exists
	postExists, err := models.PostExists(reaction.PostId)
	if err != nil || !postExists {
		utils.WriteJSON(w, map[string]string{"error": "Post not found"}, http.StatusNotFound)
		return
	}

	// Check if the user has already reacted to this post
	existingReaction, err := models.GetUserReactionForPost(userId, reaction.PostId)
	if err == nil && existingReaction != nil {
		// User has already reacted, so update the reaction type if it's different
		if existingReaction.ReactionType != reaction.ReactionType {
			if err := models.UpdateReaction(userId, reaction.PostId, reaction.ReactionType); err != nil {
				utils.WriteJSON(w, map[string]string{"error": "Failed to update reaction"}, http.StatusInternalServerError)
				return
			}
			utils.WriteJSON(w, map[string]string{"message": "Reaction updated", "type": reaction.ReactionType}, http.StatusOK)
			return
		}
		// If the reaction is the same, remove it (toggle)
		if err := models.DeleteReaction(userId, reaction.PostId); err != nil {
			utils.WriteJSON(w, map[string]string{"error": "Failed to remove reaction"}, http.StatusInternalServerError)
			return
		}
		utils.WriteJSON(w, map[string]string{"message": "Reaction removed"}, http.StatusOK)
		return
	}

	// Insert new reaction
	reactionID, err := models.InsertReaction(&reaction)
	if err != nil {
		utils.WriteJSON(w, map[string]string{"error": "Failed to save reaction"}, http.StatusInternalServerError)
		return
	}

	// Set the generated ID and return the reaction
	reaction.Id = reactionID

	utils.WriteJSON(w, reaction, http.StatusOK)
}

// GetReactions retrieves all reactions for a specific post
func GetReactions(w http.ResponseWriter, r *http.Request) {
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

	// Get reactions from database
	reactions, err := models.GetReactionsByPostId(postId)
	if err != nil {
		utils.WriteJSON(w, map[string]string{"error": "Failed to retrieve reactions"}, http.StatusInternalServerError)
		return
	}

	// Get current user's reaction, if any
	cookie, _ := r.Cookie("token")
	var currentUserReaction *utils.Reaction
	
	if cookie != nil {
		userId, err := models.Get_session(cookie.Value)
		if err == nil {
			currentUserReaction, _ = models.GetUserReactionForPost(userId, postId)
		}
	}

	// Create response with reaction counts and current user's reaction
	response := map[string]interface{}{
		"reactions": reactions,
		"counts":    models.CountReactionsByType(reactions),
		"userReaction": currentUserReaction,
	}

	utils.WriteJSON(w, response, http.StatusOK)
}

// DeleteReaction handles removing a user's reaction from a post
func DeleteReaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		utils.WriteJSON(w, map[string]string{"error": "Method not allowed"}, http.StatusMethodNotAllowed)
		return
	}

	// Get user session to verify ownership
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

	// Delete the reaction
	if err := models.DeleteReaction(userId, postId); err != nil {
		utils.WriteJSON(w, map[string]string{"error": "Failed to delete reaction"}, http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, map[string]string{"message": "Reaction removed successfully"}, http.StatusOK)
}
