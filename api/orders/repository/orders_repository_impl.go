package repository

import (
	"fmt"
	"test-wgs/config"
	"test-wgs/models"

	"github.com/jmoiron/sqlx"
)

type OrderSQL struct {
	db *sqlx.DB
}

func NewOrderSQL(db *config.DB) IOrdersSql {
	return &OrderSQL{
		db: db.SQL,
	}
}

func (r OrderSQL) CreateOrder(order models.Orders) (*sqlx.Tx, int, error) {
	var id int
	tx := r.db.MustBegin()

	execOrder, err := tx.NamedQuery(CREATE_ORDER, map[string]interface{}{
		"user_id":     order.User_id,
		"total_price": order.Total_price,
	})

	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return tx, 0, err
	}

	defer execOrder.Close()

	for execOrder.Next() {
		if err := execOrder.Scan(&id); err != nil {
			return nil, 0, err
		}
	}

	fmt.Println("OrderID", id)
	return tx, id, nil
}

func (r OrderSQL) CreateOrderDetail(tx *sqlx.Tx, orderDetails []models.OrderDetails) error {

	_, err := tx.NamedExec(CREATE_ORDER_DETAIL, orderDetails)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
