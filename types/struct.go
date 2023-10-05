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

type AdminCredential struct {
	AdminID           int    `json:"admin_id"`
	AdminUser         string `json:"admin_user"`
	AdminPassword     string `json:"admin_password"`
	AdminPasswordHash string `json:"admin_password_hash"`
	IsSuperAdmin      bool   `json:"is_super_admin"`
	AccessToken       string `json:"access_token"`
	IsApproved        bool   `json:"is_approved"`
}
