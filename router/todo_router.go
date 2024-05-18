package router

import (
	"github.com/Giafn/goTodolistApi/handler"
	"github.com/Giafn/goTodolistApi/repository"
	"github.com/Giafn/goTodolistApi/service"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRoutes(e *echo.Echo, db *gorm.DB) {
	validate := validator.New()

	todoRepository := repository.NewTodoRepository(db)
	todoService := service.NewTodoService(todoRepository)
	todoHandler := handler.NewTodoHandler(todoService, validate)

	e.GET("/todos", todoHandler.GetAllTodos)
	e.POST("/todos", todoHandler.CreateTodo)
	e.DELETE("/todos/:id", todoHandler.DeleteTodoByID)
	e.PUT("/todos/:id/done", todoHandler.DoneTodoByID)
	e.PUT("/todos/:id/edit", todoHandler.UpdateTodoTitleByID)
}
