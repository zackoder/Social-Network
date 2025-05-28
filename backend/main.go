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

	mux.HandleFunc("/login", controllers.Login)
	mux.HandleFunc("/register", controllers.Register)

	mux.Handle("POST /updatePrivacy", midleware.AuthMiddleware(controllers.UpdatePrivacy))
	mux.Handle("POST /addPost", midleware.AuthMiddleware(controllers.AddPost))
	mux.Handle("POST /followReq", midleware.AuthMiddleware(controllers.HandleFollow))
	mux.Handle("POST /creategroup", http.HandlerFunc(controllers.Creat_groupe))
	mux.Handle("/JouindGroupe", http.HandlerFunc(controllers.Jouind_Groupe))
	mux.Handle("/CreatEvent", midleware.AuthMiddleware(controllers.CreatEvent))

	mux.Handle("/groupInvitarion", http.HandlerFunc(controllers.InviteUser))

	mux.HandleFunc("/GetPostsFromGroupe", controllers.Get_all_post)

	// mux.Handle("POST /creategroup", (http.HandlerFunc(controllers.CreateGroup)))
	// mux.Handle("POST /joinReq",(http.HandlerFunc(controllers.JoinReq)))
	mux.Handle("POST /api/logout", midleware.AuthMiddleware(controllers.LogoutHandler))
	mux.Handle("GET /uploads/", midleware.AuthMiddleware(controllers.HandelPics))
	mux.Handle("GET /defaultIMG/", midleware.AuthMiddleware(controllers.HandelPics))
	mux.Handle("/api/posts", http.HandlerFunc(controllers.Posts))
	mux.HandleFunc("GET /api/getProfilePosts", midleware.AuthMiddleware(controllers.GetProfilePosts))
	mux.HandleFunc("GET /group", controllers.Group)
	mux.Handle("/event-resp", (http.HandlerFunc(controllers.EventResponse)))
	mux.Handle("/ws", midleware.AuthMiddleware(controllers.Websocket))

	// Comment handlers
	mux.Handle("GET /getComments", http.HandlerFunc(controllers.GetComments))
	mux.Handle("POST /addComment", midleware.AuthMiddleware(controllers.AddComment))

	// Reaction handlers
	mux.Handle("GET /getReactions", http.HandlerFunc(controllers.GetReactions))

	mux.Handle("/userData", midleware.AuthMiddleware(controllers.UserData))
	mux.Handle("POST /addReaction", midleware.AuthMiddleware(controllers.AddReaction))

	mux.Handle("/getNotifications", midleware.AuthMiddleware(controllers.GetNotifications))
	mux.Handle("/notiResp", midleware.AuthMiddleware(controllers.NotiResp))
	// mux.Handle("/private-messages", http.HandlerFunc(controllers.GetNotifications))

	mux.HandleFunc("GET /group/{GroupName}", controllers.Group)

	// note for walid
	// this endpoint is gonna be used to fetch the users for the post privacy
	mux.HandleFunc("GET /api/getfollowers", midleware.AuthMiddleware(controllers.GetFollowers))
	mux.HandleFunc("GET /api/getfollowinglist", controllers.GetfollowingsForProfile)
	mux.HandleFunc("GET /api/registrationData", midleware.AuthMiddleware(controllers.GetRegistrationData))
	// Notification handler
	mux.Handle("GET /api/notifications", midleware.AuthMiddleware(controllers.GetNotifications))

	// this endpoint is gonna be used to fetch the users for the chat pannel .
	// the func GetFollowers and GetFollowers look the same but they are not
	mux.HandleFunc("GET /api/getuserfriends", midleware.AuthMiddleware(controllers.GetUsers))
	fmt.Println("Server is running on port http://localhost:8080")
	http.ListenAndServe(":8080", midleware.WithCORS(mux))
}
