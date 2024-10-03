package config

import (
	"os"

	"github.com/joho/godotenv"
)

// for database.
type Config struct {
	DB_HOST string
	DB_NAME string
	DB_PORT string
	DB_USER string
	DB_PASS string
}

var JwtKey []byte 

func GetConfig() Config {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	var config Config
	config.DB_HOST = os.Getenv("DB_HOST")
	config.DB_NAME = os.Getenv("DB_NAME")
	config.DB_PORT = os.Getenv("DB_PORT")
	config.DB_USER = os.Getenv("DB_USER")
	config.DB_PASS = os.Getenv("DB_PASS")

    jwtKey := os.Getenv("JWT_KEY")
    if jwtKey == ""{
        panic("jwt key not set in env")    
    }
    JwtKey = []byte(jwtKey)

	return config
}
