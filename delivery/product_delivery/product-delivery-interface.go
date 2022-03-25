package product_delivery

import (
	"latihanFSE/usecase/product_usecase"

	"github.com/gin-gonic/gin"
)

type ProductDeliveryInterface interface {
	CreateProduct(c *gin.Context)
	GetProductList(c *gin.Context)
	GetProductDetail(c *gin.Context)
	DeleteProduct(c *gin.Context)
}

type ProductDelivery struct {
	ProductUsecase product_usecase.ProductUsecaseInterface
}

func CreateProductDelivery(ProductUsecase product_usecase.ProductUsecaseInterface) ProductDeliveryInterface {
	return &ProductDelivery{
		ProductUsecase: ProductUsecase,
	}
}
