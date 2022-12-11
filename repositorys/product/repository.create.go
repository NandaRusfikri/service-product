package repositorys

import (
	"net/http"
	"service-product/models"
	"service-product/schemas"

	"gorm.io/gorm"
)

type RepositoryCreate interface {
	CreateProductRepository(input *schemas.SchemaProduct) (*models.ModelProduct, schemas.SchemaDatabaseError)
}

type repositoryCreate struct {
	db *gorm.DB
}

func NewRepositoryCreate(db *gorm.DB) *repositoryCreate {
	return &repositoryCreate{db: db}
}

func (r *repositoryCreate) CreateProductRepository(input *schemas.SchemaProduct) (*models.ModelProduct, schemas.SchemaDatabaseError) {

	var product models.ModelProduct
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
