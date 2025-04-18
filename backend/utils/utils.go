package utils

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
	Id         int       `json:"id"`
	Privacy    string    `json:"privacy"`
	Poster     int       `json:"poster"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	Image      string    `json:"image"`
	Frendes    []string  `json:"frends"`
	CreatedAt  int       `json:"createdAt"`
	Reactions  Reactions `json:"reaction"`
}

type Reactions struct {
	Likes    int    `json:"likes"`
	Dislikes int    `json:"dislikes"`
	Action   string `json:"action"`
}
