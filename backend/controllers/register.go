// package controllers

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"os"
// 	"social-network/models"
// 	"social-network/utils"
// 	"strings"
// )

// func Register(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
// 		utils.WriteJSON(w, map[string]string{"error": "Method Not Allowd"}, http.StatusMethodNotAllowed)
// 		return
// 	}
// 	// max image size 10Mb
// 	if err := r.ParseMultipartForm(10 << 20); err != nil {
// 		utils.WriteJSON(w, map[string]string{"error": "Too larg file"}, http.StatusRequestEntityTooLarge)
// 		return
// 	}

// 	userData := r.FormValue("userData")
// 	fmt.Println(userData)
// 	filePath, err := utils.UploadImage(r)
// 	if err != nil {
// 		utils.WriteJSON(w, map[string]string{"error": "Internal Server Error"}, http.StatusMethodNotAllowed)
// 		return
// 	}
// 	if len(userData) == 0 {
// 		utils.WriteJSON(w, map[string]string{"error": "Empty request body"}, http.StatusBadRequest)
// 		return
// 	}

// 	var regesterreq utils.Regester
// 	if err := json.Unmarshal([]byte(userData), &regesterreq); err != nil {
// 		utils.WriteJSON(w, map[string]string{"error": "Internal Server Error1"}, http.StatusMethodNotAllowed)
// 		fmt.Println(err)
// 		return
// 	}

// 	regesterreq.Avatar = filePath
// 	if regesterreq.NickName == "" {
// 		regesterreq.NickName = nil
// 	}
// 	fmt.Println(regesterreq.NickName)

// 	if utils.ValidatNames(regesterreq.FirstName, regesterreq.LastName, regesterreq.NickName) && utils.ValidEmail(regesterreq.Email) && regesterreq.Password == regesterreq.ConfermPassword{
// 		hashedPss := utils.Hashpass(regesterreq.Password)
// 		if hashedPss == "" {
// 			utils.WriteJSON(w, map[string]string{"error": "Internal Server Error"}, http.StatusInternalServerError)
// 			return
// 		}
// 		regesterreq.Password = hashedPss

// 		if err := models.InsertUser(regesterreq); err != nil {
// 			if strings.Contains(err.Error(), "users.email") {
// 				utils.WriteJSON(w, map[string]string{"error": "Email has already been taken."}, http.StatusNotAcceptable)
// 			} else if strings.Contains(err.Error(), "users.nickname") {
// 				utils.WriteJSON(w, map[string]string{"error": "nickname has already been taken."}, http.StatusNotAcceptable)
// 			} else {
// 				utils.WriteJSON(w, map[string]string{"error": "Fail to register user"}, http.StatusNotAcceptable)
// 			}
// 			if err := os.Remove(filePath); err != nil {
// 				fmt.Println("removing error:", err)
// 			}
// 			return
// 		}
// 	}
// 	utils.WriteJSON(w, map[string]string{"success": "ok"}, http.StatusOK)
// }

package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"social-network/models"
	"social-network/utils"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var userData utils.User
	if err := json.NewDecoder(r.Body).Decode(&userData); err != nil {
		fmt.Println(err)
		utils.WriteJSON(w, map[string]string{"message": "invalid input data"}, http.StatusBadRequest)

		return
	}

	// if !CheckNickName(userData.Nickname) {
	// 	utils.WriteJSON(w, map[string]string{"message": "invalid nickname"}, http.StatusUnauthorized)

	// 	return
	// }

	if !utils.CheckName(userData.LastName) {
		utils.WriteJSON(w, map[string]string{"message": "invalid last name"}, http.StatusBadRequest)

		return
	}

	if !utils.CheckName(userData.FirstName) {
		utils.WriteJSON(w, map[string]string{"message": "invalid first name"}, http.StatusBadRequest)

		return
	}

	if !utils.CheckAge(userData.Age) {
		utils.WriteJSON(w, map[string]string{"message": "nvalid age"}, http.StatusBadRequest)

		return
	}

	if !utils.CheckGender(userData.Gender) {
		utils.WriteJSON(w, map[string]string{"message": "nvalid gender"}, http.StatusBadRequest)

		return
	}

	if !utils.IsValidEmail(&userData.Email) {
		utils.WriteJSON(w,map[string]string{"message":  "nvalid emai"}, http.StatusBadRequest)

		return
	}
	if len(userData.Password) < 8 || len(userData.Password) > 64 {
		utils.WriteJSON(w, map[string]string{"message": "Invalid password"}, http.StatusBadRequest)

		return
	}

	ok, err := models.IsUserRegistered(&userData)
	if err != nil {
		fmt.Println(err)
		utils.WriteJSON(w,map[string]string{"message":  "internaInternal Server Error"}, http.StatusInternalServerError)

		return
	}

	if ok {
		utils.WriteJSON(w, map[string]string{"message": "User already exists"}, http.StatusConflict)

		return
	}

	err = utils.HashPassword(&userData.Password)
	if err != nil {
		utils.WriteJSON(w,map[string]string{"message":  "Incorect password"}, http.StatusBadRequest)

		return
	}

	// userData.NickName = html.EscapeString(userData.NickName)
	err = models.RegisterUser(&userData)
	if err != nil {
		utils.WriteJSON(w, map[string]string{"message": "internaInternal Server Error"}, http.StatusInternalServerError)

		return
	}

	// Create a session and set a cookie
	userData.SessionId, err = utils.GenerateSessionID()
	if err != nil {
		utils.WriteJSON(w,map[string]string{"message":  "internaInternal Server Error"}, http.StatusInternalServerError)

		return
	}

	err = models.InsertSession(&userData)
	if err != nil {
		utils.WriteJSON(w, map[string]string{"message": "internaInternal Server Error"}, http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "token",
		Path:  "/",
		Value: userData.SessionId,
	})

	utils.WriteJSON(w, map[string]string{"message": "registred succefully"}, http.StatusOK)
}
