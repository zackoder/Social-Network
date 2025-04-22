package utils

import (
	"database/sql"
	"net/http"
	"sync"
	"time"

	utils "social-network/utils"
	models "social-network/models"
)

type Limit struct {
	LastTime int64
	Counter  int
}
type RateLimit struct {
	User map[string]Limit
	Mu   sync.Mutex
}

func NewRateLimit() *RateLimit {
	return &RateLimit{
		User: make(map[string]Limit),
	}
}

var rateLimit = NewRateLimit()

func (r *RateLimit) Allow(ip string) bool {
	r.Mu.Lock()
	defer r.Mu.Unlock()
	if _, ok := r.User[ip]; !ok {
		currentTime := time.Now()
		r.User[ip] = Limit{LastTime: currentTime.Unix(), Counter: 1}
	} else {
		limit := r.User[ip]
		currentTime := time.Now()
		if limit.LastTime == currentTime.Unix() {
			limit.Counter++
			if limit.Counter > 20 {
				return false
			}
		} else {
			limit.LastTime = currentTime.Unix()
			limit.Counter = 1
		}
	}
	return true
}

type customHandler func(w http.ResponseWriter, r *http.Request, userId int)

func AuthMiddleware( next customHandler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		allowed := rateLimit.Allow(r.RemoteAddr)
		if !allowed {
			w.WriteHeader(http.StatusTooManyRequests)
			return
		}
		userId, err := ValidUser(r)
		if err != nil {
			if err == http.ErrNoCookie {
				utils.WriteJSON(w, map[string]string{"error": "Unauthorized"}, http.StatusUnauthorized)
				return
			} else if err == sql.ErrNoRows {
				http.SetCookie(w, &http.Cookie{
					Name:    "token",
					Path:    "/",
					Value:   "",
				})
				utils.WriteJSON(w, map[string]string{"error": "Unauthorized"}, http.StatusUnauthorized)
				return
			} else {
				utils.WriteJSON(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
		}
		next(w, r, userId)
	})
}

// func IsUserRegistered(db *sql.DB, userData *utils.User) (bool, error) {
// 	var exists bool
// 	query := `SELECT EXISTS(SELECT 1 FROM users WHERE email = ?);`
// 	err := db.QueryRow(query, userData.Email).Scan(&exists)
// 	return exists, err
// }

// func RegisterUser(db *sql.DB, userData *utils.User) error {
// 	insertQuery := `INSERT INTO users (nickname, age, gender, firstname, lastname, email, password) VALUES (?, ?, ?, ?, ?, ?, ?);`
// 	result, err := db.Exec(insertQuery, userData.Nickname, userData.Age, userData.Gender, userData.FirstName, userData.LastName, userData.Email, userData.Password)
// 	if err != nil {
// 		return err
// 	}
// 	userData.ID, err = result.LastInsertId()
// 	return err
// }


func ValidUser(r *http.Request) (int, error) {
		cookie, err := r.Cookie("token")
		if err != nil {
			return 0, err
		}
	userId, err := models.Get_session(cookie.Value)
	if err != nil {
		return 0, err
	}
	return userId, nil
}


