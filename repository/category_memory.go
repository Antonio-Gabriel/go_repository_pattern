package repository

import (
	"errors"

	"github.com/Antonio-Gabriel/go_repository_pattern/entity"
)

type CategoriesMemoryDb struct {
	Categories []entity.Category
}

type CategoryRepositoryMemory struct {
	db CategoriesMemoryDb
}

func NewCategoryRepositoryMemory(db CategoriesMemoryDb) *CategoryRepositoryMemory {
	return &CategoryRepositoryMemory{db: db}
}

func (c *CategoryRepositoryMemory) Get(_id string) (entity.Category, error) {
	for _, category := range c.db.Categories {
		if category.ID == _id {
			return category, nil
		}
	}

	return entity.Category{}, errors.New("no category found with this ID")
}

func (c *CategoryRepositoryMemory) Create(category entity.Category) (entity.Category, error) {
	c.db.Categories = append(c.db.Categories, category)
	return category, nil
}
