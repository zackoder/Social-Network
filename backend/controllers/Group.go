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

func Group(w http.ResponseWriter, r *http.Request) {
}

func Creat_groupe(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriteJSON(w, map[string]string{"error": "Method Not Allowd"}, http.StatusMethodNotAllowed)
		return
	}

	var Groupe utils.Groupe
	err := json.NewDecoder(r.Body).Decode(&Groupe)
	if err != nil {
		fmt.Println(err)
		utils.WriteJSON(w, map[string]string{"error": "Bad Request"}, http.StatusBadRequest)
		return
	}

	if len(strings.TrimSpace(Groupe.Title)) < 2 || len(Groupe.Title) > 50 {
		utils.WriteJSON(w, map[string]string{"error": "invalid group title"}, http.StatusBadRequest)
		return
	}

	if len(strings.TrimSpace(Groupe.Description)) < 2 || len(Groupe.Description) > 300 {
		utils.WriteJSON(w, map[string]string{"error": "invalid group discription"}, http.StatusBadRequest)
		return
	}

	groupInserted, err := models.InsserGroupe(Groupe.Title, Groupe.Description, Groupe.CreatorId)
	if err != nil {
		fmt.Println("inserting group err", err)
		utils.WriteJSON(w, map[string]string{"error": "Internal Server Error"}, http.StatusInternalServerError)
		return
	}
	models.InsserMemmberInGroupe(groupInserted, Groupe.CreatorId, "creator")

	utils.WriteJSON(w, map[string]string{"Groupe": "criete groupe seccesfel"}, http.StatusOK)
	return
}

func Jouind_Groupe(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriteJSON(w, map[string]string{"error": "Method Not Allowd"}, http.StatusMethodNotAllowed)
		return
	}

	var requist utils.Groupe_member
	err := json.NewDecoder(r.Body).Decode(&requist)
	if err != nil {
		utils.WriteJSON(w, map[string]string{"error": "Bad Request"}, http.StatusBadRequest)
		return
	}

	if models.IsMember(requist.Groupe_id, requist.User_id) {
		utils.WriteJSON(w, map[string]string{"error": "you are already a member of this group"}, 403)
		return
	}

	err = models.InsserMemmberInGroupe(requist.Groupe_id, requist.User_id, "member")
	if err != nil {
		if strings.Contains(err.Error(), "FOREIGN KEY") {
			utils.WriteJSON(w, map[string]string{"error": "Check your data"}, http.StatusBadRequest)
			return
		}
		log.Println("inserting a member to a group", err)
		utils.WriteJSON(w, map[string]string{"error": "Internal Server Error"}, http.StatusInternalServerError)
		return
	}
	utils.WriteJSON(w, map[string]string{"prossotion": "seccesfel"}, http.StatusOK)
}

// fetch('/api/searchGroups?query=tech')
//   .then(response => response.json())
//   .then(data => {
//     console.log(data);
//     // Met à jour l'UI avec les résultats
//   });

func SearchGroupsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.WriteJSON(w, map[string]string{"error": "Method Not Allowd"}, http.StatusMethodNotAllowed)
		return
	}
	query := r.URL.Query().Get("query")
	groups, err := models.SearchGroupsInDatabase(query)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(groups)
}

func InviteUser(w http.ResponseWriter, r *http.Request /* , groupID uint */) {
	if r.Method != http.MethodPost {
		utils.WriteJSON(w, map[string]string{"error": "Method Not Allowd"}, http.StatusMethodNotAllowed)
		return
	}
	// var invitaion utils.GroupInvitation
	var noti utils.Notification
	if err := json.NewDecoder(r.Body).Decode(&noti); err != nil {
		utils.WriteJSON(w, map[string]string{"error": "Status BadRequest"}, http.StatusBadRequest)
		return
	}

	if !models.IsMember(noti.Actor_id, noti.Sender_id) {
		utils.WriteJSON(w, map[string]string{"error": "you are not a member of the group"}, http.StatusBadRequest)
		return
	}

	if models.IsMember(noti.Actor_id, noti.Target_id) {
		utils.WriteJSON(w, map[string]string{"error": "already a group member"}, 409)
		return
	}
	noti.Message = "group invitation"
	err := models.InsertNotification(noti)

	// err := models.SaveInvitation(invitaion.GroupID, invitaion.InvitedBy, invitaion.UserId)
	if err != nil {
		log.Println("saving invitation", err)
		utils.WriteJSON(w, map[string]string{"error": "Internal Server Error"}, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Invitation sent"})
}

func InsertToGroupe(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriteJSON(w, map[string]string{"error": "Method Not Allowd"}, http.StatusMethodNotAllowed)
		return
	}
}

func Get_all_post(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriteJSON(w, map[string]string{"error": "Method Not Allowd"}, http.StatusMethodNotAllowed)
		return
	}
	groupe_id_str := r.FormValue("groupe_id")
	user_id_str := r.FormValue("user_id")
	groupe_id, err1 := strconv.Atoi(groupe_id_str)
	user_id, err2 := strconv.Atoi(user_id_str)

	if err1 != nil || err2 != nil {

		utils.WriteJSON(w, map[string]string{"error": "Status BadRequest"}, http.StatusBadRequest)
		return

	}

	if !models.IsMember(groupe_id, user_id) {
		utils.WriteJSON(w, map[string]string{"error": "Access denied: you must be a member of the group to view posts."}, 403)
		return

	}

	Posts := models.QueryPosts(0, r)
	var Posts_groupe []utils.Post
	for _, v := range Posts {
		if v.Groupe_id == groupe_id {
			Posts_groupe = append(Posts_groupe, v)
		}
	}

	utils.WriteJSON(w, Posts_groupe, http.StatusOK)
}

func CreatEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriteJSON(w, map[string]string{"error": "Method Not allowd"}, http.StatusMethodNotAllowed)
		return
	}
	var event utils.Event
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		utils.WriteJSON(w, map[string]string{"error": "Status BadRequest"}, http.StatusBadRequest)
		return
	}

	if len(event.Title) > 25 || len(event.Description) > 100 || len(strings.TrimSpace(event.Description)) < 2 || len(strings.TrimSpace(event.Title)) < 2 {
		utils.WriteJSON(w, map[string]string{"error": "Status BadRequest"}, http.StatusBadRequest)
		return
	}
	if !models.IsMember(event.GroupID, event.CreatedBy) {
		utils.WriteJSON(w, map[string]string{"error": "Access denied: you must be a member of the group to creat event."}, 403)
		return
	}
	var notification utils.Notification
	err := models.InsserEventInDatabase(event)
	notification.Message = "join group request"
	notification.Actor_id = 5
	err = models.InsertNotification(notification)
	if err != nil {
		utils.WriteJSON(w, map[string]string{"error": "Internal Server Error"}, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "The event cried out successfully"})
}

func EventRrspponce(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriteJSON(w, map[string]string{"error": "Method Not allowd"}, http.StatusMethodNotAllowed)
		return
	}
	var responce utils.EventResponse
	if err := json.NewDecoder(r.Body).Decode(&responce); err != nil {
		utils.WriteJSON(w, map[string]string{"error": "Status BadRequest"}, http.StatusBadRequest)
		return
	}

	if !models.IsMember(responce.GroupeId, responce.UserID) {
		utils.WriteJSON(w, map[string]string{"error": "Access denied: you must be a member of the group to creat event."}, 403)
		return
	}
	err := models.InsserResponceInDatabase(responce)
	if err != nil {
		utils.WriteJSON(w, map[string]string{"error": "Internal Server Error"}, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "The answer was successfully added"})
}

func Event(noti utils.Notification) {

}
