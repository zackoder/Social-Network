package main

import (
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"

	"net/http"
	"social-network/controllers"
	"social-network/db"
	"social-network/models"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	models.Db = db.InitDB()
	defer models.Db.Close()
	// serving uploaded images
	http.HandleFunc("/uploads/", controllers.HandelPics)

	http.HandleFunc("/login", controllers.Login)
	http.HandleFunc("/register", controllers.Register)
	http.HandleFunc("/addPost", controllers.AddPost)
	http.HandleFunc("/api/posts", controllers.Posts)
	http.HandleFunc("/followReq", controllers.HandleFollow)
	http.HandleFunc("/updatePrivacy", controllers.UpdatePrivacy)
	http.HandleFunc("/group/{GroupName}", controllers.Group)
	http.HandleFunc("/creategroup", controllers.CreateGroup)
	http.HandleFunc("/joinReq", controllers.JoinReq)
	http.HandleFunc("/ws", controllers.Websocket)

	http.ListenAndServe(":8080", nil)
}
