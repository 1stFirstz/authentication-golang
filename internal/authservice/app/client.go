package app

import (
	"github.com/1stFirstz/authentication-golang/internal/authservice/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type client struct {
	db *gorm.DB
}

func newClient() *client {
	return &client{
		db: connectDatabase(),
	}
}

func connectDatabase() *gorm.DB {
	dsn := "host=localhost user=admin password=P@ssw0rd dbname=project port=5000 sslmode=disable TimeZone=Asia/Bangkok"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	if err := db.AutoMigrate(&entity.User{}); err != nil {
		panic("failed to migrate database")
	}
	return db
}
