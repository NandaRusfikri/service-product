package product

import (
	"service-product/models"
	repositorys "service-product/repositorys/product"
	"service-product/schemas"
)

type ServiceCreate interface {
	CreateProductService(input *schemas.SchemaProduct) (*models.ModelProduct, schemas.SchemaDatabaseError)
}

type serviceCreate struct {
	repository repositorys.RepositoryCreate
}

func NewServiceCreate(repository repositorys.RepositoryCreate) *serviceCreate {
	return &serviceCreate{repository: repository}
}

func (s *serviceCreate) CreateProductService(input *schemas.SchemaProduct) (*models.ModelProduct, schemas.SchemaDatabaseError) {

	var student schemas.SchemaProduct
	student.Name = input.Name
	student.Quantity = input.Quantity
	student.IsActive = input.IsActive

	res, err := s.repository.CreateProductRepository(&student)
	return res, err
}
