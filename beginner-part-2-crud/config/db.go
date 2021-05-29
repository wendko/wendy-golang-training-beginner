package config

import "fmt"

const (
	DBUser     = "test"
	DBPassword = "test"
	DBName     = "postgres"
	DBHost     = "0.0.0.0"
	DBPort     = "54322"
	DBType     = "postgres"
)

func GetDBType() string {
	return DBType
}

func GetPostgresConnectionString() string {
	dataBase := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		DBHost,
		DBPort,
		DBUser,
		DBName,
		DBPassword)
	return dataBase
}