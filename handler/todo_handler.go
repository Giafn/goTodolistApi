package handler

import (
	"net/http"
	"strconv"

	"github.com/Giafn/goTodolistApi/entity"
	"github.com/Giafn/goTodolistApi/service"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type TodoHandler struct {
	Service   *service.TodoService
	Validator *validator.Validate // Tambahkan validator sebagai properti
}

func NewTodoHandler(service *service.TodoService, validate *validator.Validate) *TodoHandler {
	return &TodoHandler{
		Service:   service,
		Validator: validate, // Inisialisasi properti validator
	}
}

func (handler *TodoHandler) GetAllTodos(c echo.Context) error {
	todos, err := handler.Service.GetAllTodos()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	successMessage := entity.NewAPIResponse(http.StatusOK, "Successfully show todo", todos)
	return c.JSON(http.StatusOK, successMessage)
}

func (handler *TodoHandler) CreateTodo(c echo.Context) error {
	todo := new(entity.Todo)
	if err := c.Bind(todo); err != nil {
		return err
	}

	// Validasi input menggunakan validator
	if err := handler.Validator.Struct(todo); err != nil {
		errorMessage := entity.NewAPIResponse(http.StatusBadRequest, err.Error(), nil)
		return c.JSON(http.StatusBadRequest, errorMessage)
	}

	if err := handler.Service.CreateTodo(todo); err != nil {
		errorMessage := entity.NewAPIResponse(http.StatusInternalServerError, "Failed to create todo", nil)
		return c.JSON(http.StatusInternalServerError, errorMessage)
	}

	createdId := map[string]interface{}{"created_id": todo.ID}
	successMessage := entity.NewAPIResponse(http.StatusCreated, "Successfully created todo", createdId)
	return c.JSON(http.StatusCreated, successMessage)
}

func (handler *TodoHandler) DoneTodoByID(c echo.Context) error {
	// Mendapatkan ID dari path parameter
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		errorMessage := entity.NewAPIResponse(http.StatusBadRequest, "Invalid ID", nil)
		return c.JSON(http.StatusBadRequest, errorMessage)
	}

	if _, err = handler.Service.GetTodoByID(uint(id)); err != nil {
		errorMessage := entity.NewAPIResponse(http.StatusNotFound, "ID not found", nil)
		return c.JSON(http.StatusNotFound, errorMessage)
	}

	// Menyelesaikan task menggunakan service
	err = handler.Service.MarkTodoAsDone(uint(id))
	if err != nil {
		errorMessage := entity.NewAPIResponse(http.StatusInternalServerError, "Failed to mark todo as done", nil)
		return c.JSON(http.StatusInternalServerError, errorMessage)
	}

	// Respon berhasil jika tidak ada error
	successMessage := entity.NewAPIResponse(http.StatusOK, "Successfully marked todo as done", nil)
	return c.JSON(http.StatusOK, successMessage)
}

func (handler *TodoHandler) DeleteTodoByID(c echo.Context) error {
	// Mendapatkan ID dari path parameter
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		errorMessage := entity.NewAPIResponse(http.StatusBadRequest, "Invalid ID", nil)
		return c.JSON(http.StatusBadRequest, errorMessage)
	}

	if _, err = handler.Service.GetTodoByID(uint(id)); err != nil {
		errorMessage := entity.NewAPIResponse(http.StatusNotFound, "ID not found", nil)
		return c.JSON(http.StatusNotFound, errorMessage)
	}

	// Menyelesaikan task menggunakan service
	err = handler.Service.DeleteTodoByID(uint(id))
	if err != nil {
		errorMessage := entity.NewAPIResponse(http.StatusInternalServerError, "Failed to delete todo", nil)
		return c.JSON(http.StatusInternalServerError, errorMessage)
	}

	// Respon berhasil jika tidak ada error
	successMessage := entity.NewAPIResponse(http.StatusOK, "Successfully deleted todo", nil)
	return c.JSON(http.StatusOK, successMessage)
}

func (handler *TodoHandler) UpdateTodoTitleByID(c echo.Context) error {
	// Mendapatkan ID dari path parameter
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		errorMessage := entity.NewAPIResponse(http.StatusBadRequest, "Invalid ID", nil)
		return c.JSON(http.StatusBadRequest, errorMessage)
	}

	if _, err = handler.Service.GetTodoByID(uint(id)); err != nil {
		errorMessage := entity.NewAPIResponse(http.StatusNotFound, "ID not found", nil)
		return c.JSON(http.StatusNotFound, errorMessage)
	}

	var updateTitleRequest struct {
		Title string `json:"title" validate:"required"`
	}

	if err := c.Bind(&updateTitleRequest); err != nil {
		errorMessage := entity.NewAPIResponse(http.StatusBadRequest, "Invalid request body", nil)
		return c.JSON(http.StatusBadRequest, errorMessage)
	}

	// Validasi input menggunakan validator
	if err := handler.Validator.Struct(updateTitleRequest); err != nil {
		errorMessage := entity.NewAPIResponse(http.StatusBadRequest, err.Error(), nil)
		return c.JSON(http.StatusBadRequest, errorMessage)
	}

	// Memperbarui judul task menggunakan service
	err = handler.Service.UpdateTodoTitle(uint(id), updateTitleRequest.Title)
	if err != nil {
		errorMessage := entity.NewAPIResponse(http.StatusInternalServerError, "Failed to update todo title", nil)
		return c.JSON(http.StatusInternalServerError, errorMessage)
	}

	// Respon berhasil jika tidak ada error
	successMessage := entity.NewAPIResponse(http.StatusOK, "Successfully updated todo title", nil)
	return c.JSON(http.StatusOK, successMessage)
}
