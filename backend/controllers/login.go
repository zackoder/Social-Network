package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	models "social-network/models"
	utils "social-network/utils"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var userData utils.User
	if err := json.NewDecoder(r.Body).Decode(&userData); err != nil {
		utils.WriteJSON(w, "invalid input data", http.StatusBadRequest)
		return
	}

	// if len(userData.Email) < 5 || len(userData.Password) < 5 || len(userData.Email) > 250 || len(userData.Password) > 64 {
	// 	utils.WriteJSON(w, map[string]string{"error": "invalid username/password/email"}, http.StatusBadRequest)
	// 	return
	// }

	// if utils.IsValidEmail(&userData.Email) {
	// 	 userData.Email = userData.Email, userData.Email
	// }

	password := userData.Password
	err := models.ValidCredential(&userData)
	if err != nil {
		if err == sql.ErrNoRows {
			utils.WriteJSON(w, map[string]string{"error": "Incorect Username or password"}, http.StatusUnauthorized)
			return
		}
		fmt.Println(err)
		utils.WriteJSON(w, map[string]string{"error": "internaInternal Server Error1"}, http.StatusInternalServerError)
		return
	}

	if userData.ID > 10 {
		if !utils.CheckPasswordHash(&password, &userData.Password) {
			// fmt.Println(&password,&userData.Password)
			utils.WriteJSON(w, "Incorect password", http.StatusUnauthorized)
			return
		}
	}

	userData.SessionId, err = utils.GenerateSessionID()
	if err != nil {
		fmt.Println(err)
		utils.WriteJSON(w, map[string]string{"error": "internaInternal Server Error"}, http.StatusInternalServerError)
		return
	}
	userData.SessionId, err = utils.GenerateSessionID()
	if err != nil {
		fmt.Println(err)
		utils.WriteJSON(w, map[string]string{"error": "internaInternal Server Error"}, http.StatusInternalServerError)
		return
	}

	err = models.InsertSession(&userData)
	if err != nil {
		fmt.Println(err)
		utils.WriteJSON(w, map[string]string{"error": "internaInternal Server Error"}, http.StatusInternalServerError)
		return
	}
	
	http.SetCookie(w, &http.Cookie{
		Name:  "token",
		Path:  "/",
		Value: userData.SessionId,
	})

	utils.WriteJSON(w, map[string]string{"success": "ok"}, http.StatusOK)
}
