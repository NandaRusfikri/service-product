package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	handlers "service-product/controller/product"
	repositorys "service-product/repositorys/product"
	services "service-product/services/product"
)

func InitProductRoutes(db *gorm.DB, route *gin.Engine) {

	/**
	@description All Handler Student
	*/
	productRepository := repositorys.NewProductRepositorySQL(db)
	productService := services.NewServiceProduct(productRepository)
	productHandler := handlers.NewControllerProductHTTP(productService)

	/**
	@description All Student Route
	*/
	groupRoute := route.Group("/api/v1")
	groupRoute.POST("/product", productHandler.CreateProductHandler)
	groupRoute.GET("/product", productHandler.ResultsProductHandler)
	groupRoute.GET("/product/:id", productHandler.ResultProductHandler)
	groupRoute.DELETE("/product/:id", productHandler.DeleteProductHandler)
	groupRoute.PUT("/product/:id", productHandler.UpdateProductHandler)

}
