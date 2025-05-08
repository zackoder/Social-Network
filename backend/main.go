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
	mux.Handle("/", (midleware.WithCORS(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tmp, err := template.ParseFiles("./index.html")
		if err != nil {
			fmt.Println(err)
		}
		tmp.Execute(w, nil)
	}))))
	mux.Handle("/login", midleware.WithCORS(http.HandlerFunc(controllers.Login)))
	mux.Handle("/register", midleware.WithCORS(http.HandlerFunc(controllers.Register)))
	mux.Handle("POST /addPost", midleware.WithCORS(http.HandlerFunc(controllers.AddPost)))
	mux.Handle("POST /followReq", midleware.WithCORS(http.HandlerFunc(controllers.HandleFollow)))
	mux.Handle("POST /updatePrivacy", midleware.WithCORS(http.HandlerFunc(controllers.UpdatePrivacy)))
	mux.Handle("POST /creategroup", midleware.WithCORS(http.HandlerFunc(controllers.Creat_groupe)))
	http.HandleFunc("/JouindGroupe", controllers.Jouind_Groupe)
	http.HandleFunc("/GetPostsFromGroupe",controllers.Get_all_post)
	http.HandleFunc("CreatEvent",controllers.CreatEvent)

	http.HandleFunc("/GetPostsFromGroupe", controllers.Get_all_post)
	mux.HandleFunc("GET /uploads/", controllers.HandelPics)
	mux.HandleFunc("/api/posts", controllers.Posts)
	mux.HandleFunc("GET /group/{GroupName}", controllers.Group)
	mux.HandleFunc("/ws", controllers.Websocket)

	http.ListenAndServe(":8080", mux)
}
