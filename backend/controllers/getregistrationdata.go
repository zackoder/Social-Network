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
	Response := map[string]any{}
	profileOwnerIDStr := r.URL.Query().Get("id")
	profileOwnerInt, err := strconv.Atoi(profileOwnerIDStr)
	if err != nil{
		utils.WriteJSON(w, map[string]string{"error":"Bad Request1111"},http.StatusBadRequest)
		return
	}
	
	profileId := strconv.Itoa(userID)
	registrationData , err := models.GetRegistration(profileOwnerIDStr) 
	profileStasus,err:= models.GetProfileStatus(profileOwnerInt, userID)
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
	Response["registration_data"] = registrationData 
	Response["profile_status"] = profileStasus 
	utils.WriteJSON(w,Response,200)
}