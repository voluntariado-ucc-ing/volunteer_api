package auth

type Credentials struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	NewPassword string `json:"new_password,omitempty"`
}

type MailCredentials struct {
	Mail string `json:"mail"`
	Pass string `json:"pass"`
}

type MailRequest struct {
	Volunteers []MailCredentials `json:"volunteers"`
}
