package product

import (
	"service-product/enitites"
	"service-product/schemas"
)

type ServiceProduct interface {
	CreateProductService(input *schemas.SchemaProduct) (*enitites.EntityProduct, schemas.SchemaDatabaseError)
	DeleteProductService(input *schemas.SchemaProduct) (*enitites.EntityProduct, schemas.SchemaDatabaseError)
	GetProductByIDService(id uint64) (enitites.EntityProduct, schemas.SchemaDatabaseError)
	ResultsProductService() ([]*enitites.EntityProduct, schemas.SchemaDatabaseError)
	UpdateProductService(input *schemas.SchemaProduct) (*enitites.EntityProduct, schemas.SchemaDatabaseError)
}
