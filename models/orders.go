package models

type Orders struct {
	ID          int
	User_id     string
	Total_price float64
}

type OrderDetails struct {
	ID           int
	Order_id     int
	Product_sku  string
	Products_qty int
}
