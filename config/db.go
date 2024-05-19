package config

import (
	"fmt"
	"log"

	"github.com/Giafn/goTodolistApi/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type connection struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
}

func InitDB(cfg *Config) *gorm.DB {
	conn := connection{
		DBHost:     cfg.DBHost,
		DBUser:     cfg.DBUser,
		DBPassword: cfg.DBPassword,
		DBName:     cfg.DBName,
		DBPort:     cfg.DBPort,
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		conn.DBHost, conn.DBUser, conn.DBPassword, conn.DBName, conn.DBPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&entity.Todo{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}
