package main

import (
	"fmt"
	"net/http"
	"text/template"

	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"

	"social-network/controllers"
	"social-network/db"
	"social-network/midleware"
	"social-network/models"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	models.Db = db.InitDB()
	defer models.Db.Close()
	mux := http.NewServeMux()

	mux.Handle("/", (http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tmp, err := template.ParseFiles("./index.html")
		if err != nil {
			fmt.Println(err)
		}
		tmp.Execute(w, nil)
	})))

	mux.Handle("/login", http.HandlerFunc(controllers.Login))
	mux.Handle("/register", http.HandlerFunc(controllers.Register))

	mux.Handle("POST /addPost", midleware.AuthMiddleware(controllers.AddPost))
	mux.Handle("POST /followReq", midleware.AuthMiddleware(controllers.HandleFollow))
	mux.Handle("POST /updatePrivacy", http.HandlerFunc(controllers.UpdatePrivacy))
	mux.Handle("POST /creategroup", http.HandlerFunc(controllers.Creat_groupe))
	mux.Handle("/JouindGroupe", http.HandlerFunc(controllers.Jouind_Groupe))
	mux.Handle("/CreatEvent", midleware.AuthMiddleware(controllers.CreatEvent))

	mux.Handle("/groupInvitarion", http.HandlerFunc(controllers.InviteUser))
	mux.HandleFunc("/GetPostsFromGroupe", controllers.Get_all_post)

	mux.Handle("POST /api/logout", http.HandlerFunc(controllers.LogoutHandler))

	mux.HandleFunc("GET /uploads/", controllers.HandelPics)
	mux.Handle("/api/posts", http.HandlerFunc(controllers.Posts))
	mux.HandleFunc("GET /api/getProfilePosts", controllers.GetProfilePosts)
	mux.HandleFunc("GET /group", controllers.Group)
	mux.Handle("/event-resp", (http.HandlerFunc(controllers.EventResponse)))
	mux.HandleFunc("/ws", controllers.Websocket)

	// Comment handlers
	mux.Handle("POST /addComment", http.HandlerFunc(controllers.AddComment))
	mux.Handle("GET /getComments", http.HandlerFunc(controllers.GetComments))

	// Reaction handlers
	mux.Handle("POST /addReaction", http.HandlerFunc(controllers.AddReaction))
	mux.Handle("GET /getReactions", http.HandlerFunc(controllers.GetReactions))

	mux.HandleFunc("GET /api/getfollowers", controllers.GetFollowers)
	mux.HandleFunc("GET /api/registrationData", controllers.GetRegistrationData)
	mux.Handle("/userData", http.HandlerFunc(controllers.UserData))

	mux.Handle("/getNotifications", midleware.AuthMiddleware(controllers.GetNotifications))
	mux.Handle("/notiResp", midleware.AuthMiddleware(controllers.NotiResp))
	// mux.Handle("/private-messages", http.HandlerFunc(controllers.GetNotifications))

	fmt.Println("localhost:8080")
	http.ListenAndServe(":8080", midleware.WithCORS(mux))
}
