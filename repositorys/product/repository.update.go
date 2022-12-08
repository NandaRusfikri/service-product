package repositorys

import (
	"net/http"
	"service-product/models"
	"service-product/schemas"

	"gorm.io/gorm"


)

type RepositoryUpdate interface {
	UpdateProductRepository(input *schemas.SchemaProduct) (*models.ModelProduct, schemas.SchemaDatabaseError)
}

type repositoryUpdate struct {
	db *gorm.DB
}

func NewRepositoryUpdate(db *gorm.DB) *repositoryUpdate {
	return &repositoryUpdate{db: db}
}

func (r *repositoryUpdate) UpdateProductRepository(input *schemas.SchemaProduct) (*models.ModelProduct, schemas.SchemaDatabaseError) {

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

	students.Name = input.Name
	students.Quantity = input.Quantity
	students.IsActive = input.IsActive
	students.Price= input.Price
	//students.Bid = input.Bid

	updateStudent := db.Debug().Where("id = ?", input.ID).Updates(students)

	if updateStudent.RowsAffected < 1 {
		errorCode <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_02",
		}
		return &students, <-errorCode
	}
	errorCode <- schemas.SchemaDatabaseError{}
	return &students, <-errorCode
}
