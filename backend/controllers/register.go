package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"social-network/models"
	"social-network/utils"
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

	if filePath == "" {
		regesterreq.Avatar = "/defaultIMG/defaulte.jpg"
		fmt.Println("file path", regesterreq.Avatar)

	}
	if regesterreq.NickName == "" {
		regesterreq.NickName = nil
	}
	if regesterreq.Age < 15 {
		utils.WriteJSON(w, map[string]string{"error": "you need to be older than 15 to register"}, http.StatusBadRequest)
		return
		// you need to be older than 15 to register
	}

	if utils.ValidatNames(regesterreq.FirstName, regesterreq.LastName, regesterreq.NickName) && utils.ValidEmail(regesterreq.Email) && regesterreq.Password == regesterreq.ConfermPassword {
		hashedPss := utils.Hashpass(regesterreq.Password)
		if hashedPss == "" {
			utils.WriteJSON(w, map[string]string{"error": "Internal Server Error"}, http.StatusInternalServerError)
			return
		}
		regesterreq.Password = hashedPss
		regesterreq.ID, err = models.InsertUser(regesterreq)
		if err != nil {
			if strings.Contains(err.Error(), "users.email") {
				utils.WriteJSON(w, map[string]string{"error": "Email has already been taken."}, http.StatusNotAcceptable)
			} else if strings.Contains(err.Error(), "users.nickname") {
				utils.WriteJSON(w, map[string]string{"error": "nickname has already been taken."}, http.StatusNotAcceptable)
			} else {
				utils.WriteJSON(w, map[string]string{"error": "Fail to register user"}, http.StatusNotAcceptable)
			}
			if filePath != "" {
				err := os.Remove(filePath)
				fmt.Println("removing error:", err)
			}
			return
		}
		utils.WriteJSON(w, map[string]string{"success": "ok"}, http.StatusOK)
		return
	}
}

// package controllers

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"strings"

// 	"social-network/models"
// 	"social-network/utils"
// )

// func Register(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
// 		utils.WriteJSON(w, map[string]string{"error": "Method Not Allowd"}, http.StatusMethodNotAllowed)
// 		return
// 	}
// 	var registrstionFormRequest utils.User
// 	userData := r.FormValue("userData")

// 	// max image size 10Mb
// 	if err := r.ParseMultipartForm(10 << 20); err != nil {
// 		utils.WriteJSON(w, map[string]string{"error": "Too larg file"}, http.StatusRequestEntityTooLarge)
// 		return
// 	}
// 	if err := json.Unmarshal([]byte(userData), &registrstionFormRequest); err != nil {
// 		utils.WriteJSON(w, map[string]string{"error": "Internal Server Error1"}, http.StatusMethodNotAllowed)
// 		fmt.Println(err)
// 		return
// 	}
// 	filePath, err := utils.UploadImage(r)
// 	if err != nil {
// 		utils.WriteJSON(w, map[string]string{"error": "Internal Server Error"}, http.StatusMethodNotAllowed)
// 		return
// 	}

// 	fmt.Println("requiiiiiiissssttttt", userData)
// 	fmt.Println("ppppppppppppppppppppppppppppppppppppppppppppppppppppppp", registrstionFormRequest)
// 	registrstionFormRequest.Avatar = filePath
// 	if registrstionFormRequest.Avatar == "" {
// 		registrstionFormRequest.Avatar = "/defaultIMG/defaulte.jpg"
// 	}

// 	if registrstionFormRequest.Nickname == "" {
// 		registrstionFormRequest.Nickname = ""
// 	}

// 	// if err := json.NewDecoder(r.Body).Decode(&userData); err != nil {
// 	// 	fmt.Println(err)
// 	// 	utils.WriteJSON(w, map[string]string{"message": "invalid input data"}, http.StatusBadRequest)

// 	// 	return
// 	// }

// 	// if !CheckNickName(userData.Nickname) {
// 	// 	utils.WriteJSON(w, map[string]string{"message": "invalid nickname"}, http.StatusUnauthorized)

// 	// 	return
// 	// }

// 	if !utils.CheckName(registrstionFormRequest.LastName) {
// 		utils.WriteJSON(w, map[string]string{"error": "invalid last name"}, http.StatusBadRequest)

// 		return
// 	}

// 	if !utils.CheckName(registrstionFormRequest.FirstName) {
// 		utils.WriteJSON(w, map[string]string{"error": "invalid first name"}, http.StatusBadRequest)

// 		return
// 	}

// 	// if !utils.CheckAge(registrstionFormRequest.Age) {
// 	// 	utils.WriteJSON(w, map[string]string{"error": "Invalid age"}, http.StatusBadRequest)

// 	// 	return
// 	// }

// 	if !utils.CheckGender(registrstionFormRequest.Gender) {
// 		utils.WriteJSON(w, map[string]string{"error": "Invalid gender"}, http.StatusBadRequest)

// 		return
// 	}

// 	if !utils.IsValidEmail(&registrstionFormRequest.Email) {
// 		utils.WriteJSON(w, map[string]string{"error": "Invalid emai"}, http.StatusBadRequest)

// 		return
// 	}
// 	if len(registrstionFormRequest.Password) < 8 || len(registrstionFormRequest.Password) > 64 {
// 		utils.WriteJSON(w, map[string]string{"error": "Invalid password"}, http.StatusBadRequest)

// 		return
// 	}

// 	ok, err := models.IsUserRegistered(&registrstionFormRequest)
// 	if err != nil {
// 		fmt.Println(err)
// 		utils.WriteJSON(w, map[string]string{"error": "internaInternal Server Error"}, http.StatusInternalServerError)

// 		return
// 	}

// 	if ok {
// 		utils.WriteJSON(w, map[string]string{"error": "User already exists"}, http.StatusConflict)

// 		return
// 	}

// 	err = utils.HashPassword(&registrstionFormRequest.Password)
// 	if err != nil {
// 		utils.WriteJSON(w, map[string]string{"error": "Incorect password"}, http.StatusBadRequest)

// 		return
// 	}

// 	// registrstionFormRequest.NickName = html.EscapeString(registrstionFormRequest.NickName)
// 	err = models.RegisterUser(&registrstionFormRequest)
// 	if err != nil {
// 		if strings.Contains(err.Error(), "nickname") {
// 			utils.WriteJSON(w, map[string]string{"error": "nickname already used"}, http.StatusBadRequest)
// 		}
// 		return
// 	}

// 	// Create a session and set a cookie
// 	registrstionFormRequest.SessionId, err = utils.GenerateSessionID()
// 	if err != nil {
// 		utils.WriteJSON(w, map[string]string{"error": "please try again"}, http.StatusInternalServerError)
// 		return
// 	}

// 	err = models.InsertSession(&registrstionFormRequest)
// 	if err != nil {
// 		utils.WriteJSON(w, map[string]string{"error": ""}, http.StatusInternalServerError)
// 		fmt.Println(err)
// 		return
// 	}

// 	http.SetCookie(w, &http.Cookie{
// 		Name:  "token",
// 		Path:  "/",
// 		Value: registrstionFormRequest.SessionId,
// 	})

// 	utils.WriteJSON(w, map[string]string{"message": "registred succefully"}, http.StatusOK)
// }
