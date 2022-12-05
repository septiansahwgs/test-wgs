package usecase

import (
	"test-wgs/api/product/repository"
	"test-wgs/models"
)

type productUseCase struct {
	productSql repository.IProductSql
}

func NewProductUseCase(productSql repository.IProductSql) IProductUsecase {
	return &productUseCase{
		productSql: productSql,
	}
}

func (prd productUseCase) GetAllProduct(price, name, date string) (models.Product, error) {

	var orderBy string

	if len(price) > 0 {
		if price == "high" {
			orderBy = "ORDER BY product_price DESC"
		} else if price == "low" {
			orderBy = "ORDER BY product_price ASC"
		}

	}

	if len(name) > 0 {
		if name == "ASC" {
			orderBy = "ORDER BY product_name ASC"
		} else if name == "DESC" {
			orderBy = "ORDER BY product_name DESC"
		}

	}

	if len(date) > 0 {
		if date == "latest" {
			orderBy = "ORDER BY created_date DESC"
		} else if date == "longest" {
			orderBy = "ORDER BY created_date ASC"
		}

	}

	data, err := prd.productSql.GetProductBySKU(orderBy)
	if err != nil {
		return models.Product{}, err
	}

	return data, nil
}
