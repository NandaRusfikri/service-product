package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"service-product/helpers"
	"service-product/schemas"
	services "service-product/services/product"
	"strconv"
)

type ProductControllerHTTP struct {
	service services.ServiceProduct
}

func NewControllerProductHTTP(service services.ServiceProduct) *ProductControllerHTTP {
	return &ProductControllerHTTP{service: service}
}
func (controller *ProductControllerHTTP) CreateProductHandler(ctx *gin.Context) {

	var input schemas.SchemaProduct
	ctx.ShouldBindJSON(&input)

	_, err := controller.service.CreateProductService(&input)

	switch err.Type {
	case "error_01":
		helpers.APIResponse(ctx, "Npm Product already exist", err.Code, http.MethodPost, nil)
		return
	case "error_02":
		helpers.APIResponse(ctx, "Create new Product account failed", err.Code, http.MethodPost, nil)
		return
	default:
		helpers.APIResponse(ctx, "Create new Product account successfully", http.StatusCreated, http.MethodPost, nil)
	}

}

func (controller *ProductControllerHTTP) DeleteProductHandler(ctx *gin.Context) {

	var input schemas.SchemaProduct
	input.ID, _ = strconv.ParseUint(ctx.Param("id"), 10, 64)

	_, err := controller.service.DeleteProductService(&input)

	switch err.Type {
	case "error_01":
		helpers.APIResponse(ctx, "Product data is not exist or deleted", err.Code, http.MethodDelete, nil)
		return
	case "error_02":
		helpers.APIResponse(ctx, "Delete Product data failed", err.Code, http.MethodDelete, nil)
		return
	default:
		helpers.APIResponse(ctx, "Delete Product data successfully", http.StatusOK, http.MethodDelete, nil)
	}
}

func (controller *ProductControllerHTTP) ResultProductHandler(ctx *gin.Context) {

	ID, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)

	res, err := controller.service.GetProductByIDService(ID)

	switch err.Type {
	case "error_01":
		helpers.APIResponse(ctx, "Product data is not exist or deleted", err.Code, http.MethodGet, nil)
		return
	default:
		helpers.APIResponse(ctx, "Result Product data successfully", http.StatusOK, http.MethodGet, res)
	}
}

func (controller *ProductControllerHTTP) ResultsProductHandler(ctx *gin.Context) {

	res, err := controller.service.ResultsProductService()

	switch err.Type {
	case "error_01":
		helpers.APIResponse(ctx, "Products data is not exists", err.Code, http.MethodPost, nil)
	default:
		helpers.APIResponse(ctx, "Results Products data successfully", http.StatusOK, http.MethodPost, res)
	}
}

func (controller *ProductControllerHTTP) UpdateProductHandler(ctx *gin.Context) {

	var input schemas.SchemaProduct
	input.ID, _ = strconv.ParseUint(ctx.Param("id"), 10, 64)
	ctx.ShouldBindJSON(&input)

	_, err := controller.service.UpdateProductService(&input)

	switch err.Type {
	case "error_01":
		helpers.APIResponse(ctx, "Product data is not exist or deleted", http.StatusNotFound, http.MethodPost, nil)
	case "error_02":
		helpers.APIResponse(ctx, "Update Product data failed", http.StatusForbidden, http.MethodPost, nil)
	default:
		helpers.APIResponse(ctx, "Update Product data sucessfully", http.StatusOK, http.MethodPost, nil)
	}
}
