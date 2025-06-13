package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"social-network/models"
	"social-network/utils"
)

func Creat_groupe(w http.ResponseWriter, r *http.Request, user_id int) {
	if r.Method != http.MethodPost {
		utils.WriteJSON(w, map[string]string{"error": "Method Not Allowd"}, http.StatusMethodNotAllowed)
		return
	}

	var Groupe utils.Groupe
	err := json.NewDecoder(r.Body).Decode(&Groupe)

	if err != nil {
		utils.WriteJSON(w, map[string]string{"error": "Bad Request"}, http.StatusBadRequest)
		return
	}
	Groupe.CreatorId = user_id

	if len(strings.TrimSpace(Groupe.Title)) < 2 || len(Groupe.Title) > 50 {
		utils.WriteJSON(w, map[string]string{"error": "invalid group title"}, http.StatusBadRequest)
		return
	}

	if len(strings.TrimSpace(Groupe.Description)) < 2 || len(Groupe.Description) > 300 {
		utils.WriteJSON(w, map[string]string{"error": "invalid group discription"}, http.StatusBadRequest)
		return
	}

	Groupe.CreatorId = user_id

	log.Println("group:", Groupe.CreatorId)

	Groupe.Id, err = models.InsertGroupe(Groupe.Title, Groupe.Description, user_id)
	err = models.InsserMemmberInGroupe(Groupe.Id, user_id, "creator")

	if err != nil {

		if strings.Contains(err.Error(), "groups.name") {
			utils.WriteJSON(w, map[string]string{"error": "This group already exists"}, http.StatusBadRequest)
			return
		}
		utils.WriteJSON(w, map[string]string{"error": "Internal Server Error"}, http.StatusInternalServerError)
		return
	}
	// models.InsserMemmberInGroupe(groupe_id, Groupe.CreatorId, "creator")

	utils.WriteJSON(w, Groupe, http.StatusOK)
	return
}

func Join_Group(w http.ResponseWriter, r *http.Request, user_id int) {
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

	requist.User_id = user_id
	if models.IsMember(requist.Groupe_id, requist.User_id) {
		utils.WriteJSON(w, map[string]string{"error": "you are already a member of this group"}, 403)
		return
	}

	if !models.CheckGroup(requist.Groupe_id) {
		utils.WriteJSON(w, map[string]string{"error": "Group not found"}, http.StatusNotFound)
		return
	}

	var noti utils.Notification
	noti.Target_id = models.GetGroupOwner(requist)
	if noti.Target_id == user_id {
		utils.WriteJSON(w, map[string]string{"error": "Forbidden"}, http.StatusForbidden)
		return
	}
	noti.Actor_id = requist.Groupe_id
	noti.Sender_id = requist.User_id
	noti.Message = "join request"
	err = models.InsertNotification(noti)
	if err != nil {
		log.Println("error lksdfjgksdfglkjdsglkjgl", err)
		os.Exit(10)
		if err.Error() != "" && strings.Contains(err.Error(), "FOREIGN KEY") {
			utils.WriteJSON(w, map[string]string{"error": "check your data"}, http.StatusBadRequest)
		} else {
			utils.WriteJSON(w, map[string]string{"error": "Internal Server Error"}, http.StatusInternalServerError)
			log.Println(err)
		}
		return
	}

	// if !models.IsMember(requist.Groupe_id, requist.User_id) {
	// 	err = models.InsserMemmberInGroupe(requist.Groupe_id, requist.User_id, "member")
	// 	if err != nil {
	// 		utils.WriteJSON(w, map[string]string{"error": "Internal Server Error"}, http.StatusInternalServerError)
	// 		log.Println(err)
	// 	}
	// 	utils.WriteJSON(w, map[string]string{"prossotion": "seccesfel"}, http.StatusOK)
	// } else {
	// 	utils.WriteJSON(w, map[string]string{"error": "you are redy member in this group"}, 403)
	// 	return
	// }
	utils.WriteJSON(w, map[string]string{"prossotion": "succeeded"}, http.StatusOK)
}

func Getgroupmsgs(w http.ResponseWriter, r *http.Request, user_id int) {
	group_id, err := strconv.Atoi(r.URL.Query().Get("groupId"))
	offset := r.URL.Query().Get("offset")
	if err != nil {
		utils.WriteJSON(w, map[string]string{"error": "invalid data"}, http.StatusForbidden)
		return
	}
	if !models.IsMember(group_id, user_id) {
		utils.WriteJSON(w, map[string]string{"error": "you can't send messages unless you are a member "}, http.StatusForbidden)
		return
	}
	msgs, err := models.SelectGroupMSGs(group_id, user_id, offset, r.Host)
	if err != nil {
		log.Println("group messages err:", err)
	}
	utils.WriteJSON(w, msgs, 200)
}

func AllGroups(w http.ResponseWriter, r *http.Request, user_id int) {
	if r.Method != http.MethodGet {
		utils.WriteJSON(w, map[string]string{"error": "Method Not Allowd"}, http.StatusMethodNotAllowed)
		return
	}

	Groups := models.GetAllGroups(user_id)
	utils.WriteJSON(w, Groups, 200)
}

func GetGroupsJoined(w http.ResponseWriter, r *http.Request, user_id int) {
	cookie, err := r.Cookie("token")
	if err != nil {
		utils.WriteJSON(w, map[string]string{"error": "You don't have access."}, http.StatusForbidden)
		return
	}
	if r.Method != http.MethodGet {
		utils.WriteJSON(w, map[string]string{"error": "Method Not Allowd"}, http.StatusMethodNotAllowed)
		return
	}

	userId, err := models.Get_session(cookie.Value)
	if err != nil {
		utils.WriteJSON(w, map[string]string{"error": "Invalid session"}, http.StatusUnauthorized)
		return
	}

	Groups := models.GetGroupsOfMember(userId)

	utils.WriteJSON(w, Groups, 200)
}

func GetGroupsCreatedByUser(w http.ResponseWriter, r *http.Request, user_id int) {
	if r.Method != http.MethodGet {
		utils.WriteJSON(w, map[string]string{"error": "Method Not Allowd"}, http.StatusMethodNotAllowed)
		return
	}

	Groups := models.GroupsCreatedByUser(user_id)
	cookie, err := r.Cookie("token")
	if err != nil {
		utils.WriteJSON(w, map[string]string{"error": "You don't have access."}, http.StatusForbidden)
		return
	}
	userId, err := models.Get_session(cookie.Value)
	if err != nil {
		utils.WriteJSON(w, map[string]string{"error": "Invalid session"}, http.StatusUnauthorized)
		return
	}
	Groups = models.GroupsCreatedByUser(userId)

	utils.WriteJSON(w, Groups, 200)
}

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

	Broadcast(noti.Target_id, noti)
	utils.WriteJSON(w, map[string]string{"message": "Invitation sent"}, http.StatusCreated)
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

func CreatEvent(w http.ResponseWriter, r *http.Request, userId int) {
	log.Println(r.Method)
	if r.Method != http.MethodPost {
		log.Println("haaa fin ana ")
		utils.WriteJSON(w, map[string]string{"error": "Method Not allowd"}, http.StatusMethodNotAllowed)
		return
	}

	var notification utils.Notification
	var event utils.Event
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		utils.WriteJSON(w, map[string]string{"error": "Status BadRequest1"}, http.StatusBadRequest)
		return
	}
	notification.Target_id = event.GroupID
	event.CreatedBy = userId
	if len(event.Title) > 25 || len(event.Description) > 100 || len(strings.TrimSpace(event.Description)) < 2 || len(strings.TrimSpace(event.Title)) < 2 {
		utils.WriteJSON(w, map[string]string{"error": "Status BadRequest2"}, http.StatusBadRequest)
		return
	}

	if !models.IsMember(event.GroupID, userId) {
		utils.WriteJSON(w, map[string]string{"error": "Access denied: you must be a member of the group to creat event."}, 403)
		return
	}

	notification.Actor_id, err = models.InsserEventInDatabase(event)
	notification.Message = "event"
	notification.Sender_id = userId
	err = models.InsertNotification(notification)
	if err != nil {
		utils.WriteJSON(w, map[string]string{"error": "Internal Server Error"}, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "The event cried out successfully"})
}

func EventResponce(w http.ResponseWriter, r *http.Request, userId int) {
	if r.Method != http.MethodPost {
		utils.WriteJSON(w, map[string]string{"error": "Method Not allowd"}, http.StatusMethodNotAllowed)
		return
	}
	var responce utils.EventResponse
	if err := json.NewDecoder(r.Body).Decode(&responce); err != nil {
		utils.WriteJSON(w, map[string]string{"error": "Status BadRequest 3"}, http.StatusBadRequest)
		return
	}

	responce.UserID = userId

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

func GetPostsGroupe(w http.ResponseWriter, r *http.Request, userId int) {
	if r.Method != http.MethodPost {
		utils.WriteJSON(w, map[string]string{"error": "Method Not allowd"}, http.StatusMethodNotAllowed)
		return
	}
	type RequestDAta struct {
		Groupe_id int `json:"id"`
	}
	var Groupe_id RequestDAta
	err := json.NewDecoder(r.Body).Decode(&Groupe_id)
	if err != nil {
		utils.WriteJSON(w, map[string]string{"error": "Status BadRequest 3"}, http.StatusBadRequest)
		return
	}
	if !models.IsMember(Groupe_id.Groupe_id, userId) {
		utils.WriteJSON(w, map[string]string{"error": "Access denied: you must be a member of the group to Fetchs Posts."}, 403)
		return
	}
	Posts, err := models.GetPostsFromDatabase(Groupe_id.Groupe_id, r)
	fmt.Println(err)
	if err != nil {
		utils.WriteJSON(w, map[string]string{"error": "Internal Server Error."}, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(Posts)
}

func GetEvents(w http.ResponseWriter, r *http.Request, userId int) {
	if r.Method != http.MethodPost {
		utils.WriteJSON(w, map[string]string{"error": "Method Not allowd"}, http.StatusMethodNotAllowed)
		return
	}
	var groupeId int
	err := json.NewDecoder(r.Body).Decode(&groupeId)
	if err != nil {
		utils.WriteJSON(w, map[string]string{"error": "Status BadRequest 3"}, http.StatusBadRequest)
		return
	}
	if !models.IsMember(groupeId, userId) {
		utils.WriteJSON(w, map[string]string{"error": "Access denied: you must be a member of the group to Fetchs Events."}, 403)
		return
	}
	Events, err := models.GetEventsFromDatabase(groupeId, userId)
	if err != nil {
		utils.WriteJSON(w, map[string]string{"error": "Internal Server Error."}, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(Events)
}

func GetUserIdByCookie(r *http.Request) (int, error) {
	cookie, err0 := r.Cookie("token")
	if err0 != nil {
		return 0, err0
	}
	userId, err := models.Get_session(cookie.Value)
	if err != nil {
		return 0, err
	}
	return userId, nil
}

func GetGroup(w http.ResponseWriter, r *http.Request, user_id int) {
	group_id, err := strconv.Atoi(r.URL.Query().Get("groupId"))
	if err != nil {
		utils.WriteJSON(w, map[string]string{"error": "invalid data"}, http.StatusNotFound)
		return
	}
	group, err := models.GetOneGroup(group_id)

	group.Id = group_id
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("hhhhh khawi")
			utils.WriteJSON(w, map[string]string{"error": "This group does not exist!"}, http.StatusBadRequest)
		} else {
			utils.WriteJSON(w, map[string]string{"error": "Internal Server Error"}, http.StatusInternalServerError)
			log.Println("scan group error", err)
		}
		return
	}
	log.Println(group_id, "userid", group)
	utils.WriteJSON(w, group, 200)
}
