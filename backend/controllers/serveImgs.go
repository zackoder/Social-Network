package controllers

import (
	"net/http"
	"social-network/utils"
	"strings"
)

func HandelPics(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.WriteJSON(w, map[string]string{"error:": "mehtod not allowed"}, http.StatusMethodNotAllowed)
		return
	}

	/*
		cookie, err := r.Cookie("token")
		if err != nil {
			utils.WriteJSON(w, map[string]string{"error:": "Unauthorized"}, http.StatusUnauthorized)
			return
		}
	*/
	path := strings.TrimPrefix(r.URL.Path, "/uploads/")
	if path == "" {
		utils.WriteJSON(w, map[string]string{"error": "forbidden"}, http.StatusForbidden)
		return
	}

	http.ServeFile(w, r, "./uploads/"+path)
}
