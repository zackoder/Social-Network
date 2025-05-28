package controllers

import (
	"net/http"
	"strconv"

	"social-network/models"
	"social-network/utils"
)

func GetMessages(w http.ResponseWriter, r *http.Request, user_id int) {
	receiver_id, err := strconv.Atoi(r.URL.Query().Get("receiver_id"))
	offset := r.URL.Query().Get("offset")
	if err != nil {
		utils.WriteJSON(w, map[string]string{"error": "invalide data"}, http.StatusForbidden)
		return
	}
	models.QueryMsgs(receiver_id, offset)
}
