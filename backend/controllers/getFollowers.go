package controllers

import (
	"net/http"

	"social-network/models"
	"social-network/utils"
)

func GetFollowers(w http.ResponseWriter, r *http.Request, userID int) {
	if r.Method != "GET" {
		utils.WriteJSON(w, map[string]string{"error": "Method Not Allowed"}, http.StatusMethodNotAllowed)
		return
	}
	// profileOwnerIDStr := r.URL.Query().Get("id")
	// profileOwnerID,err := strconv.Atoi(profileOwnerIDStr)
	// if err != nil {
	// 	utils.WriteJSON(w,map[string]string{"error":"Internal Server Error"},http.StatusInternalServerError)
	// 	return
	// }
	followers, err := models.GetFollowers(userID)
	if err != nil {
		utils.WriteJSON(w, map[string]string{"error": "We Can't Fetsh Followers List"}, http.StatusInternalServerError)
		return
	}
	utils.WriteJSON(w, followers, http.StatusAccepted)
}
