package dto

type CreateOrderDTO struct {
	UserID      string              `json:"user_id"`
	TotalPrice  float64             `json:"total_price"`
	OrderDetail []CreateOrderDetail `json:"order_detail"`
}

type CreateOrderDetail struct {
	ProductSKU string `json:"product_sku"`
	ProductQty int    `json:"product_qty"`
}
