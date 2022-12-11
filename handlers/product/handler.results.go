package handlers

import (
	"context"
	"errors"
	"github.com/golang/protobuf/ptypes/empty"
	"net/http"
	"service-product/helpers"
	pb "service-product/proto"
	services "service-product/services/product"
	"strconv"

	"github.com/gin-gonic/gin"
)

type HandlerResults struct {
	service services.ServiceResults
}

func NewHandlerResultsProduct(service services.ServiceResults) *HandlerResults {
	return &HandlerResults{service: service}
}

func (h *HandlerResults) ResultsProductHandler(ctx *gin.Context) {

	res, err := h.service.ResultsProductService()

	switch err.Type {
	case "error_01":
		helpers.APIResponse(ctx, "Products data is not exists", err.Code, http.MethodPost, nil)
	default:
		helpers.APIResponse(ctx, "Results Products data successfully", http.StatusOK, http.MethodPost, res)
	}
}

func (h *ServiceProductHandler) ListProductRPC(ctx context.Context, empty *empty.Empty, res *pb.ResponseModelProductList) error {

	ListProduct, err := h.serviceResults.ResultsProductService()

	var ListProto []*pb.EntityProtoProduct

	switch err.Code {
	case 500:
		return errors.New("Internal Server Error")
	default:
		for _, product := range ListProduct {
			data := pb.EntityProtoProduct{
				Id:       strconv.FormatInt(product.ID, 10),
				Name:     product.Name,
				Quantity: product.Quantity,
				Price:    product.Price,
				IsActive: product.IsActive,
			}
			ListProto = append(ListProto, &data)
		}

	}
	res.List = ListProto

	return nil
}
