package models

import (
	"github.com/SupratickDey/go_gin_backend/lib/common"
	"github.com/jinzhu/gorm"
)

// Post data model
type Post struct {
	gorm.Model
	Text   string `sql:"type:text;"`
	User   User   `gorm:"foreignkey:UserID"`
	UserID uint
}

// Serialize serializes posts data
func (p Post) Serialize() common.JSON {
	return common.JSON{
		"id":         p.ID,
		"text":       p.Text,
		"user":       p.User.Serialize(),
		"created_at": p.CreatedAt,
	}
}
