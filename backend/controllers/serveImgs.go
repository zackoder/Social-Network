package controllers

import (
	"net/http"
	"strings"

	"social-network/utils"
)

func HandelPics(w http.ResponseWriter, r *http.Request, user int) {
	if r.Method != http.MethodGet {
		utils.WriteJSON(w, map[string]string{"error:": "mehtod not allowed"}, http.StatusMethodNotAllowed)
		return
	}
	
	path := r.URL.Path
	validpath := strings.TrimPrefix(r.URL.Path, "/uploads/")
	validpath = strings.TrimPrefix(r.URL.Path, "/defaultIMG/")
	if validpath == "" {
		utils.WriteJSON(w, map[string]string{"error": "Forbidden"}, http.StatusForbidden)
		return
	}

	http.ServeFile(w, r, "."+path)
}
