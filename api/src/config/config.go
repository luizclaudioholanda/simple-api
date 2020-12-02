package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// ConnectionString database connection string
	ConnectionString = ""

	// Port Database port number
	Port = 0
)

//LoadConfig loads application configuration
func LoadConfig() {

	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Port, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		Port = 5001
	}

	ConnectionString = fmt.Sprintf("host=localhost port=5432 user=%s dbname=%s sslmode=disable password=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PWD"),
	)
}
