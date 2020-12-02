package database

import (
	"api/src/config"

	"github.com/jinzhu/gorm"

	// GORM Connection management
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Connect Estabilish database connection
func Connect() (*gorm.DB, error) {

	db, err := gorm.Open("postgres", config.ConnectionString)

	if err != nil {
		return nil, err
	}

	return db, nil
}
