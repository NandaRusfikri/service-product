package handlers

import (
	"context"
	"errors"
	"github.com/golang/protobuf/ptypes/empty"
	"net/http"
	"service-product/helpers"
	pb "service-product/proto"
	"service-product/schemas"
	services "service-product/services/product"

	"github.com/gin-gonic/gin"
)

type handlerDelete struct {
	service services.ServiceDelete
}

func NewHandlerDeleteProduct(service services.ServiceDelete) *handlerDelete {
	return &handlerDelete{service: service}
}

func (h *handlerDelete) DeleteProductHandler(ctx *gin.Context) {

	var input schemas.SchemaProduct
	input.ID = ctx.Param("id")



	_, err := h.service.DeleteProductService(&input)

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

func (h *ServiceProductHandler) DeleteProductRPC(ctx context.Context, req *pb.DeleteProductRequest, res *empty.Empty)  error {

	var input schemas.SchemaProduct
	input.ID = req.Id

	_, err := h.serviceDelete.DeleteProductService(&input)

	switch err.Code {
	case 500:
		return errors.New("Internal Server Error")
	case 404:
		return errors.New("Product data is not exist or deleted")
	case 409:
		return errors.New("Delete Product data failed")
	default:
		return nil
	}
}

