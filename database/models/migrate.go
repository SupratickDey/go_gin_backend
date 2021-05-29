package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

// Migrate auto-migrates models using ORM
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&User{}, &Post{})
	// set up foreign keys
	db.Model(&Post{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	fmt.Println("Auto Migration has been processed")
}
