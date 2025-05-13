package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"social-network/models"
	"social-network/utils"
	"strings"
)

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriteJSON(w, map[string]string{"error": "Method Not Allowd"}, http.StatusMethodNotAllowed)
		return
	}
	// max image size 10Mb
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		utils.WriteJSON(w, map[string]string{"error": "Too larg file"}, http.StatusRequestEntityTooLarge)
		return
	}
	
	userData := r.FormValue("userData")
	filePath, err := utils.UploadImage(r)
	if err != nil {
		utils.WriteJSON(w, map[string]string{"error": "Internal Server Error"}, http.StatusMethodNotAllowed)
		return
	}
	if len(userData) == 0 {
		utils.WriteJSON(w, map[string]string{"error": "Empty request body"}, http.StatusBadRequest)
		return
	}


	var regesterreq utils.Regester
	if err := json.Unmarshal([]byte(userData), &regesterreq); err != nil {
		utils.WriteJSON(w, map[string]string{"error": "Internal Server Error1"}, http.StatusMethodNotAllowed)
		fmt.Println(err)
		return
	}

	regesterreq.Avatar = filePath
	if regesterreq.NickName == "" {
		regesterreq.NickName = nil
	}
	fmt.Println(regesterreq.NickName)
	
	if utils.ValidatNames(regesterreq.FirstName, regesterreq.LastName, regesterreq.NickName) && utils.ValidEmail(regesterreq.Email) && regesterreq.Password == regesterreq.ConfermPassword{
		hashedPss := utils.Hashpass(regesterreq.Password)
		if hashedPss == "" {
			utils.WriteJSON(w, map[string]string{"error": "Internal Server Error"}, http.StatusInternalServerError)
			return
		}
		regesterreq.Password = hashedPss

		if err := models.InsertUser(regesterreq); err != nil {
			if strings.Contains(err.Error(), "users.email") {
				utils.WriteJSON(w, map[string]string{"error": "Email has already been taken."}, http.StatusNotAcceptable)
			} else if strings.Contains(err.Error(), "users.nickname") {
				utils.WriteJSON(w, map[string]string{"error": "nickname has already been taken."}, http.StatusNotAcceptable)
			} else {
				utils.WriteJSON(w, map[string]string{"error": "Fail to register user"}, http.StatusNotAcceptable)
			}
			if err := os.Remove(filePath); err != nil {
				fmt.Println("removing error:", err)
			}
			return
		}
	}
	utils.WriteJSON(w, map[string]string{"success": "ok"}, http.StatusOK)
}