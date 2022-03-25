package product_repository

import (
	"latihanFSE/models/entity"

	"gorm.io/gorm"
)

func (repo *ProductRepo) CreateProduct(product *entity.Product) *gorm.DB {

	product.Status = "inactive"
	result := repo.mysqlConn.Omit("MakerID", "CheckerID", "SignerID").Create(&product)
	return result
}
