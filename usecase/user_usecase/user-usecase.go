package user_usecase

import (
	"net/http"

	"latihanFSE/models/dto"
	"latihanFSE/models/entity"

	"github.com/google/uuid"
)

func (u *UserUsecase) CreateUser(request dto.CreateUserRequest) dto.HttpResponse {

	user := entity.User{
		PersonalNumber: request.PersonalNumber,
		Email:          request.Email,
		Name:           request.Name,
		Password:       request.Password,
	}

	result := u.UserRepo.CreateUser(&user)

	if result.Error != nil {
		return dto.DBErrorResponse(result.Error)
	}

	return dto.HttpResponse{
		StatusCode: http.StatusCreated,
		Status:     "ok",
		Error:      nil,
		Data:       entity.ResultUserId{ID: user.ID},
	}
}

func (u *UserUsecase) GetUserList() dto.HttpResponse {

	UserList, result := u.UserRepo.GetUserList()

	if result.Error != nil {
		return dto.DBErrorResponse(result.Error)
	}

	return dto.HttpResponse{
		StatusCode: http.StatusOK,
		Status:     "ok",
		Error:      nil,
		Data:       UserList,
	}
}

func (u *UserUsecase) GetUserDetail(ID string) dto.HttpResponse {
	uuID, _ := uuid.Parse(ID)
	UserDetail, result := u.UserRepo.GetUserDetail(uuID)

	if result.Error != nil {
		return dto.UserNotFoundResponse(result.Error)
	}

	return dto.HttpResponse{
		StatusCode: http.StatusOK,
		Status:     "ok",
		Error:      nil,
		Data:       UserDetail,
	}

}
