package utils

import "time"

// import "time"

type Regester struct {
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName"`
	NickName        any    `json:"nickName"`
	Age             int    `json:"age"`
	Gender          string `json:"gender"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfermPassword string `json:"confermPassword"`
	About_Me        string `json:"AboutMe"`
	Avatar          string `json:"avatar"`
}

type Login struct {
	Logininpt string `json:"login"`
	Password  string `json:"password"`
}

type Post struct {
	Id          int       `json:"id"`
	Privacy     string    `json:"privacy"`
	Poster_id   int       `json:"poster"`
	Poster_name string    `json:"first_name`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	Image       string    `json:"image"`
	Avatar      string    `json:"image"`
	Friendes    []int     `json:"friends"`
	CreatedAt   int       `json:"createdAt"`
	Reactions   Reactions `json:"reaction"`
	Groupe_id   int       `json:"groupe_id"`
}

type Reactions struct {
	Likes    int    `json:"likes"`
	Dislikes int    `json:"dislikes"`
	Action   string `json:"action"`
}

type NewGroup struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Notification struct {
	Id        int    `json:"id"`
	Message   string `json:"message"`
	Sender_id int    `json:"sender"`
	Actor_id  int    `json:"actor_id"`
	Target_id int    `json:"target"`
	Type      string `json:"-"`
}

type User struct {
	ID        int64
	Nickname  string `json:"nickname"`
	Age       uint8  `json:"age"`
	Gender    string `json:"gender"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	SessionId string
	Avatar    string `json:"avatar"`
	AboutMe   string `json:"aboutme"`
	Privacy   string `json:"privacy"`
}
type Session struct {
	Id        int    `json:"id"`
	UserId    int    `json:"user_id"`
	Token     string `json:"token"`
	CreatedAt string `json:"created_at"`
}

type Message struct {
	Sender_id   int    `json:"sender_id"`
	Reciever_id int    `json:"reciever_id"`
	Type        string `json:"type"`
	Group_id    int    `json:"group_id"`
	Content     string `json:"content"`
	Mime        string `json:"mime"`
	Filename    string `json:"filename"`
}

type Err struct {
	Error string `json:"error"`
}

type GroupInvitation struct {
	GroupID   int `json:"groupe_id"`
	InvitedBy int `json:"invited_By"`
	UserId    int `json:"invited"`
	CreatedAt time.Time
}
type Groupe struct {
	Id          int
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatorId   int    `json:"cretorid"`
	FirstName   string `json:"first_name"`
	LasttName   string `json:"last_name"`
}
type Groupe_member struct {
	User_id   int `json:"user_id"`
	Groupe_id int `json:"groupe_id"`
}
type Event struct {
	GroupID     int       `json:"groupe_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	EventTime   time.Time `json:"event_time"`
	CreatedBy   int       `json:"created_by"`
}

type EventResponse struct {
	UserID   int    `json:"user_id"`
	EventID  int    `json:"event_id"`
	GroupeId int    `json:"groupe_id"`
	Response string `json:"responce"`
}
