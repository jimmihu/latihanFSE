package product_usecase

import (
	"latihanFSE/models/dto"
	"latihanFSE/models/entity"
	"net/http"
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
		Data:       entity.ResultUserId{ID: product.ID},
	}
}
