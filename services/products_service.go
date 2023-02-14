package services

import (
	"github.com/kennedybaraka/app-server/models"
	"gorm.io/gorm"
)

type ProductService interface {
	CreateProduct(db *gorm.DB, product models.Product) (models.Product, error)
}

func CreateProduct(db *gorm.DB, product models.Product) (models.Product, error) {
	err := db.Create(&product).Error
	if err != nil {
		return models.Product{}, err
	}
	return product, nil
}
