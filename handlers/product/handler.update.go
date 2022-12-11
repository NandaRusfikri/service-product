package handlers

import (
	"context"
	"errors"
	"net/http"
	"service-product/helpers"
	pb "service-product/proto"
	"service-product/schemas"
	services "service-product/services/product"
	"strconv"

	"github.com/gin-gonic/gin"
)

type handlerUpdate struct {
	service services.ServiceUpdate
}

func NewHandlerUpdateProduct(service services.ServiceUpdate) *handlerUpdate {
	return &handlerUpdate{service: service}
}

func (h *handlerUpdate) UpdateProductHandler(ctx *gin.Context) {

	var input schemas.SchemaProduct
	input.ID, _ = strconv.ParseInt(ctx.Param("id"), 10, 64)
	ctx.ShouldBindJSON(&input)

	_, err := h.service.UpdateProductService(&input)

	switch err.Type {
	case "error_01":
		helpers.APIResponse(ctx, "Product data is not exist or deleted", http.StatusNotFound, http.MethodPost, nil)
	case "error_02":
		helpers.APIResponse(ctx, "Update Product data failed", http.StatusForbidden, http.MethodPost, nil)
	default:
		helpers.APIResponse(ctx, "Update Product data sucessfully", http.StatusOK, http.MethodPost, nil)
	}
}

func (h *ServiceProductHandler) UpdateProductRPC(ctx context.Context, req *pb.EntityProtoProduct, res *pb.EntityProtoProduct) error {

	var input schemas.SchemaProduct
	input.ID, _ = strconv.ParseInt(req.Id, 10, 64)
	input.Name = req.Name
	input.Price = req.Price
	input.Quantity = req.Quantity
	input.IsActive = req.IsActive

	Update, err := h.serviceUpdate.UpdateProductService(&input)

	switch err.Code {
	case 500:
		return errors.New("Internal Server Error")
	case 404:
		return errors.New("Product data is not exist or deleted")
	case 409:
		return errors.New("Update Product data failed")
	default:
		res.Id = strconv.FormatInt(Update.ID, 10)
		res.Name = Update.Name
		res.Price = Update.Price
		res.Quantity = Update.Quantity
		res.IsActive = Update.IsActive
	}
	return nil

}
