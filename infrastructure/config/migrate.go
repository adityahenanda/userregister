package config

import (
	"userregister/models"

	"github.com/jinzhu/gorm"
)

func Migrate(idb *gorm.DB) {
	idb.Debug().AutoMigrate(
		&models.UserAccount{},
	)
}
