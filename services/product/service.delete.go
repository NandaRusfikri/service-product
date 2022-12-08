package product

import (
	"service-product/models"
	repositorys "service-product/repositorys/product"
	"service-product/schemas"
)

type ServiceDelete interface {
	DeleteProductService(input *schemas.SchemaProduct) (*models.ModelProduct, schemas.SchemaDatabaseError)
}

type serviceDelete struct {
	repository repositorys.RepositoryDelete
}

func NewServiceDelete(repository repositorys.RepositoryDelete) *serviceDelete {
	return &serviceDelete{repository: repository}
}

func (s *serviceDelete) DeleteProductService(input *schemas.SchemaProduct) (*models.ModelProduct, schemas.SchemaDatabaseError) {

	var student schemas.SchemaProduct
	student.ID = input.ID

	res, err := s.repository.DeleteProductRepository(&student)
	return res, err
}
