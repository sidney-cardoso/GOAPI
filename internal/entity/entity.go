package entity

import "github.com/google/uuid"

type Category struct {
	ID   string
	Name string
}
type Product struct {
	ID          string
	Name        string
	Description string
	Price       float64
	CategoryID  string
	ImageURL    string
}

func NewCategory(name string) *Category {
	return &Category{
		ID:   uuid.New().String(),
		Name: name,
	}
}

func NewProduct(name, description, categoryID, imageURL string, price float64) *Product {
	return &Product{
		ID:          uuid.New().String(),
		Name:        name,
		Description: description,
		CategoryID:  categoryID,
		ImageURL:    imageURL,
		Price:       price,
	}
}
