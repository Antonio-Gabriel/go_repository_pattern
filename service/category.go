package service

import "github.com/Antonio-Gabriel/go_repository_pattern/entity"

type CategoryService struct {
	Repository entity.CategoryRepository
}

func NewCategoryService(repository entity.CategoryRepository) *CategoryService {
	return &CategoryService{Repository: repository}
}

func (c *CategoryService) FindById(_id string) (entity.Category, error) {
	category, err := c.Repository.Get(_id)

	if err != nil {
		return entity.Category{}, err
	}

	return category, nil
}

func (c *CategoryService) Create(name string) (entity.Category, error) {
	category := entity.NewCategory()
	category.Name = name
	categoryRepository, err := c.Repository.Create(*category)

	if err != nil {
		return entity.Category{}, err
	}

	return categoryRepository, nil
}
