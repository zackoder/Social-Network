package main

import (
	"fmt"

	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/gorilla/websocket"

	"net/http"
	"social-network/controllers"
	"social-network/db"
	"social-network/models"

	_ "github.com/mattn/go-sqlite3"
)

type Msg struct {
	Type string `json:type`
	Msg  string `json:content`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // for development; restrict this in production
	},
}

func main() {
	models.Db = db.InitDB()
	defer models.Db.Close()
	http.HandleFunc("/login", controllers.Login)
	http.HandleFunc("/register", controllers.Register)
	http.HandleFunc("/addPost", controllers.AddPost)
	http.HandleFunc("/uploads/", controllers.HandelPics)
	http.HandleFunc("/api/posts", controllers.Posts)
	http.HandleFunc("/followReq", controllers.HandleFollow)
	http.HandleFunc("/updatePrivacy", controllers.UpdatePrivacy)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(w, r, nil)
		for {
			// msg := Msg{}
			// error := conn.ReadJSON(&msg)

			_, msg, err := conn.ReadMessage()
			if err != nil {
				conn.Close()
				return
			}
			fmt.Println(string(msg))
		}
	})

	fmt.Println("http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
