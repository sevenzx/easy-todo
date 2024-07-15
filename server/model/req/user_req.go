package req

type UserRegister struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
