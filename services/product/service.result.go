package product

import (
	"service-product/models"
	repositorys "service-product/repositorys/product"
	"service-product/schemas"
)

type ServiceResult interface {
	ResultProductService(input *schemas.SchemaProduct) (*models.ModelProduct, schemas.SchemaDatabaseError)
}

type serviceResult struct {
	repository repositorys.RepositoryResult
}

func NewServiceResult(repository repositorys.RepositoryResult) *serviceResult {
	return &serviceResult{repository: repository}
}

func (s *serviceResult) ResultProductService(input *schemas.SchemaProduct) (*models.ModelProduct, schemas.SchemaDatabaseError) {

	var student schemas.SchemaProduct
	student.ID = input.ID

	res, err := s.repository.ResultProductRepository(&student)
	return res, err
}
