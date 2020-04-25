package config

import "os"

const (
	address = "EMAIL_ADDRESS"
	pass    = "EMAIL_PASSWORD"
)

var (
	emailAddress  = os.Getenv(address)
	emailPassword = os.Getenv(pass)
)

func GetMailCredentials() (string, string) {
	return emailAddress, emailPassword
}
