package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"social-network/models"
	"social-network/utils"
	"strconv"
)




func GetRegistrationData(w http.ResponseWriter, r *http.Request,userID int)  {
	profileOwnerIDStr := r.URL.Query().Get("id")
	profileId := strconv.Itoa(userID)
	registrationData , err := models.GetRegistration(profileOwnerIDStr) 
	if err != nil {
		if err == sql.ErrNoRows {
			utils.WriteJSON(w, map[string]string{"error": "User Not Found"},http.StatusNotFound)
			}else{
				fmt.Println("im here",err)
				utils.WriteJSON(w, map[string]string{"error": "User Not Found"},http.StatusInternalServerError)
			}
			return
		}
	registrationData.ProfileOner =	profileOwnerIDStr == profileId
	if registrationData.Avatar != ""{
		registrationData.Avatar = r.Host + registrationData.Avatar
	}
	utils.WriteJSON(w,registrationData,200)
}