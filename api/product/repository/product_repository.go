package repository

import "test-wgs/models"

type IProductSql interface {
	GetProductBySKU(SKU string) (models.Product, error)
	GetProductBySKUs(SKUs []string) ([]models.Product, error)
}
