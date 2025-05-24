package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"social-network/models"
	"social-network/utils"
	"strconv"
)

// GetNotifications fetches notifications for a user with pagination support
func GetNotifications(w http.ResponseWriter, r *http.Request, userID int) {

	// Parse pagination parameters
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")

	// Default values
	limit := 5
	offset := 0

	// Parse limit if provided
	if limitStr != "" {
		parsedLimit, err := strconv.Atoi(limitStr)
		if err == nil && parsedLimit > 0 {
			limit = parsedLimit
		}
	}

	// Parse offset if provided
	if offsetStr != "" {
		parsedOffset, err := strconv.Atoi(offsetStr)
		if err == nil && parsedOffset >= 0 {
			offset = parsedOffset
		}
	}

	// Get notifications from the database
	notifications, err := models.GetNotifications(userID, limit, offset)
	if err != nil {
		fmt.Println("Error fetching notifications:", err)
		http.Error(w, "Failed to fetch notifications", http.StatusInternalServerError)
		return
	}

	// Set content type and serialize the response
	w.Header().Set("Content-Type", "application/json")

	response := struct {
		Notifications []utils.Notification `json:"notifications"`
		HasMore       bool                 `json:"hasMore"`
	}{
		Notifications: notifications,
		HasMore:       len(notifications) == limit, // If we got exactly the requested limit, there may be more
	}

	json.NewEncoder(w).Encode(response)
}
