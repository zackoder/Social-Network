package controllers

import (
	"net/http"
	"social-network/models"
	"social-network/utils"
	"strconv"
)



func GetFollowers(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET"{
		utils.WriteJSON(w,map[string]string{"error":"Method Not Allowed"},http.StatusMethodNotAllowed)
		return
	}
	profileOwnerIDStr := r.URL.Query().Get("id")
	profileOwnerID,err := strconv.Atoi(profileOwnerIDStr)
	if err != nil {
		utils.WriteJSON(w,map[string]string{"error":"Internal Server Error"},http.StatusInternalServerError)
		return
	}
 	followers,err := models.GetFollowers(profileOwnerID)
	if err != nil {
		utils.WriteJSON(w,map[string]string{"error":"We Can't Fetsh Followers List"},http.StatusInternalServerError)
		return
	}
	utils.WriteJSON(w, followers, http.StatusAccepted)

}