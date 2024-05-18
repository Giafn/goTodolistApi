package service

import (
	"github.com/Giafn/goTodolistApi/entity"
	"github.com/Giafn/goTodolistApi/repository"
)

type TodoService struct {
	Repo *repository.TodoRepository
}

func NewTodoService(repo *repository.TodoRepository) *TodoService {
	return &TodoService{Repo: repo}
}

func (service *TodoService) GetAllTodos() ([]entity.Todo, error) {
	return service.Repo.GetAllTodos()
}

func (service *TodoService) GetTodoByID(id uint) (*entity.Todo, error) {
	return service.Repo.GetTodoByID(id)
}

func (service *TodoService) CreateTodo(todo *entity.Todo) error {
	return service.Repo.CreateTodo(todo)
}

func (service *TodoService) MarkTodoAsDone(id uint) error {
	// Ambil todo dari repository berdasarkan ID
	todo, err := service.Repo.GetTodoByID(id)
	if err != nil {
		return err
	}

	todo.Completed = true

	if err := service.Repo.UpdateTodo(todo); err != nil {
		return err
	}

	return nil
}

func (service *TodoService) UpdateTodoTitle(id uint, title string) error {
	// Ambil todo dari repository berdasarkan ID
	todo, err := service.Repo.GetTodoByID(id)
	if err != nil {
		return err
	}

	todo.Title = title

	if err := service.Repo.UpdateTodo(todo); err != nil {
		return err
	}

	return nil
}

func (service *TodoService) DeleteTodoByID(id uint) error {
	return service.Repo.DeleteTodoByID(id)
}
