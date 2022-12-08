package handlers

import (
	"context"
	"errors"
	"net/http"
	"service-product/helpers"
	pb "service-product/proto"
	"service-product/schemas"
	services "service-product/services/product"

	"github.com/gin-gonic/gin"
)

type HandlerCreate struct {
	service services.ServiceCreate
}

func NewHandlerCreateProduct(service services.ServiceCreate) *HandlerCreate {
	return &HandlerCreate{service: service}
}

func (h *HandlerCreate) CreateProductHandler(ctx *gin.Context) {

	var input schemas.SchemaProduct
	ctx.ShouldBindJSON(&input)


	_, err := h.service.CreateProductService(&input)

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

func (h *ServiceProductHandler) CreateProductRPC(ctx context.Context, param *pb.CreateProductRequest,res *pb.ModelProtoProduct) error {

	input:= schemas.SchemaProduct{
		IsActive: param.IsActive,
		Name: param.Name,
		Quantity: param.Quantity,
	}
	Create, err := h.serviceCreate.CreateProductService(&input)
	switch err.Code {
	case 500:
		return errors.New("Internal Server Error")
	case 409:
		return errors.New("Name Product already exist")
	case 402:
		return errors.New("Create new Product account failed")
	default:
		res.Id = Create.ID
		res.Name = Create.Name
		res.Quantity = Create.Quantity
		res.IsActive = Create.IsActive
	}
	return nil

}

