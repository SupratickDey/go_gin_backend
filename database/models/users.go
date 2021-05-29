package models

import (
	"github.com/SupratickDey/go_gin_backend/lib/common"
	"github.com/jinzhu/gorm"
)

// User data model
type User struct {
	gorm.Model
	Email        string
	FirstName    string
	LastName     string
	Gender       string
	PasswordHash string
}

// Serialize serializes user data
func (u *User) Serialize() common.JSON {
	return common.JSON{
		"id":         u.ID,
		"email":      u.Email,
		"first_name": u.FirstName,
		"last_name":  u.LastName,
		"gender":     u.Gender,
	}
}

func (u *User) Read(m common.JSON) {
	u.ID = uint(m["id"].(float64))
	u.Email = m["email"].(string)
	u.FirstName = m["first_name"].(string)
	u.LastName = m["last_name"].(string)
	u.Gender = m["gender"].(string)
}
