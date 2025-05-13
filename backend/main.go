package main

import (
	"fmt"
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

	mux.Handle("/login",(http.HandlerFunc(controllers.Login)))
	mux.Handle("/register",(http.HandlerFunc(controllers.Register)))
	mux.Handle("POST /addPost",(http.HandlerFunc(controllers.AddPost)))
	mux.Handle("POST /followReq",(http.HandlerFunc(controllers.HandleFollow)))
	mux.Handle("POST /updatePrivacy",(http.HandlerFunc(controllers.UpdatePrivacy)))
	mux.Handle("POST /creategroup", (http.HandlerFunc(controllers.CreateGroup)))
	mux.Handle("POST /joinReq",(http.HandlerFunc(controllers.JoinReq)))
	mux.Handle("POST /api/logout",(http.HandlerFunc(controllers.LogoutHandler)))

	models.Db = db.InitDB()
	defer models.Db.Close()
	mux.HandleFunc("GET /uploads/", controllers.HandelPics)
	mux.HandleFunc("GET /api/getProfilePosts",midleware.AuthMiddleware(controllers.GetProfilePosts))
	mux.HandleFunc("GET /api/posts", controllers.Posts)
	mux.HandleFunc("GET /group/{GroupName}", controllers.Group)
	
	// note for walid 
	// this endpoint is gonna be used to fetch the users for the post privacy  
	mux.HandleFunc("GET /api/getfollowers",midleware.AuthMiddleware(controllers.GetFollowers))
	mux.HandleFunc("GET /api/registrationData",midleware.AuthMiddleware(controllers.GetRegistrationData))
	// this endpoint is gonna be used to fetch the users for the chat pannel . 
	// the func GetFollowers and GetFollowers they look the same but they are not 🕶😎
	mux.HandleFunc("GET /api/getuserfriends",midleware.AuthMiddleware(controllers.GetUsers))
    mux.HandleFunc("/ws", controllers.Websocket)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", midleware.WithCORS(mux))
}




