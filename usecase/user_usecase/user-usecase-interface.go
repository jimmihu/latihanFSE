package user_usecase

import (
	"latihanFSE/models/dto"
	"latihanFSE/repository/user_repository"
	"latihanFSE/usecase/jwt_usecase"
)

type UserUsecaseInterface interface {
	CreateUser(dto.CreateUserRequest) dto.HttpResponse
	GetUserList() dto.HttpResponse
	GetRoleList() dto.HttpResponse
	GetUserDetail(string) dto.HttpResponse
	DeleteUser(string) dto.HttpResponse
	UpdateUser(string, dto.UpdateUserRequest) dto.HttpResponse
	LoginUser(request dto.LoginRequest) dto.HttpResponse
}

type UserUsecase struct {
	UserRepo   user_repository.UserRepoInterface
	JwtUsecase jwt_usecase.JwtUsecaseInterface
}

func CreateUserUsecase(UserRepo user_repository.UserRepoInterface, JwtUsecase jwt_usecase.JwtUsecaseInterface) UserUsecaseInterface {
	return &UserUsecase{
		UserRepo:   UserRepo,
		JwtUsecase: JwtUsecase,
	}
}
