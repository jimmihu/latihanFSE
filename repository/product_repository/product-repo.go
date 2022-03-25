package product_repository

import (
	"latihanFSE/models/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (repo *ProductRepo) CreateProduct(product *entity.Product) *gorm.DB {

	product.Status = "inactive"
	result := repo.mysqlConn.Omit("MakerID", "CheckerID", "SignerID").Create(&product)
	return result
}

func (repo *ProductRepo) GetProductList() ([]entity.ProductList, *gorm.DB) {
	ProductListResult := []entity.ProductList{}
	Result := repo.mysqlConn.Model(&entity.Product{}).Find(&ProductListResult)
	return ProductListResult, Result
}

func (repo *ProductRepo) GetProductDetail(ID uuid.UUID) (entity.ProductDetail, *gorm.DB) {
	ProductDetailResult := entity.ProductDetail{}
	Result := repo.mysqlConn.Model(&entity.Product{}).Joins("Maker").Joins("Checker").Joins("Signer").First(&ProductDetailResult, ID)
	return ProductDetailResult, Result
}

func (repo *ProductRepo) DeleteProduct(ID uuid.UUID) *gorm.DB {
	Result := repo.mysqlConn.Delete(&entity.Product{}, ID)
	return Result
}

func (repo *ProductRepo) UpdateProduct(ID uuid.UUID, product *entity.Product) *gorm.DB {

	Result := repo.mysqlConn.Model(entity.Product{ID: ID}).Updates(&product)
	return Result
}

func (repo *ProductRepo) CheckProduct(ID uuid.UUID, product *entity.Product) *gorm.DB {
	product.Status = "approved"
	Result := repo.mysqlConn.Model(entity.Product{ID: ID}).Updates(&product)
	return Result
}

func (repo *ProductRepo) PublishProduct(ID uuid.UUID, product *entity.Product) *gorm.DB {
	product.Status = "active"
	Result := repo.mysqlConn.Model(entity.Product{ID: ID}).Updates(&product)
	return Result
}
