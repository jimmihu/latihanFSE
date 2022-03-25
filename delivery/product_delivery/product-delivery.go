package product_delivery

import (
	"encoding/json"
	"latihanFSE/models/dto"

	"github.com/gin-gonic/gin"
)

func (p *ProductDelivery) CreateProduct(c *gin.Context) {
	d := json.NewDecoder(c.Request.Body)
	d.DisallowUnknownFields()

	CreateProductRequest := dto.CreateProductRequest{}
	err := d.Decode(&CreateProductRequest)
	if err != nil {
		dto.JsonRequestErrorResponse(c, err)
		return
	}

	response := p.ProductUsecase.CreateProduct(CreateProductRequest)
	c.JSON(response.StatusCode, response)
}
