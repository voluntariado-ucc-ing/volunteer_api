package config

import "os"

const (
	address = "EMAIL_ADDRESS"
	pass    = "EMAIL_PASSWORD"

	dbName     = "DATABASE_NAME"
	dbPassword = "DB_PASSWORD"
	dbUser     = "DB_USER"

	SmtpHost    = "smtp.gmail.com"
	SmtpAddress = "smtp.gmail.com:587"
)

var (
	emailAddress     = os.Getenv(address)
	emailPassword    = os.Getenv(pass)
	databaseName     = os.Getenv(dbName)
	databaseUser     = os.Getenv(dbUser)
	databasePassword = os.Getenv(dbPassword)
)

func GetMailCredentials() (string, string) {
	return emailAddress, emailPassword
}

func GetDatabaseName() string {
	return databaseName
}

func GetDatabaseUser() string {
	return databaseUser
}

func GetDatabasePassword() string {
	return databasePassword
}
