package database

import (
	"database/sql"

	"github.com/sidney-cardoso/goapi/internal/entity"
)

type ProductDB struct {
	db *sql.DB
}

func NewProductDB(db *sql.DB) *ProductDB {
	return &ProductDB{db: db}
}

func (prodDB *ProductDB) GetProducts() ([]*entity.Product, error) {
	rows, err := prodDB.db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*entity.Product
	for rows.Next() {
		var product entity.Product
		if err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.CategoryID,
			&product.ImageURL,
		); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products, nil
}

func (prodDB *ProductDB) GetProduct(id string) (*entity.Product, error) {
	var product entity.Product

	err := prodDB.db.
		QueryRow("SELECT * FROM products WHERE id = ?", id).Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.CategoryID,
		&product.ImageURL,
	)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (prodDB *ProductDB) GetProductByCategory(categoryID string) ([]*entity.Product, error) {
	var products []*entity.Product
	rows, err := prodDB.db.Query("SELECT * FROM products WHERE category_id = ?", categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product entity.Product
		if err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.CategoryID,
			&product.ImageURL,
		); err != nil {
			return nil, err
		}
	}
	return products, nil
}

func (prodDB *ProductDB) CreateProduct(product *entity.Product) (string, error) {
	_, err := prodDB.db.Exec("INSERT INTO products VALUES (?, ?)", product.ID, product.Name)
	if err != nil {
		return "", err
	}
	return product.ID, nil
}
