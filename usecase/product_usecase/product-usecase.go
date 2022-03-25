package product_usecase

import (
	"latihanFSE/models/dto"
	"latihanFSE/models/entity"
	"net/http"

	"github.com/google/uuid"
)

func (p *ProductUsecase) CreateProduct(request dto.CreateProductRequest) dto.HttpResponse {

	product := entity.Product{
		Name:        request.Name,
		Description: request.Description,
	}

	result := p.ProductRepo.CreateProduct(&product)

	if result.Error != nil {
		return dto.DBErrorResponse(result.Error)
	}

	return dto.HttpResponse{
		StatusCode: http.StatusCreated,
		Status:     "ok",
		Error:      nil,
		Data:       entity.ProductIDResult{ID: product.ID},
	}
}

func (p *ProductUsecase) GetProductList() dto.HttpResponse {

	ProductList, result := p.ProductRepo.GetProductList()

	if result.Error != nil {
		return dto.DBErrorResponse(result.Error)
	}

	return dto.HttpResponse{
		StatusCode: http.StatusOK,
		Status:     "ok",
		Error:      nil,
		Data:       ProductList,
	}
}

func (p *ProductUsecase) GetProductDetail(ID string) dto.HttpResponse {
	uuID, _ := uuid.Parse(ID)
	ProductDetail, result := p.ProductRepo.GetProductDetail(uuID)

	if result.Error != nil {
		return dto.ProductNotFoundResponse(result.Error)
	}

	return dto.HttpResponse{
		StatusCode: http.StatusOK,
		Status:     "ok",
		Error:      nil,
		Data:       ProductDetail,
	}

}

func (p *ProductUsecase) DeleteProduct(ID string) dto.HttpResponse {
	uuID, _ := uuid.Parse(ID)
	result := p.ProductRepo.DeleteProduct(uuID)

	if result.RowsAffected == 0 {
		return dto.ProductNotFoundResponse(result.Error)
	}

	if result.Error != nil {
		return dto.DBErrorResponse(result.Error)
	}

	return dto.HttpResponse{
		StatusCode: http.StatusOK,
		Status:     "ok",
		Error:      nil,
		Data:       nil,
	}
}

func (p *ProductUsecase) UpdateProduct(ID string, request dto.UpdateProductRequest) dto.HttpResponse {
	uuID, _ := uuid.Parse(ID)
	product := entity.Product{
		Name:        request.Name,
		Description: request.Description,
	}
	result := p.ProductRepo.UpdateProduct(uuID, &product)

	if result.RowsAffected == 0 {
		return dto.ProductNotFoundResponse(result.Error)
	}

	if result.Error != nil {
		return dto.DBErrorResponse(result.Error)
	}

	return dto.HttpResponse{
		StatusCode: http.StatusOK,
		Status:     "ok",
		Error:      nil,
		Data:       entity.ProductIDResult{ID: uuID},
	}
}

func (p *ProductUsecase) CheckProduct(ID string, request dto.UpdateProductRequest) dto.HttpResponse {
	uuID, _ := uuid.Parse(ID)
	product := entity.Product{
		Name:        request.Name,
		Description: request.Description,
	}
	result := p.ProductRepo.CheckProduct(uuID, &product)

	if result.RowsAffected == 0 {
		return dto.ProductNotFoundResponse(result.Error)
	}

	if result.Error != nil {
		return dto.DBErrorResponse(result.Error)
	}

	return dto.HttpResponse{
		StatusCode: http.StatusOK,
		Status:     "ok",
		Error:      nil,
		Data:       entity.ProductIDResult{ID: uuID},
	}
}

func (p *ProductUsecase) PublishProduct(ID string, request dto.UpdateProductRequest) dto.HttpResponse {
	uuID, _ := uuid.Parse(ID)
	product := entity.Product{
		Name:        request.Name,
		Description: request.Description,
	}
	result := p.ProductRepo.PublishProduct(uuID, &product)

	if result.RowsAffected == 0 {
		return dto.ProductNotFoundResponse(result.Error)
	}

	if result.Error != nil {
		return dto.DBErrorResponse(result.Error)
	}

	return dto.HttpResponse{
		StatusCode: http.StatusOK,
		Status:     "ok",
		Error:      nil,
		Data:       entity.ProductIDResult{ID: uuID},
	}
}
