package database

import (
	"database/sql"

	"github.com/sidney-cardoso/goapi/internal/entity"
)

type CategoryDB struct {
	db *sql.DB
}

func NewCategoryDB(db *sql.DB) *CategoryDB {
	return &CategoryDB{
		db: db,
	}
}

func (catDB *CategoryDB) GetCategories() ([]*entity.Category, error) {
	rows, err := catDB.db.Query("SELECT * FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*entity.Category

	for rows.Next() {
		var category entity.Category

		if err := rows.Scan(
			&category.ID,
			&category.Name,
		); err != nil {
			return nil, err
		}

		categories = append(categories, &category)
	}
	return categories, nil
}

func (catDB *CategoryDB) GetCategory(id string) (*entity.Category, error) {
	var category entity.Category
	err := catDB.db.QueryRow("SELECT * FROM categories WHERE id = ?", id).Scan(
		&category.ID,
		&category.Name,
	)
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (catDB *CategoryDB) CreateCategory(category *entity.Category) (string, error) {
	_, err := catDB.db.Exec("INSERT INTO category (id, name) VALUES (?, ?)", category.ID, category.Name)
	if err != nil {
		return "", err
	}

	return category.ID, nil
}
