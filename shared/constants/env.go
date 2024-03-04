package constants

import "github.com/vans-id/agit-technical-test-api.git/shared/util"

var (
	APP_PORT        = util.Getenv("APP_PORT")
	APP_DB_USERNAME = util.Getenv("APP_DB_USERNAME")
	APP_DB_PASSWORD = util.Getenv("APP_DB_PASSWORD")
	APP_DB_HOST     = util.Getenv("APP_DB_HOST")
	APP_DB_PORT     = util.Getenv("APP_DB_PORT")
	APP_DB_NAME     = util.Getenv("APP_DB_NAME")
	APP_SECRET_KEY  = util.Getenv("APP_SECRET_KEY")
)
