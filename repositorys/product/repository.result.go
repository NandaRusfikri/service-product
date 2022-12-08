package repositorys

import (
	"net/http"
	"service-product/models"
	"service-product/schemas"

	"gorm.io/gorm"


)

type RepositoryResult interface {
	ResultProductRepository(input *schemas.SchemaProduct) (*models.ModelProduct, schemas.SchemaDatabaseError)
}

type repositoryResult struct {
	db *gorm.DB
}

func NewRepositoryResult(db *gorm.DB) *repositoryResult {
	return &repositoryResult{db: db}
}

func (r *repositoryResult) ResultProductRepository(input *schemas.SchemaProduct) (*models.ModelProduct, schemas.SchemaDatabaseError) {

	var students models.ModelProduct
	db := r.db.Model(&students)
	errorCode := make(chan schemas.SchemaDatabaseError, 1)

	resultStudents := db.Debug().First(&students)

	if resultStudents.RowsAffected < 1 {
		errorCode <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_01",
		}
		return &students, <-errorCode
	}

	return &students, <-errorCode
}
