package usecase

import (
	"fmt"
	"test-wgs/api/orders/dto"
	"test-wgs/api/orders/repository"
	productRepo "test-wgs/api/product/repository"
	"test-wgs/models"
)

type orderUseCase struct {
	orderSql   repository.IOrdersSql
	productSQl productRepo.IProductSql
}

func NewOrdersUseCase(orderSql repository.IOrdersSql, productSql productRepo.IProductSql) IOrderUsecase {
	return &orderUseCase{
		orderSql:   orderSql,
		productSQl: productSql,
	}
}

func (ord orderUseCase) CreateOrder(orderDTO dto.CreateOrderDTO) (float64, error) {
	var totalPrice float64
	var orderDetails []models.OrderDetails
	productSKU := orderDTO.OrderDetail[0].ProductSKU
	productQty := orderDTO.OrderDetail[0].ProductQty
	fmt.Println("teees")
	products, err := ord.productSQl.GetProductBySKU(productSKU)
	if err != nil {
		return totalPrice, nil
	}

	totalPrice = products.Price

	if productSKU == "120P90" && productQty == 3 {
		totalPrice = products.Price * 2
	}

	if productSKU == "A304SD" && productQty == 3 {
		totalPrice = products.Price * 3
		afterDiscount := totalPrice - (totalPrice * 10 / 100)

		totalPrice = afterDiscount
	}

	if len(orderDTO.OrderDetail) > 1 {
		totalPrice = ord.getTotalPrices(orderDTO.OrderDetail)
	}

	tx, orderID, err := ord.orderSql.CreateOrder(models.Orders{
		User_id:     orderDTO.UserID,
		Total_price: totalPrice,
	})

	if err != nil {
		return totalPrice, err
	}

	if productSKU == "43N23P" {
		orderDetails = append(orderDetails, ord.handlerBundling(orderID))
	}

	for _, orderDetailDTO := range orderDTO.OrderDetail {
		var orderDetail models.OrderDetails
		orderDetail.Order_id = orderID
		orderDetail.Product_sku = orderDetailDTO.ProductSKU
		orderDetail.Products_qty = orderDetailDTO.ProductQty

		orderDetails = append(orderDetails, orderDetail)
	}

	if err := ord.orderSql.CreateOrderDetail(tx, orderDetails); err != nil {
		return totalPrice, err
	}

	return totalPrice, nil
}

func (ord orderUseCase) getTotalPrices(orderDetailsDTO []dto.CreateOrderDetail) float64 {
	var SKUs []string
	for _, v := range orderDetailsDTO {
		SKUs = append(SKUs, v.ProductSKU)
	}

	products, err := ord.productSQl.GetProductBySKUs(SKUs)
	if err != nil {
		return 0
	}

	var result float64
	for _, v := range products {
		result += v.Price
	}

	return result
}

func (ord orderUseCase) handlerBundling(orderID int) models.OrderDetails {
	var orderDetail models.OrderDetails

	orderDetail.Order_id = orderID
	orderDetail.Product_sku = "234324"
	orderDetail.Products_qty = 1

	return orderDetail
}
