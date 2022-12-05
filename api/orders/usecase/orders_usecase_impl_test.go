package usecase

import (
	"log"
	"test-wgs/api/orders/dto"
	"test-wgs/api/orders/repository"
	productRepo "test-wgs/api/product/repository"
	"test-wgs/config"
	"testing"

	"github.com/magiconair/properties/assert"
)

func Test_orderUseCase_CreateOrder(t *testing.T) {
	db, err := config.ConnectSQLTest()
	if err != nil {
		log.Fatal("error db connection ", err)
	}

	type fields struct {
		orderSql   repository.IOrdersSql
		productSQl productRepo.IProductSql
	}
	type args struct {
		orderDTO dto.CreateOrderDTO
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		expected float64
	}{
		// TODO: Add test cases.
		{
			name:   "Bundling MacBook",
			fields: fields{orderSql: repository.NewOrderSQL(db), productSQl: productRepo.NewProductSQL(db)},
			args: args{orderDTO: dto.CreateOrderDTO{
				UserID:     "USR-001",
				TotalPrice: 5399.99,
				OrderDetail: []dto.CreateOrderDetail{
					{
						ProductSKU: "43N23P",
						ProductQty: 1,
					},
				},
			}},
			expected: 5399.99,
		},
		{
			name:   "Google Home for The Price of 2",
			fields: fields{orderSql: repository.NewOrderSQL(db), productSQl: productRepo.NewProductSQL(db)},
			args: args{orderDTO: dto.CreateOrderDTO{
				UserID:     "USR-002",
				TotalPrice: 149.97, //harga total ketika quantity product di kali 3
				OrderDetail: []dto.CreateOrderDetail{
					{
						ProductSKU: "120P90",
						ProductQty: 3,
					},
				},
			}},
			expected: 99.98,
		},
		{
			name:   "3 Alexa Speakers will have a 10 discount ",
			fields: fields{orderSql: repository.NewOrderSQL(db), productSQl: productRepo.NewProductSQL(db)},
			args: args{orderDTO: dto.CreateOrderDTO{
				UserID:     "USR-003",
				TotalPrice: 328.5, //harga total ketika quantity product di kali 3
				OrderDetail: []dto.CreateOrderDetail{
					{
						ProductSKU: "A304SD",
						ProductQty: 3,
					},
				},
			}},
			expected: 295.65,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ord := orderUseCase{
				orderSql:   tt.fields.orderSql,
				productSQl: tt.fields.productSQl,
			}
			totalPrice, _ := ord.CreateOrder(tt.args.orderDTO)
			assert.Equal(t, tt.expected, totalPrice)
		})
	}
}
