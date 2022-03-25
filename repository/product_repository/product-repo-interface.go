package product_repository

import (
	"latihanFSE/models/entity"

	"gorm.io/gorm"
)

type ProductRepoInterface interface {
	CreateProduct(product *entity.Product) *gorm.DB
}

type ProductRepo struct {
	mysqlConn *gorm.DB
}

func GetProductRepo(mysqlConn *gorm.DB) ProductRepoInterface {
	return &ProductRepo{
		mysqlConn: mysqlConn,
	}
}
