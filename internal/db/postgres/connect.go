package postgres

import (
	"college/internal/usecase/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func MustConnect() *gorm.DB {
	db, err := gorm.Open(postgres.Open("postgres://pg:pg@127.0.0.1:4000/devices"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to db" + err.Error())
	}

	if err = db.AutoMigrate(
		&models.Mouse{},
		&models.KeyBoard{},
		&models.Screen{},
		&models.Case{},
	); err != nil {
		panic("failed to migrate" + err.Error())
	}

	return db
}
