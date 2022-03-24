package user_usecase

import (
	"latihanFSE/models/dto"
	"latihanFSE/repository/user_repository"
)

type UserUsecaseInterface interface {
	CreateUser(dto.CreateUserRequest) dto.HttpResponse
}

type UserUsecase struct {
	UserRepo user_repository.UserRepoInterface
}

func CreateUsercase(UserRepo user_repository.UserRepoInterface) UserUsecaseInterface {
	return &UserUsecase{
		UserRepo: UserRepo,
	}
}
