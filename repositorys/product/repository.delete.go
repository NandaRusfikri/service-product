package repositorys

import (
	"net/http"
	"service-product/models"
	"service-product/schemas"

	"gorm.io/gorm"


)

type RepositoryDelete interface {
	DeleteProductRepository(input *schemas.SchemaProduct) (*models.ModelProduct, schemas.SchemaDatabaseError)
}

type repositoryDelete struct {
	db *gorm.DB
}

func NewRepositoryDelete(db *gorm.DB) *repositoryDelete {
	return &repositoryDelete{db: db}
}

func (r *repositoryDelete) DeleteProductRepository(input *schemas.SchemaProduct) (*models.ModelProduct, schemas.SchemaDatabaseError) {

	var students models.ModelProduct
	db := r.db.Model(&students)
	errorCode := make(chan schemas.SchemaDatabaseError, 1)

	students.ID = input.ID
	checkStudentId := db.Debug().First(&students)

	if checkStudentId.RowsAffected < 1 {
		errorCode <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_01",
		}
		return &students, <-errorCode
	}

	deleteStudentId := db.Debug().Delete(&students)

	if deleteStudentId.RowsAffected < 1 {
		errorCode <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_02",
		}
		return &students, <-errorCode
	}
	errorCode <- schemas.SchemaDatabaseError{}

	return &students, <-errorCode
}
