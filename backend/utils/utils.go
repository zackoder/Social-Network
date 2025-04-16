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
