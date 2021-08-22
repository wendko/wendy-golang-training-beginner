package config

import (
	"fmt"
	"os"
)

func GetDBType() string {
	dBType := os.Getenv("DBType")
	return dBType
}

func GetPostgresConnectionString() string {
	dBHost := os.Getenv("DBHost")
	dBPort := os.Getenv("DBPort")
	dBUser := os.Getenv("DBUser")
	dBName := os.Getenv("DBName")
	dBPassword := os.Getenv("DBPassword")
	dataBase := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		dBHost,
		dBPort,
		dBUser,
		dBName,
		dBPassword)
	return dataBase
}