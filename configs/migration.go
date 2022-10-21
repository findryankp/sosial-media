package configs

import (
	"fmt"
	"sosialMedia/models"
)

func InitMigrate() {
	DB.AutoMigrate(
		&models.User{},
		&models.Photo{},
		&models.SosialMedia{},
		&models.Comment{},
	)

	fmt.Println("Migration done.")
}
