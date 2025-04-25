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
	fmt.Println("test-------------------",userData.Email)
	fmt.Println("test", userData.Password)

	// if len(userData.Email) < 4 || len(userData.Password) < 5 || len(userData.Email) > 30 || len(userData.Password) > 64 {
	// 	utils.WriteJSON(w, "invalid username/password/email", http.StatusBadRequest)
	// 	return
	// }

	// if utils.IsValidEmail(&userData.Email) {
	// 	 userData.Email = userData.Email, userData.Email
	// }

	password := userData.Password
	err := models.ValidCredential(&userData)
	if err != nil {
		if err == sql.ErrNoRows {
			utils.WriteJSON(w, "Incorect Username or password", http.StatusUnauthorized)
			return
		}
		fmt.Println(err)
		utils.WriteJSON(w, "internaInternal Server Error1", http.StatusInternalServerError)
		return
	}

	if !utils.CheckPasswordHash(&password, &userData.Password) {
		utils.WriteJSON(w, "Incorect password", http.StatusUnauthorized)
		return
	}
	ok, err := models.GetActiveSession(&userData)
	if err != nil {
		utils.WriteJSON(w, "internaInternal Server Error2", http.StatusInternalServerError)
		return
	}
	fmt.Println(ok)
	if ok {
		err = models.DeleteSession(&userData)
		if err != nil {
			fmt.Println(err)
			utils.WriteJSON(w, "internaInternal Server Error3", http.StatusInternalServerError)
			return
		}
	}

	userData.SessionId, err = utils.GenerateSessionID()
	if err != nil {
		fmt.Println(err)
		utils.WriteJSON(w, "internaInternal Server Error4", http.StatusInternalServerError)
		return
	}

	err = models.InsertSession(&userData)
	if err != nil {
		fmt.Println(err)
		utils.WriteJSON(w, "internaInternal Server Erro5", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "token",
		Path:  "/",
		Value: userData.SessionId,
	})

	utils.WriteJSON(w, map[string]string{"success": "ok"}, http.StatusOK)
}
