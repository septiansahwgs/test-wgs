package repository

import (
	"test-wgs/models"

	"github.com/jmoiron/sqlx"
)

type IOrdersSql interface {
	CreateOrder(models.Orders) (*sqlx.Tx, int, error)
	CreateOrderDetail(tx *sqlx.Tx, orderDetails []models.OrderDetails) error
}
