package repositorys

import (
	"service-product/enitites"
	"service-product/schemas"
)

type ProductRepository interface {
	CreateProductRepository(input *schemas.SchemaProduct) (*enitites.EntityProduct, schemas.SchemaDatabaseError)
	DeleteProductRepository(input *schemas.SchemaProduct) (*enitites.EntityProduct, schemas.SchemaDatabaseError)
	GetProductByIDRepository(id uint64) (enitites.EntityProduct, schemas.SchemaDatabaseError)
	ListProductRepository() ([]*enitites.EntityProduct, schemas.SchemaDatabaseError)
	UpdateProductRepository(input *schemas.SchemaProduct) (*enitites.EntityProduct, schemas.SchemaDatabaseError)
}
