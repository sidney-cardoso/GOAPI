package service

import (
	"github.com/sidney-cardoso/goapi/internal/database"
	"github.com/sidney-cardoso/goapi/internal/entity"
)

type ProductService struct {
	ProductDB database.ProductDB
}

func NewProductService(productDB database.ProductDB) *ProductService {
	return &ProductService{ProductDB: productDB}
}

func (prodServ *ProductService) GetProducts() ([]*entity.Product, error) {
	products, err := prodServ.ProductDB.GetProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (prodServ *ProductService) GetProduct(id string) (*entity.Product, error) {
	product, err := prodServ.ProductDB.GetProduct(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (prodServ *ProductService) GetProductByCategory(categoryID string) ([]*entity.Product, error) {
	products, err := prodServ.ProductDB.GetProductByCategory(categoryID)
	if err != nil {
		return nil, err
	}
	return products, err
}

func (prodServ *ProductService) CreateProduct(name, description, category_id, image_url string, price float64) (*entity.Product, error) {
	product := entity.NewProduct(
		name,
		description,
		category_id,
		image_url,
		price,
	)

	_, err := prodServ.ProductDB.CreateProduct(product)
	if err != nil {
		return nil, err
	}

	return product, nil
}
