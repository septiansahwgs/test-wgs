package orders

import (
	"net/http"
	"test-wgs/api/orders/dto"
	"test-wgs/api/orders/usecase"

	"github.com/gin-gonic/gin"
)

type Orders struct {
	UseCase usecase.IOrderUsecase
}

func (r Orders) Orders(route *gin.RouterGroup) {
	route.POST("/order", r.CreateOrder)
}

func (o Orders) CreateOrder(ctx *gin.Context) {
	var orderDTO dto.CreateOrderDTO
	if errBind := ctx.BindJSON(&orderDTO); errBind != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": errBind.Error()})
		return
	}

	totalPrice, err := o.UseCase.CreateOrder(orderDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"total_price": totalPrice,
	})
}
