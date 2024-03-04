package pg

import (
	"fmt"

	"github.com/vans-id/agit-technical-test-api.git/pkg/constants"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func SetupDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("postgres://%v:%v@%v:%v/%v",
		constants.APP_DB_USERNAME,
		constants.APP_DB_PASSWORD,
		constants.APP_DB_HOST,
		constants.APP_DB_PORT,
		constants.APP_DB_NAME,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}
