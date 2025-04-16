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
	http.HandleFunc("/login", controllers.Login)
	http.HandleFunc("/register", controllers.Register)
	http.HandleFunc("/addPost", controllers.AddPost)

	http.ListenAndServe(":8080", nil)
}
