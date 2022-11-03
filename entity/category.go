package entity

import uuid "github.com/satori/go.uuid"

type CategoryRepository interface {
	Get(_id string) (Category, error)
	Create(category Category) (Category, error)
}

type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewCategory() *Category {
	category := Category{
		ID: uuid.NewV4().String(),
	}

	return &category
}
