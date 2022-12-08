package product

import (
	"service-product/models"
	repositorys "service-product/repositorys/product"
	"service-product/schemas"
)

type ServiceResults interface {
	ResultsProductService() ([]*models.ModelProduct, schemas.SchemaDatabaseError)
}

type serviceResults struct {
	repository repositorys.RepositoryResults
}

func NewServiceResults(repository repositorys.RepositoryResults) *serviceResults {
	return &serviceResults{repository: repository}
}

func (s *serviceResults) ResultsProductService() ([]*models.ModelProduct, schemas.SchemaDatabaseError) {

	res, err := s.repository.ResultsProductRepository()
	//fmt.Printf("serviceResults %+v\n",res)
	return res, err
}
