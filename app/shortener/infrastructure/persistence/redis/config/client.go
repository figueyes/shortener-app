package config

import (
	"fmt"
	"os"
	"strconv"
)

func CreateDbRedisConnection() *DbConnection {

	address := os.Getenv(fmt.Sprintf("%s%s", os.Getenv("APP_NAMESPACE"), "_REDIS_ADDRESS"))
	password := os.Getenv(fmt.Sprintf("%s%s", os.Getenv("APP_NAMESPACE"), "_REDIS_PASSWORD"))
	db := os.Getenv(fmt.Sprintf("%s%s", os.Getenv("APP_NAMESPACE"), "_REDIS_DB"))
	dbInt, _ := strconv.Atoi(db)

	//if len(address) == 0 || len(db) == 0 || len(password) == 0 {
	//	log.Fatal("data connection invalid")
	//}

	connection := NewRedisConnection(Config().
		SetDB(dbInt).
		SetPassword(password).
		SetAddress(address),
	)
	return connection
}
