package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/figueyes/shortener-app/app/shortener/infrastructure/log"
)

func CreateDbMongoConnection() *DbConnection {
	var port int
	var err error

	dbHost := os.Getenv(fmt.Sprintf("%s%s", os.Getenv("APP_NAMESPACE"), "_MONGODB_HOST"))
	dbPort := os.Getenv(fmt.Sprintf("%s%s", os.Getenv("APP_NAMESPACE"), "_MONGODB_PORT"))
	dbDatabase := os.Getenv(fmt.Sprintf("%s%s", os.Getenv("APP_NAMESPACE"), "_MONGODB_DATABASE"))
	dbUsername := os.Getenv(fmt.Sprintf("%s%s", os.Getenv("APP_NAMESPACE"), "_MONGODB_USERNAME"))
	dbPassword := os.Getenv(fmt.Sprintf("%s%s", os.Getenv("APP_NAMESPACE"), "_MONGODB_PASSWORD"))

	//if len(dbHost) == 0 || len(dbPort) == 0 || len(dbDatabase) == 0 {
	//	log.Fatal("data connection invalid")
	//}
	if port, err = strconv.Atoi(dbPort); err != nil {
		log.Error("invalid port error %s", dbPort)
	}
	connection := NewMongoConnection(Config().
		Host(dbHost).
		Port(port).
		DatabaseName(dbDatabase).
		User(dbUsername).
		Password(dbPassword),
	)
	return connection
}
