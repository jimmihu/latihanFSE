package product_usecase

import (
	"latihanFSE/models/dto"
	"latihanFSE/repository/product_repository"
)

type ProductUsecaseInterface interface {
	CreateProduct(request dto.CreateProductRequest) dto.HttpResponse
	GetProductList() dto.HttpResponse
	GetProductDetail(string) dto.HttpResponse
	DeleteProduct(string) dto.HttpResponse
}

type ProductUsecase struct {
	ProductRepo product_repository.ProductRepoInterface
}

func CreateProductUsecase(ProductRepo product_repository.ProductRepoInterface) ProductUsecaseInterface {
	return &ProductUsecase{
		ProductRepo: ProductRepo,
	}
}
