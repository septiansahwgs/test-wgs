package usecase

import "test-wgs/models"

type IProductUsecase interface {
	GetAllProduct(price, name, date string) (models.Product, error)
}
