package main

import (
	"fmt"

	"github.com/Giafn/goTodolistApi/config"
	"github.com/Giafn/goTodolistApi/router"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func main() {
	cfg := config.LoadConfig()
	db := config.InitDB(cfg)
	e := echo.New()
	validator := validator.New()

	router.InitRoutes(e, db, validator)

	fmt.Printf("Server started at :%s\n", cfg.Port)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.Port)))
}
