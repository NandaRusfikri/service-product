package handlers

import (
	"context"
	"errors"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/emptypb"
	"net/http"
	pb "service-product/proto/product"
	"service-product/schemas"
	services "service-product/services/product"
	"strconv"
)

type ServiceProductHandlerServer struct {
	service services.ServiceProduct
	pb.UnimplementedServiceProductHandlerServer
}

func NewControllerProductRPC(service services.ServiceProduct) *ServiceProductHandlerServer {
	return &ServiceProductHandlerServer{
		service: service,
	}
}

func (controller *ServiceProductHandlerServer) GetProductByIdRPC(ctx context.Context, param *pb.ProductByIdRequest) (res *pb.EntityProtoProduct, err error) {

	id, ErrParse := strconv.ParseUint(param.Id, 10, 64)
	if ErrParse != nil {
		return res, errors.New("Id must be number")
	}

	Product, errProduct := controller.service.GetProductByIDService(id)

	switch errProduct.Code {
	case 500:
		return res, errors.New("Internal Server Error")
	case http.StatusNotFound:
		return res, nil
	default:
		res.Id = strconv.Itoa(int(Product.ID))
		res.Name = Product.Name
		res.Quantity = Product.Quantity
		res.IsActive = Product.IsActive
		res.Price = Product.Price
	}
	return res, nil

}

func (controller *ServiceProductHandlerServer) CreateProductRPC(ctx context.Context, param *pb.CreateProductRequest) (*pb.EntityProtoProduct, error) {

	var resp pb.EntityProtoProduct
	input := schemas.SchemaProduct{
		IsActive: param.IsActive,
		Name:     param.Name,
		Quantity: param.Quantity,
		Price:    param.Price,
	}
	Create, err := controller.service.CreateProductService(&input)

	switch err.Code {
	case 500:
		return nil, errors.New("Internal Server Error")
	case 409:
		return nil, errors.New("Name Product already exist")
	case 402:
		return nil, errors.New("Create new Product account failed")
	default:
		resp.Id = strconv.Itoa(int(Create.ID))
		resp.Name = Create.Name
		resp.Quantity = Create.Quantity
		resp.IsActive = Create.IsActive
		resp.Price = Create.Price
	}
	return &resp, nil

}

func (controller *ServiceProductHandlerServer) DeleteProductRPC(ctx context.Context, req *pb.ProductByIdRequest) (*emptypb.Empty, error) {

	var input schemas.SchemaProduct
	input.ID, _ = strconv.ParseUint(req.Id, 10, 64)

	_, err := controller.service.DeleteProductService(&input)

	switch err.Code {
	case 500:
		return nil, errors.New("Internal Server Error")
	case 404:
		return nil, errors.New("Product data is not exist or deleted")
	case 409:
		return nil, errors.New("Delete Product data failed")
	default:
		return nil, nil
	}
}

func (controller *ServiceProductHandlerServer) ListProductRPC(ctx context.Context, empty *empty.Empty) (*pb.ResponseEntityProductList, error) {

	var res pb.ResponseEntityProductList
	ListProduct, err := controller.service.ResultsProductService()

	var ListProto []*pb.EntityProtoProduct

	switch err.Code {
	case 500:
		return nil, errors.New("Internal Server Error")
	default:
		for _, product := range ListProduct {
			data := pb.EntityProtoProduct{
				Id:       strconv.FormatUint(product.ID, 10),
				Name:     product.Name,
				Quantity: product.Quantity,
				Price:    product.Price,
				IsActive: product.IsActive,
			}
			ListProto = append(ListProto, &data)
		}

	}
	res.List = ListProto

	return &res, nil
}

func (controller *ServiceProductHandlerServer) UpdateProductRPC(ctx context.Context, req *pb.EntityProtoProduct) (*pb.EntityProtoProduct, error) {

	var input schemas.SchemaProduct
	var resp pb.EntityProtoProduct
	input.ID, _ = strconv.ParseUint(req.Id, 10, 64)
	input.Name = req.Name
	input.Price = req.Price
	input.Quantity = req.Quantity
	input.IsActive = req.IsActive

	Update, err := controller.service.UpdateProductService(&input)

	switch err.Code {
	case 500:
		return nil, errors.New("Internal Server Error")
	case 404:
		return nil, errors.New("Product data is not exist or deleted")
	case 409:
		return nil, errors.New("Update Product data failed")
	default:
		resp.Id = strconv.FormatUint(Update.ID, 10)
		resp.Name = Update.Name
		resp.Price = Update.Price
		resp.Quantity = Update.Quantity
		resp.IsActive = Update.IsActive
	}
	return &resp, nil

}
