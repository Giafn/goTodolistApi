package router

import (
	"net/http"

	"github.com/Giafn/goTodolistApi/handler"
	"github.com/Giafn/goTodolistApi/repository"
	"github.com/Giafn/goTodolistApi/service"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate) {
	todoRepository := repository.NewTodoRepository(db)
	todoService := service.NewTodoService(todoRepository)
	todoHandler := handler.NewTodoHandler(todoService, validate)
	// Membuat grup /api
	api := e.Group("/api")

	// Menambahkan route untuk hello world
	api.GET("", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Hello This Todo App With Gorm and Echo",
		})
	})

	api.GET("/todos", todoHandler.GetAllTodos)
	api.POST("/todos", todoHandler.CreateTodo)
	api.DELETE("/todos/:id", todoHandler.DeleteTodoByID)
	api.PUT("/todos/:id/done", todoHandler.DoneTodoByID)
	api.PUT("/todos/:id/edit", todoHandler.UpdateTodoTitleByID)
}
