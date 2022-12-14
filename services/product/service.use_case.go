package product

import (
	"service-product/enitites"
	repositorys "service-product/repositorys/product"
	"service-product/schemas"
)

type serviceProductImpl struct {
	repository repositorys.ProductRepository
}

func NewServiceProduct(repository repositorys.ProductRepository) *serviceProductImpl {
	return &serviceProductImpl{repository: repository}
}

func (s *serviceProductImpl) GetProductByIDService(id uint64) (enitites.EntityProduct, schemas.SchemaDatabaseError) {

	res, err := s.repository.GetProductByIDRepository(id)
	return res, err
}

func (s *serviceProductImpl) CreateProductService(input *schemas.SchemaProduct) (*enitites.EntityProduct, schemas.SchemaDatabaseError) {

	var student schemas.SchemaProduct
	student.Name = input.Name
	student.Quantity = input.Quantity
	student.IsActive = input.IsActive
	student.Price = input.Price

	res, err := s.repository.CreateProductRepository(&student)
	return res, err
}

func (s *serviceProductImpl) DeleteProductService(input *schemas.SchemaProduct) (*enitites.EntityProduct, schemas.SchemaDatabaseError) {

	Product, err := s.repository.GetProductByIDRepository(input.ID)
	if Product.ID < 1 {
		return nil, schemas.SchemaDatabaseError{Code: 404, Message: "Product ID not found"}
	}
	var student schemas.SchemaProduct
	student.ID = input.ID
	res, err := s.repository.DeleteProductRepository(&student)
	return res, err
}

func (s *serviceProductImpl) ResultsProductService() ([]*enitites.EntityProduct, schemas.SchemaDatabaseError) {

	res, err := s.repository.ListProductRepository()
	return res, err
}

func (s *serviceProductImpl) UpdateProductService(input *schemas.SchemaProduct) (*enitites.EntityProduct, schemas.SchemaDatabaseError) {

	Product, err := s.repository.GetProductByIDRepository(input.ID)
	if Product.ID < 1 {
		return nil, schemas.SchemaDatabaseError{Code: 404, Message: "Product ID not found"}
	}

	var student schemas.SchemaProduct
	student.ID = Product.ID
	student.Name = input.Name
	student.Quantity = input.Quantity
	student.IsActive = input.IsActive
	student.Price = input.Price

	res, err := s.repository.UpdateProductRepository(&student)
	return res, err
}
