package storage

import (
	"log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	config "github.com/wendko/wendy-golang-training-beginner/beginner-part-2-crud/config"
)

var DB *gorm.DB

func NewDB(params ...string) *gorm.DB {
	var err error
	conString := config.GetPostgresConnectionString()

	log.Print(conString)

	DB, err = gorm.Open(config.GetDBType(), conString)

	if err != nil {
		log.Panic(err)
	}

	return DB
}

func GetDBInstance() *gorm.DB {
	return DB
}