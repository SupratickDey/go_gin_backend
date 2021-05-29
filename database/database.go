package database

import (
	"github.com/SupratickDey/go_gin_backend/database/models"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func Initilize()(*gorm.DB,error){
	dbConfig := os.Getenv("DB_CONFIG")
	db, err := gorm.Open("postgres", dbConfig)
	db.LogMode(true) // logs SQL
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Connected to database")
	models.Migrate(db)
	return db, err
}
