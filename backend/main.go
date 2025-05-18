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
	mux.HandleFunc("/JouindGroupe", controllers.Jouind_Groupe)
	mux.HandleFunc("/GetPostsFromGroupe", controllers.Get_all_post)
	mux.HandleFunc("/CreatEvent", controllers.CreatEvent)

	mux.Handle("POST /api/logout", midleware.WithCORS(http.HandlerFunc(controllers.LogoutHandler)))

	mux.HandleFunc("GET /uploads/", controllers.HandelPics)
	mux.Handle("/api/posts", midleware.WithCORS(http.HandlerFunc(controllers.Posts)))
	mux.HandleFunc("GET /api/getProfilePosts", controllers.GetProfilePosts)
	mux.HandleFunc("GET /group/{GroupName}", controllers.Group)
	mux.HandleFunc("/ws", controllers.Websocket)

	// Comment handlers
	mux.Handle("POST /addComment", midleware.WithCORS(http.HandlerFunc(controllers.AddComment)))
	mux.Handle("GET /getComments", midleware.WithCORS(http.HandlerFunc(controllers.GetComments)))

	// Reaction handlers
	mux.Handle("POST /addReaction", midleware.WithCORS(http.HandlerFunc(controllers.AddReaction)))
	mux.Handle("GET /getReactions", midleware.WithCORS(http.HandlerFunc(controllers.GetReactions)))

	mux.HandleFunc("GET /api/getfollowers", controllers.GetFollowers)
	mux.HandleFunc("GET /api/registrationData", controllers.GetRegistrationData)
	mux.Handle("/userData", midleware.WithCORS(http.HandlerFunc(controllers.UserData)))

	fmt.Println("localhost:8080")
	http.ListenAndServe(":8080", midleware.WithCORS(mux))
}
