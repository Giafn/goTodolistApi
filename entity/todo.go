package entity

import "time"

type Todo struct {
	ID        uint      `gorm:"primaryKey"`
	Title     string    `json:"title" validate:"required"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
