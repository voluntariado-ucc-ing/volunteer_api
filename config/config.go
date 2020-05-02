package config

import "os"

const (
	address = "EMAIL_ADDRESS"
	pass    = "EMAIL_PASSWORD"

	dbName     = "DB_NAME"
	dbPassword = "DB_PASS"
	dbUser     = "DB_USER"
	dbHost     = "DB_HOST"

	SmtpHost    = "smtp.gmail.com"
	SmtpAddress = "smtp.gmail.com:587"
)

var (
	emailAddress     = os.Getenv(address)
	emailPassword    = os.Getenv(pass)
	databaseName     = os.Getenv(dbName)
	databaseUser     = os.Getenv(dbUser)
	databasePassword = os.Getenv(dbPassword)
	databaseHost     = os.Getenv(dbHost)
)

func GetMailCredentials() (string, string) {
	return emailAddress, emailPassword
}

func GetDatabaseHost() string {
	return databaseHost
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
