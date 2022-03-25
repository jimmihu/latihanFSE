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
		Data:       entity.UserIDResult{ID: user.ID},
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

func (u *UserUsecase) GetRoleList() dto.HttpResponse {

	RoleList, result := u.UserRepo.GetRoleList()

	if result.Error != nil {
		return dto.DBErrorResponse(result.Error)
	}

	return dto.HttpResponse{
		StatusCode: http.StatusOK,
		Status:     "ok",
		Error:      nil,
		Data:       RoleList,
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

func (u *UserUsecase) DeleteUser(ID string) dto.HttpResponse {
	uuID, _ := uuid.Parse(ID)
	result := u.UserRepo.DeleteUser(uuID)

	if result.RowsAffected == 0 {
		return dto.UserNotFoundResponse(result.Error)
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

func (u *UserUsecase) UpdateUser(ID string, request dto.UpdateUserRequest) dto.HttpResponse {
	uuID, _ := uuid.Parse(ID)
	user := entity.User{
		PersonalNumber: request.PersonalNumber,
		Name:           request.Name,
		Email:          request.Email,
		Password:       request.Password,
		Active:         request.Active,
		RoleID:         request.Role.ID,
	}
	result := u.UserRepo.UpdateUser(uuID, &user)

	if result.RowsAffected == 0 {
		return dto.UserNotFoundResponse(result.Error)
	}

	if result.Error != nil {
		return dto.DBErrorResponse(result.Error)
	}

	return dto.HttpResponse{
		StatusCode: http.StatusOK,
		Status:     "ok",
		Error:      nil,
		Data:       entity.UserIDResult{ID: uuID},
	}
}
