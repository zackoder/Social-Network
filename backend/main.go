package main

import (
	"net/http"

	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"

	"social-network/controllers"
	"social-network/db"
	"social-network/midleware"
	"social-network/models"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	mux := http.NewServeMux()
	
	mux.Handle("/login", midleware.WithCORS(http.HandlerFunc(controllers.Login)))
	mux.Handle("/register", midleware.WithCORS(http.HandlerFunc(controllers.Register)))
	mux.Handle("POST /addPost", midleware.WithCORS(http.HandlerFunc(controllers.AddPost)))
	mux.Handle("POST /followReq", midleware.WithCORS(http.HandlerFunc(controllers.HandleFollow)))
	mux.Handle("POST /updatePrivacy", midleware.WithCORS(http.HandlerFunc(controllers.UpdatePrivacy)))
	mux.Handle("POST /creategroup", midleware.WithCORS(http.HandlerFunc(controllers.CreateGroup)))
	mux.Handle("POST /joinReq", midleware.WithCORS(http.HandlerFunc(controllers.JoinReq)))
	mux.Handle("POST /api/logout", midleware.WithCORS(http.HandlerFunc(controllers.LogoutHandler)))
	mux.Handle("POST /api/reactions", midleware.WithCORS(http.HandlerFunc(controllers.HandelReactions)))
	models.Db = db.InitDB()
	defer models.Db.Close()
	mux.HandleFunc("GET /uploads/", controllers.HandelPics)
	mux.HandleFunc("GET /api/getProfilePosts", controllers.GetProfilePosts)
	mux.HandleFunc("GET /api/posts", controllers.Posts)
	mux.HandleFunc("GET /group/{GroupName}", controllers.Group)
	mux.HandleFunc("GET /api/getfollowers", controllers.GetFollowers)
	mux.HandleFunc("GET /api/registrationData", controllers.GetRegistrationData)
	mux.HandleFunc("/ws", controllers.Websocket)
	
	http.ListenAndServe(":8080", mux)
}
