package product

import (
	"net/http"
	"test-wgs/api/product/usecase"

	"github.com/gin-gonic/gin"
)

type Product struct {
	UseCase usecase.IProductUsecase
}

func (r Product) User(route *gin.RouterGroup) {
	route.GET("/product", r.GetAllProduct)
}

func (r Product) GetAllProduct(c *gin.Context) {
	price := c.Query("price")
	name := c.Query("name")
	date := c.Query("date")

	data, err := r.UseCase.GetAllProduct(price, name, date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}
