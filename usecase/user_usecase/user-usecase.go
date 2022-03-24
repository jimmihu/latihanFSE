package user_usecase

import (
	"net/http"

	"latihanFSE/models/dto"
	"latihanFSE/models/entity"
)

func (u *UserUseCase) CreateUser(request dto.CreateUserRequest) dto.HttpResponse {

	user := entity.User{
		PersonalNumber: request.PersonalNumber,
		Email:          request.Email,
		Name:           request.Name,
		Password:       request.Password,
	}

	result := u.UserRepo.CreateUser(&user)

	if result.Error != nil {
		return dto.ErrorDBResponse(result.Error)
	}

	return dto.HttpResponse{
		StatusCode: http.StatusCreated,
		Status:     "ok",
		Error:      nil,
		Data:       entity.ResultUserId{ID: user.ID},
	}
}
