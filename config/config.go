package config

import "os"

const (
	address = "EMAIL_ADDRESS"
	pass    = "EMAIL_PASSWORD"

	dbName     = "DB_NAME"
	dbPassword = "DB_PASS"
	dbUser     = "DB_USER"
	dbHost     = "DB_HOST"
	dbPort     = "DB_PORT"

	redisPort = "REDIS_PORT"
	redisHost = "REDIS_HOST"

	SmtpHost    = "smtp.gmail.com"
	SmtpAddress = "smtp.gmail.com:587"
)

var (
	emailAddress     = "voluntariadoing.noreply@ucc.edu.ar" //os.Getenv(address)
	emailPassword    = "ysl*gzzjic4Taok"                    //os.Getenv(pass)
	databaseName     = os.Getenv(dbName)
	databaseUser     = os.Getenv(dbUser)
	databasePassword = os.Getenv(dbPassword)
	databaseHost     = os.Getenv(dbHost)
	databasePort     = os.Getenv(dbPort)

	redisPortValue = os.Getenv(redisPort)
	redisHostValue = os.Getenv(redisHost)
)

func GetRedisPort() string {
	return redisPortValue
}

func GetRedisHost() string {
	return redisHostValue
}

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

func GetDatabasePort() string {
	return databasePort
}
