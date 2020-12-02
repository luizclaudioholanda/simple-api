package model

import (
	"github.com/jinzhu/gorm"
)

// Book represent a book in database
type Book struct {
	gorm.Model
	Name     string
	IsbnCode string
	Author   string
}
