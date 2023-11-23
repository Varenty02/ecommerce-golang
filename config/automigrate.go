package config

import (
	usermodel "Ecommerce/module/model"
	"gorm.io/gorm"
)

func AutoMigrateDB(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(
		//auto migrate for user
		&usermodel.User{},
		&usermodel.UserCreate{},
		&usermodel.UserLogin{},
		&usermodel.UserDelete{},

	)
	return db
}
