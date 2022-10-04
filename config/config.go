package config

import (
	"fmt"
	"go-rest-api-orders/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host   = "localhost"
	port   = "5432"
	user   = "postgres"
	pass   = "root"
	dbname = "order_management"
)

func ConnectGorm() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, pass, dbname,
	)

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic(err)
	}

	db.Debug().AutoMigrate(models.Order{}, models.Item{})

	return db
}
