package controllers

import (
	"net/http"
	"social-network/models"
	"social-network/utils"
)

func HandleFollow(w http.ResponseWriter, r *http.Request) {
	follower := r.URL.Query().Get("follower")
	followed := r.URL.Query().Get("followed")
	privacy,err := models.GetProfilePrivecy(followed) 
	if err != nil{
		utils.WriteJSON(w, map[string]string{"resp": "user not found"}, 404)
		return	
	}
	if privacy == "public" {

		resp, err := models.InserOrUpdate(follower, followed)
		if err != nil {
			utils.WriteJSON(w, map[string]string{"resp": "try to follow net time"}, 404)
			return
		}
		utils.WriteJSON(w, map[string]string{"resp": resp}, 200)
	}else if privacy == "privat" {
		
	}
}

func UpdatePrivacy(w http.ResponseWriter, r *http.Request) {
	id := 2
	test := models.UpdateProfile(id)
	utils.WriteJSON(w, map[string]string{"test": test}, 200)
}
