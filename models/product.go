package models

import "time"

type Product struct {
	ID            int       `json:"id"`
	SKU           string    `json:"sku"`
	Name          string    `json:"name"`
	Price         float64   `json:"price"`
	Inventory_qty int       `json:"inventory_quantity"`
	Created_at    time.Time `json:"created_at"`
}
