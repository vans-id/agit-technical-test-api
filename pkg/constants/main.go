package constants

import (
	"os"
	"time"
)

var (
	APP_PORT        = os.Getenv("APP_PORT")
	APP_DB_USERNAME = os.Getenv("APP_DB_USERNAME")
	APP_DB_PASSWORD = os.Getenv("APP_DB_PASSWORD")
	APP_DB_HOST     = os.Getenv("APP_DB_HOST")
	APP_DB_PORT     = os.Getenv("APP_DB_PORT")
	APP_DB_NAME     = os.Getenv("APP_DB_NAME")

	JWT_ISS = "agit-technical-test-api"
	JWT_IAT = time.Now().Unix()
	JWT_EXP = time.Now().Add(time.Hour * 24).Unix()
)
