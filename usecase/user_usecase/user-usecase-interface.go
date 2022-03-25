package user_usecase

import (
	"latihanFSE/models/dto"
	"latihanFSE/repository/user_repository"
)

type UserUsecaseInterface interface {
	CreateUser(dto.CreateUserRequest) dto.HttpResponse
	GetUserList() dto.HttpResponse
	GetRoleList() dto.HttpResponse
	GetUserDetail(string) dto.HttpResponse
	DeleteUser(string) dto.HttpResponse
	UpdateUser(string, dto.UpdateUserRequest) dto.HttpResponse
}

type UserUsecase struct {
	UserRepo user_repository.UserRepoInterface
}

func CreateUserUsecase(UserRepo user_repository.UserRepoInterface) UserUsecaseInterface {
	return &UserUsecase{
		UserRepo: UserRepo,
	}
}
