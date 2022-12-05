package repository

import (
	"fmt"
	"test-wgs/config"
	"test-wgs/models"

	"github.com/jmoiron/sqlx"
)

type ProductSQL struct {
	db *sqlx.DB
}

func NewProductSQL(db *config.DB) IProductSql {
	return &ProductSQL{
		db: db.SQL,
	}
}

func (P ProductSQL) GetProductBySKU(productSKU string) (product models.Product, err error) {
	if err = P.db.Get(&product, `SELECT id, sku, "name", price, inventory_qty, created_at
	FROM products WHERE sku = $1`, productSKU); err != nil {
		fmt.Println("ERROR", err)
		return product, err
	}

	return product, nil
}

func (P ProductSQL) GetProductBySKUs(SKUs []string) ([]models.Product, error) {
	var products []models.Product
	query, args, _ := sqlx.In(`SELECT id, sku, "name", price, inventory_qty, created_at
	FROM products WHERE sku IN (?)`, SKUs)

	query = P.db.Rebind(query)
	rows, err := P.db.Queryx(query, args...)
	if err != nil {
		fmt.Println("ERROR", err)
		return nil, err
	}

	for rows.Next() {
		var product models.Product
		if errBind := rows.StructScan(&product); errBind != nil {
			fmt.Println("ERROR", errBind)
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}
