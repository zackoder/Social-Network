package utils

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
	Id        int       `json:"id"`
	Privacy   string    `json:"privacy"`
	Poster    int       `json:"poster"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Image     string    `json:"image"`
	Friendes  []string  `json:"friends"`
	CreatedAt int       `json:"createdAt"`
	Reactions Reactions `json:"reaction"`
}

type Reactions struct {
	Likes    int    `json:"likes"`
	Dislikes int    `json:"dislikes"`
	Action   string `json:"action"`
}

type User struct {
	ID         int64
	Nickname   string `json:"nickname"`
	Age        uint8  `json:"age"`
	Gender     string `json:"gender"`
	FirstName  string `json:"firstname"`
	LastName   string `json:"lastname"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	SessionId  string
}
type Session struct {
	Id        int    `json:"id"`
	UserId    int    `json:"user_id"`
	Token     string `json:"token"`
	CreatedAt string `json:"created_at"`
}
