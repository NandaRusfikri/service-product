package product

import (
	"service-product/models"
	repositorys "service-product/repositorys/product"
	"service-product/schemas"
)

type ServiceUpdate interface {
	UpdateProductService(input *schemas.SchemaProduct) (*models.ModelProduct, schemas.SchemaDatabaseError)
}

type serviceUpdate struct {
	repository repositorys.RepositoryUpdate
}

func NewServiceUpdate(repository repositorys.RepositoryUpdate) *serviceUpdate {
	return &serviceUpdate{repository: repository}
}

func (s *serviceUpdate) UpdateProductService(input *schemas.SchemaProduct) (*models.ModelProduct, schemas.SchemaDatabaseError) {

	var student schemas.SchemaProduct
	student.ID = input.ID
	student.Name = input.Name
	student.Quantity = input.Quantity
	student.IsActive = input.IsActive
	student.Price = input.Price

	res, err := s.repository.UpdateProductRepository(&student)
	return res, err
}
