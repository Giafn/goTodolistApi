package main

import (
	"fmt"
	"log"

	"github.com/Giafn/goTodolistApi/entity"
	"github.com/Giafn/goTodolistApi/router"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dsn := "host=127.0.0.1 user=gia password=gia dbname=todolist port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&entity.Todo{})
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()

	router.InitRoutes(e, db)

	fmt.Println("Server started at :8000")
	e.Logger.Fatal(e.Start(":8000"))
}
