package repositorys

import (
	"net/http"
	"service-product/models"
	"service-product/schemas"

	"gorm.io/gorm"


)

type RepositoryResults interface {
	ResultsProductRepository() ([]*models.ModelProduct, schemas.SchemaDatabaseError)
}

type repositoryResults struct {
	db *gorm.DB
}

func NewRepositoryResults(db *gorm.DB) *repositoryResults {
	return &repositoryResults{db: db}
}

func (r *repositoryResults) ResultsProductRepository() ([]*models.ModelProduct, schemas.SchemaDatabaseError) {

	var students []*models.ModelProduct
	db := r.db.Model(&students)
	errorCode := make(chan schemas.SchemaDatabaseError, 1)


	resultsStudents := db.Debug().Find(&students)

	if resultsStudents.RowsAffected < 1 {
		errorCode <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_01",
		}
		return students, <-errorCode
	}


	errorCode <- schemas.SchemaDatabaseError{}
	return students, <-errorCode
}
