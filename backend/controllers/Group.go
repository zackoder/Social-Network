package controllers

import (
	"encoding/json"
	"net/http"
	"social-network/models"
	"social-network/utils"
	"strconv"
	"strings"
)

func Group(w http.ResponseWriter, r *http.Request) {

}

func CreateGroup(w http.ResponseWriter, r *http.Request) {
	var group utils.NewGroup
	id := 5
	json.NewDecoder(r.Body).Decode(&group)
	err := models.InsertNewGroup(&group, id)
	if err != nil {
		if strings.Contains(err.Error(), "groups.name") {
			utils.WriteJSON(w, map[string]string{"err": "chouse anoter name for the group"}, http.StatusForbidden)
			return
		}
		utils.WriteJSON(w, map[string]string{"err": "couldent insert group"}, http.StatusInternalServerError)
		return
	}
	models.InsertMumber(group.Id, id)
	Manager.AddGroup(group.Id, id)
	utils.WriteJSON(w, group, 200)
}

func JoinReq(w http.ResponseWriter, r *http.Request) {
	group_id, err := strconv.Atoi(r.URL.Query().Get("group_id"))
	if err != nil {
		utils.WriteJSON(w, map[string]string{"error": "invalid group id"}, http.StatusNotFound)
		return
	}
	
	user_id, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		utils.WriteJSON(w, map[string]string{"error": "invalid user id"}, http.StatusNotFound)
		return
	}

	if err := models.InsertMumber(group_id, user_id); err != nil {
		if strings.Contains(err.Error(), "UNIQUE") {
			utils.WriteJSON(w, map[string]string{"error": "mumber alredy exists"}, http.StatusForbidden)
		} else if strings.Contains(err.Error(), "FOREIGN KEY") {
			utils.WriteJSON(w, map[string]string{"error": "check the provided data"}, http.StatusNotFound)
		} else {
			utils.WriteJSON(w, map[string]string{"error": "internal server error"}, http.StatusInternalServerError)
		}
		return
	}
	utils.WriteJSON(w, map[string]string{"successe": "mumber added"}, http.StatusOK)
}
