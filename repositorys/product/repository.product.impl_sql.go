package repositorys

import (
	"gorm.io/gorm"
	"net/http"
	"service-product/enitites"
	"service-product/schemas"
)

type productRepositoryImplSQL struct {
	db *gorm.DB
}

func NewProductRepositorySQL(db *gorm.DB) ProductRepository {
	return &productRepositoryImplSQL{db: db}
}

func (r *productRepositoryImplSQL) CreateProductRepository(input *schemas.SchemaProduct) (*enitites.EntityProduct, schemas.SchemaDatabaseError) {

	var product enitites.EntityProduct
	db := r.db.Model(&product)
	errorCode := make(chan schemas.SchemaDatabaseError, 1)

	checkProductExist := db.Debug().First(&product, "name = ?", input.Name)

	if checkProductExist.RowsAffected > 0 {
		errorCode <- schemas.SchemaDatabaseError{
			Code: http.StatusConflict,
			Type: "error_01",
		}
		return &product, <-errorCode
	}

	product.Name = input.Name
	product.Quantity = input.Quantity
	product.Price = input.Price
	product.IsActive = input.IsActive

	addNewProduct := db.Debug().Create(&product)

	if addNewProduct.RowsAffected < 1 {
		errorCode <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_02",
		}
		return &product, <-errorCode
	}
	errorCode <- schemas.SchemaDatabaseError{}

	return &product, <-errorCode
}

func (r *productRepositoryImplSQL) DeleteProductRepository(input *schemas.SchemaProduct) (*enitites.EntityProduct, schemas.SchemaDatabaseError) {

	var students enitites.EntityProduct
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

func (r *productRepositoryImplSQL) GetProductByIDRepository(id uint64) (enitites.EntityProduct, schemas.SchemaDatabaseError) {

	var students enitites.EntityProduct
	errorCode := make(chan schemas.SchemaDatabaseError, 1)
	db := r.db.Model(&students)
	students.ID = id
	resultProduct := db.Debug().First(&students, id)
	if resultProduct.RowsAffected < 1 {
		errorCode <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_01",
		}
		return students, <-errorCode
	}

	errorCode <- schemas.SchemaDatabaseError{}

	return students, <-errorCode
}

func (r *productRepositoryImplSQL) ListProductRepository() ([]*enitites.EntityProduct, schemas.SchemaDatabaseError) {

	var students []*enitites.EntityProduct
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

func (r *productRepositoryImplSQL) UpdateProductRepository(input *schemas.SchemaProduct) (*enitites.EntityProduct, schemas.SchemaDatabaseError) {

	var students enitites.EntityProduct
	db := r.db.Model(&students)
	errorCode := make(chan schemas.SchemaDatabaseError, 1)

	students.ID = input.ID
	students.Name = input.Name
	students.Quantity = input.Quantity
	students.IsActive = input.IsActive
	students.Price = input.Price

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
