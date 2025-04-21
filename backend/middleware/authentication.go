package utils

import (
	"database/sql"
	"net/http"
	"sync"
	"time"

	utils "social-network/utils"
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

type customHandler func(w http.ResponseWriter, r *http.Request, db *sql.DB, userId int)

func AuthMiddleware(db *sql.DB, next customHandler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		allowed := rateLimit.Allow(r.RemoteAddr)
		if !allowed {
			w.WriteHeader(http.StatusTooManyRequests)
			return
		}
		userId, err := ValidUser(r, db)
		if err != nil {
			if err == http.ErrNoCookie {
				utils.WriteJSON(w, map[string]string{"error": "Unauthorized"}, http.StatusUnauthorized)
				return
			} else if err == sql.ErrNoRows {
				http.SetCookie(w, &http.Cookie{
					Name:    "token",
					Path:    "/",
					Value:   "",
					Expires: time.Unix(0, 0),
				})
				utils.WriteJSON(w, map[string]string{"error": "Unauthorized"}, http.StatusUnauthorized)
				return
			} else {
				utils.WriteJSON(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
		}
		next(w, r, db, userId)
	})
}

// func IsUserRegistered(db *sql.DB, userData *utils.User) (bool, error) {
// 	var exists bool
// 	query := `SELECT EXISTS(SELECT 1 FROM users WHERE email = ? OR nickname = ?);`
// 	err := db.QueryRow(query, userData.Email, userData.Nickname).Scan(&exists)
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

// func InsertSession(db *sql.DB, userData *utils.User) error {
// 	_, err := db.Exec("INSERT INTO sessions ( user_id, token, expires_at) VALUES (?, ?, ?)", userData.ID, userData.SessionId, userData.Expiration)
// 	return err
// }

// func GetActiveSession(db *sql.DB, userData *utils.User) (bool, error) {
// 	var exists bool
// 	currentTime := time.Now()
// 	fmt.Println(currentTime)
// 	query := `SELECT EXISTS(SELECT 1 FROM sessions WHERE user_id = ?  AND expires_at > ?);`
// 	err := db.QueryRow(query, userData.ID, currentTime).Scan(&exists)
// 	if err != nil {
// 		return false, err
// 	}
// 	return exists, nil
// }

// func DeleteSession(db *sql.DB, userData *utils.User) error {
// 	query := `DELETE FROM sessions WHERE user_id =  ?;`
// 	_, err := db.Exec(query, userData.ID)
// 	return err
// }

// func ValidCredential(db *sql.DB, userData *utils.User) error {
// 	query := `SELECT id, password FROM users WHERE nickname = ? OR email = ?;`
// 	err := db.QueryRow(query, userData.Nickname, userData.Email).Scan(&userData.ID, &userData.Password)
// 	if err != nil {
// 		return err
// 	}
// 	return err
// }

func ValidUser(r *http.Request, db *sql.DB) (int, error) {
	token := r.Header.Get("token")
	if token == "" {
		// Fallback to checking cookies
		cookie, err := r.Cookie("token")
		if err != nil {
			return 0, err
		}
		token = cookie.Value
	}
	userId, err := Get_session(token, db)
	if err != nil {
		return 0, err
	}
	return userId, nil
}

func Get_session(ses string, db *sql.DB) (int, error) {
	var sessionid int
	query := `SELECT user_id FROM sessions WHERE token = ?`
	err := db.QueryRow(query, ses).Scan(&sessionid)
	if err != nil {
		return 0, err
	}
	return sessionid, nil
}
