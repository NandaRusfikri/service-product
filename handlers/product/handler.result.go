package handlers

import (
	"net/http"
	"service-product/helpers"
	"service-product/schemas"
	services "service-product/services/product"

	"github.com/gin-gonic/gin"
)

type HandlerResult struct {
	service services.ServiceResult
}

func NewHandlerResultProduct(service services.ServiceResult) *HandlerResult {
	return &HandlerResult{service: service}
}

func (h *HandlerResult) ResultProductHandler(ctx *gin.Context) {

	var input schemas.SchemaProduct
	input.ID = ctx.Param("id")



	res, err := h.service.ResultProductService(&input)

	switch err.Type {
	case "error_01":
		helpers.APIResponse(ctx, "Product data is not exist or deleted", err.Code, http.MethodGet, nil)
		return
	default:
		helpers.APIResponse(ctx, "Result Product data successfully", http.StatusOK, http.MethodGet, res)
	}
}

//func (h *HandlerResult) GetProduct(ctx *gin.Context) {
//
//	var input schemas.SchemaProduct
//	input.ID = ctx.Param("id")
//
//}
