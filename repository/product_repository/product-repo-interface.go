package product_repository

import (
	"latihanFSE/models/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductRepoInterface interface {
	CreateProduct(product *entity.Product) *gorm.DB
	GetProductList() ([]entity.ProductList, *gorm.DB)
	GetProductDetail(ID uuid.UUID) (entity.ProductDetail, *gorm.DB)
	DeleteProduct(ID uuid.UUID) *gorm.DB
}

type ProductRepo struct {
	mysqlConn *gorm.DB
}

func GetProductRepo(mysqlConn *gorm.DB) ProductRepoInterface {
	return &ProductRepo{
		mysqlConn: mysqlConn,
	}
}
