package handlers

import (
	"context"
	"errors"
	"github.com/golang/protobuf/ptypes/empty"
	pb "service-product/proto"
	"service-product/schemas"
	services "service-product/services/product"
	"strconv"
)

type ProductControllerGRPC struct {
	service services.ServiceProduct
}

func NewControllerProductRPC(
	service services.ServiceProduct) *ProductControllerGRPC {
	return &ProductControllerGRPC{
		service: service,
	}
}

func (controller *ProductControllerGRPC) CreateProductRPC(ctx context.Context, param *pb.CreateProductRequest, res *pb.EntityProtoProduct) error {

	input := schemas.SchemaProduct{
		IsActive: param.IsActive,
		Name:     param.Name,
		Quantity: param.Quantity,
		Price:    param.Price,
	}
	Create, err := controller.service.CreateProductService(&input)

	switch err.Code {
	case 500:
		return errors.New("Internal Server Error")
	case 409:
		return errors.New("Name Product already exist")
	case 402:
		return errors.New("Create new Product account failed")
	default:
		res.Id = strconv.Itoa(int(Create.ID))
		res.Name = Create.Name
		res.Quantity = Create.Quantity
		res.IsActive = Create.IsActive
		res.Price = Create.Price
	}
	return nil

}

func (controller *ProductControllerGRPC) DeleteProductRPC(ctx context.Context, req *pb.DeleteProductRequest, res *empty.Empty) error {

	var input schemas.SchemaProduct
	input.ID, _ = strconv.ParseInt(req.Id, 10, 64)

	_, err := controller.service.DeleteProductService(&input)

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

func (controller *ProductControllerGRPC) ListProductRPC(ctx context.Context, empty *empty.Empty, res *pb.ResponseEntityProductList) error {

	ListProduct, err := controller.service.ResultsProductService()

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

func (controller *ProductControllerGRPC) UpdateProductRPC(ctx context.Context, req *pb.EntityProtoProduct, res *pb.EntityProtoProduct) error {

	var input schemas.SchemaProduct
	input.ID, _ = strconv.ParseInt(req.Id, 10, 64)
	input.Name = req.Name
	input.Price = req.Price
	input.Quantity = req.Quantity
	input.IsActive = req.IsActive

	Update, err := controller.service.UpdateProductService(&input)

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
