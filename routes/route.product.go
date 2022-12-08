package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	handlers "service-product/handlers/product"
	repositorys "service-product/repositorys/product"
	services "service-product/services/product"
)

func InitStudentRoutes(db *gorm.DB, route *gin.Engine) {

	/**
	@description All Handler Student
	*/
	createStudentRepository := repositorys.NewRepositoryCreate(db)
	createStudentService := services.NewServiceCreate(createStudentRepository)
	createStudentHandler := handlers.NewHandlerCreateProduct(createStudentService)

	resultsStudentRepository := repositorys.NewRepositoryResults(db)
	resultsStudentService := services.NewServiceResults(resultsStudentRepository)
	resultsStudentHandler := handlers.NewHandlerResultsProduct(resultsStudentService)

	resultStudentRepository := repositorys.NewRepositoryResult(db)
	resultStudentService := services.NewServiceResult(resultStudentRepository)
	resultStudentHandler := handlers.NewHandlerResultProduct(resultStudentService)

	deleteStudentRepository := repositorys.NewRepositoryDelete(db)
	deleteStudentService := services.NewServiceDelete(deleteStudentRepository)
	deleteStudentHandler := handlers.NewHandlerDeleteProduct(deleteStudentService)

	updateStudentRepository := repositorys.NewRepositoryUpdate(db)
	updateStudentService := services.NewServiceUpdate(updateStudentRepository)
	updateStudentHandler := handlers.NewHandlerUpdateProduct(updateStudentService)






	/**
	@description All Student Route
	*/
	groupRoute := route.Group("/api/v1")
	groupRoute.POST("/product", createStudentHandler.CreateProductHandler)
	groupRoute.GET("/product", resultsStudentHandler.ResultsProductHandler)
	groupRoute.GET("/product/:id", resultStudentHandler.ResultProductHandler)
	groupRoute.DELETE("/product/:id", deleteStudentHandler.DeleteProductHandler)
	groupRoute.PUT("/product/:id", updateStudentHandler.UpdateProductHandler)


}
