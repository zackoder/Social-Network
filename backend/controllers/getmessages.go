package controllers

import (
	"log"
	"net/http"
	"strconv"

	"social-network/models"
	"social-network/utils"
)

func GetMessages(w http.ResponseWriter, r *http.Request, user_id int) {
	var message utils.Message
	var err error
	message.Reciever_id, err = strconv.Atoi(r.URL.Query().Get("receiver_id"))
	offset := r.URL.Query().Get("offset")
	if err != nil {
		utils.WriteJSON(w, map[string]string{"error": "invalide data"}, http.StatusForbidden)
		return
	}
	message.Sender_id = user_id

	messages, err := models.QueryMsgs(message, r.Host, offset)
	if err != nil {
		log.Println("query error:", err)
		utils.WriteJSON(w, map[string]string{"error": "Internal Server Error11"}, http.StatusInternalServerError)
		return
	}
	utils.WriteJSON(w, messages, 200)
}
