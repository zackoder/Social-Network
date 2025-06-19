package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"social-network/models"
	"social-network/utils"
)

func GetNotifications(w http.ResponseWriter, r *http.Request, userId int) {
	notifications, err := models.SelectNotifications(userId)
	if err != nil {
		log.Println(err)
	}
	utils.WriteJSON(w, notifications, 200)
}

func NotiResp(w http.ResponseWriter, r *http.Request, userId int) {
	var noti utils.Notification
	err := json.NewDecoder(r.Body).Decode(&noti)
	if err != nil {
		log.Println("decoding json11", err)
		return
	}
	resp := noti.Message
	println("im the resp", noti.Message)
	models.SelectOneNoti(&noti)
	if resp == "rejected" || resp == "accepted" {
		if resp == "rejected" {
			err = models.DeleteNoti(noti.Id)
			if err != nil {
				if err == sql.ErrNoRows {
					utils.WriteJSON(w, map[string]string{"error": "there is no such notification"}, http.StatusBadRequest)
					return
				}
				log.Println(err)
			}
			return
		}
		if noti.Message == "group invitation" {
			err := models.InsserMemmberInGroupe(noti.Actor_id, noti.Target_id, "member")
			if err != nil {
				if strings.Contains(err.Error(), "UNIQUE") {
					utils.WriteJSON(w, map[string]string{"error": "you are already a member of that group"}, http.StatusBadRequest)
				} else {
					utils.WriteJSON(w, map[string]string{"error": "internal server error"}, http.StatusInternalServerError)
					log.Println("inserting member error", err)
					return
				}
			}
		} else if noti.Message == "follow request" {
			err := models.InsertFollow(noti.Actor_id, strconv.Itoa(noti.Target_id))
			if err != nil {
				log.Println(err)
			}
		} else if noti.Message == "join request" {
			err := models.InsserMemmberInGroupe(noti.Actor_id, noti.Sender_id, "member")
			if err != nil {
				if strings.Contains(err.Error(), "UNIQUE") {
					utils.WriteJSON(w, map[string]string{"error": "you are already a member of that group"}, http.StatusBadRequest)
				} else {
					utils.WriteJSON(w, map[string]string{"error": "internal server error"}, http.StatusInternalServerError)
					log.Println("inserting member error", err)
					return
				}
			}
		}
		models.DeleteNoti(noti.Id)
	} else {
		utils.WriteJSON(w, map[string]string{"error": "Bad Request"}, http.StatusBadRequest)
		return
	}
	utils.WriteJSON(w, map[string]string{"successful": "ok"}, 200)
}
