package service

import (
	"github.com/sidney-cardoso/goapi/internal/database"
	"github.com/sidney-cardoso/goapi/internal/entity"
)

type CategoryService struct {
	CategoryDB database.CategoryDB
}

func NewCategoryService(categoryDB database.CategoryDB) *CategoryService {
	return &CategoryService{
		CategoryDB: categoryDB,
	}
}

func (catServ *CategoryService) GetCategories() ([]*entity.Category, error) {
	categories, err := catServ.CategoryDB.GetCategories()
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (catServ *CategoryService) CreateCategory(name string) (*entity.Category, error) {
	category := entity.NewCategory(name)
	_, err := catServ.CategoryDB.CreateCategory(category)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (catServ *CategoryService) GetCategory(id string) (*entity.Category, error) {
	category, err := catServ.CategoryDB.GetCategory(id)
	if err != nil {
		return nil, err
	}
	return category, nil
}
