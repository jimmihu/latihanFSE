package product_delivery

import (
	"encoding/json"
	"latihanFSE/models/dto"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (p *ProductDelivery) CreateProduct(c *gin.Context) {
	d := json.NewDecoder(c.Request.Body)
	d.DisallowUnknownFields()
	UserID, _ := c.Get("user_id")
	UserMakerID, _ := UserID.(uuid.UUID)

	CreateProductRequest := dto.CreateProductRequest{
		MakerID: UserMakerID,
	}
	err := d.Decode(&CreateProductRequest)
	if err != nil {
		dto.JsonRequestErrorResponse(c, err)
		return
	}

	response := p.ProductUsecase.CreateProduct(CreateProductRequest)
	c.JSON(response.StatusCode, response)
}

func (p *ProductDelivery) GetProductList(c *gin.Context) {
	response := p.ProductUsecase.GetProductList()
	c.JSON(response.StatusCode, response)
}

func (p *ProductDelivery) GetProductDetail(c *gin.Context) {
	ID := c.Param("id")
	response := p.ProductUsecase.GetProductDetail(ID)
	c.JSON(response.StatusCode, response)
}

func (p *ProductDelivery) DeleteProduct(c *gin.Context) {
	ID := c.Param("id")
	response := p.ProductUsecase.DeleteProduct(ID)
	c.JSON(response.StatusCode, response)
}

func (p *ProductDelivery) UpdateProduct(c *gin.Context) {
	ID := c.Param("id")
	d := json.NewDecoder(c.Request.Body)
	d.DisallowUnknownFields()

	UpdateProductRequest := dto.UpdateProductRequest{}
	err := d.Decode(&UpdateProductRequest)
	if err != nil {
		dto.JsonRequestErrorResponse(c, err)
		return
	}

	response := p.ProductUsecase.UpdateProduct(ID, UpdateProductRequest)
	c.JSON(response.StatusCode, response)
}

func (p *ProductDelivery) CheckProduct(c *gin.Context) {
	ID := c.Param("id")
	d := json.NewDecoder(c.Request.Body)
	d.DisallowUnknownFields()
	UserID, _ := c.Get("user_id")
	CheckerID, _ := UserID.(uuid.UUID)

	CheckProductRequest := dto.CheckProductRequest{
		CheckerID: CheckerID,
	}
	err := d.Decode(&CheckProductRequest)
	if err != nil {
		dto.JsonRequestErrorResponse(c, err)
		return
	}

	response := p.ProductUsecase.CheckProduct(ID, CheckProductRequest)
	c.JSON(response.StatusCode, response)
}

func (p *ProductDelivery) PublishProduct(c *gin.Context) {
	ID := c.Param("id")
	d := json.NewDecoder(c.Request.Body)
	d.DisallowUnknownFields()
	UserID, _ := c.Get("user_id")
	SignerID, _ := UserID.(uuid.UUID)

	PublishProductRequest := dto.PublishProductRequest{
		SignerID: SignerID,
	}
	err := d.Decode(&PublishProductRequest)
	if err != nil {
		dto.JsonRequestErrorResponse(c, err)
		return
	}

	response := p.ProductUsecase.PublishProduct(ID, PublishProductRequest)
	c.JSON(response.StatusCode, response)
}
