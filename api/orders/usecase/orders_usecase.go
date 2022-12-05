package usecase

import "test-wgs/api/orders/dto"

type IOrderUsecase interface {
	CreateOrder(dto.CreateOrderDTO) (float64, error)
}
