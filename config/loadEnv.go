package config

import (
	"os"
	"fmt"
)

// Code to load data from .env file
// and from here exports those .env data to be used across the app

var JWTKey string = os.Getenv("JWT_SECRET_KEY")

//Checking that an environment variable is present or not.
userName, ok := os.LookupEnv("DB_USERNAME")
if !ok {
  fmt.Println("DB_USERNAME is not present")
} else {
  fmt.Printf("DB_USERNAME Host: %s\n", userName)
}

dbURL := os.ExpandEnv("postgres://${DB_USERNAME}:${DB_PASSWORD}@$DB_HOST:$DB_PORT/$DB_NAME?sslmode=disable")
fmt.Println("DB URL: ", dbURL)
