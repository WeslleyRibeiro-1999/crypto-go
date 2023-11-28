package db

import (
	"os"

	"github.com/WeslleyRibeiro-1999/crypto-go/users/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(os.Getenv("DB_CONNECTION_USER")), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.User{})

	return db, nil
}
