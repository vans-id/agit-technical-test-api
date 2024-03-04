package migration

import (
	"github.com/vans-id/agit-technical-test-api.git/internal/entity"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	u := &entity.User{}
	e := &entity.Employee{}

	db.Migrator().DropTable(u)
	db.Migrator().DropTable(e)

	db.AutoMigrate(u, e)
}
