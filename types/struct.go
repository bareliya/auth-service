package types

type UserCred struct {
	UserId      int
	UserName    string
	FirstName   string
	LastName    string
	Hash        string
	AccessToken string
}

type UserRegister struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"user_name"`
	Password  string `json:"password"`
}
