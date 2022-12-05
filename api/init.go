package api

import (
	"log"
	"test-wgs/api/orders"
	ordersRepo "test-wgs/api/orders/repository"
	ordersUseCase "test-wgs/api/orders/usecase"
	"test-wgs/api/product"
	productRepo "test-wgs/api/product/repository"
	productUseCase "test-wgs/api/product/usecase"
	"test-wgs/config"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	db, err := config.ConnectSQL()
	if err != nil {
		log.Fatal("error db connection ", err)
	}
	r.GET("", func(c *gin.Context) {
		return
	})

	// r.Use(utils.CORSMiddleware())
	v1 := r.Group("/api")

	iProductSql := productRepo.NewProductSQL(db)
	iProductUseCase := productUseCase.NewProductUseCase(iProductSql)
	ProductController := product.Product{UseCase: iProductUseCase}
	ProductController.User(v1)

	iOrderSql := ordersRepo.NewOrderSQL(db)
	iOrderUseCase := ordersUseCase.NewOrdersUseCase(iOrderSql, iProductSql)
	orderController := orders.Orders{UseCase: iOrderUseCase}
	orderController.Orders(v1)

	r.Run(":8002")
}
