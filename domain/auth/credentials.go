package auth

type Credentials struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	NewPassword string `json:"new_password,omitempty"`
}
