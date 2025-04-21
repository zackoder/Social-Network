package main

import (
	"net/http"

	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"

	"social-network/controllers"
	"social-network/db"
	"social-network/models"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	models.Db = db.InitDB()
	defer models.Db.Close()
	http.HandleFunc("/login", controllers.Login)
	http.HandleFunc("/register", controllers.Register)
	http.HandleFunc(" POST /addPost", controllers.AddPost)
	http.HandleFunc("/uploads/", controllers.HandelPics)
	http.HandleFunc("/api/posts", controllers.Posts)
	http.HandleFunc("/followReq", controllers.HandleFollow)
	http.HandleFunc("/updatePrivacy", controllers.UpdatePrivacy)
	http.HandleFunc("GET /api/getProfilePosts",controllers.GetProfilePosts)

	http.ListenAndServe(":8080", nil)
}
