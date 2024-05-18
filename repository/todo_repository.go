package repository

import (
	"github.com/Giafn/goTodolistApi/entity"
	"gorm.io/gorm"
)

type TodoRepository struct {
	Db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *TodoRepository {
	return &TodoRepository{Db: db}
}

func (repo *TodoRepository) GetAllTodos() ([]entity.Todo, error) {
	var todos []entity.Todo
	if err := repo.Db.Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

func (repo *TodoRepository) CreateTodo(todo *entity.Todo) error {
	if err := repo.Db.Create(todo).Error; err != nil {
		return err
	}
	return nil
}

func (repo *TodoRepository) GetTodoByID(id uint) (*entity.Todo, error) {
	var todo entity.Todo
	if err := repo.Db.First(&todo, id).Error; err != nil {
		return nil, err
	}
	return &todo, nil
}

func (repo *TodoRepository) UpdateTodo(todo *entity.Todo) error {
	if err := repo.Db.Save(todo).Error; err != nil {
		return err
	}
	return nil
}

func (repo *TodoRepository) DeleteTodoByID(id uint) error {
	if err := repo.Db.Delete(&entity.Todo{}, id).Error; err != nil {
		return err
	}
	return nil
}
