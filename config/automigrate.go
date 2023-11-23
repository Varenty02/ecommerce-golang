package config

import (
	usermodel "Ecommerce/module/model"
	"gorm.io/gorm"
)

func AutoMigrateDB(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(
		&usermodel.UserCreate{},
	)
	return db
}
