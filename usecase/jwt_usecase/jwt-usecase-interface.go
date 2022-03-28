package jwt_usecase

import (
	"latihanFSE/repository/user_repository"

	"github.com/google/uuid"
)

type JwtUsecaseInterface interface {
	GenerateToken(UserID uuid.UUID, Email string) (string, error)
}

type JwtUsecase struct {
	UserRepo user_repository.UserRepoInterface
}

func CreateJwtUseCase(UserRepo user_repository.UserRepoInterface) JwtUsecaseInterface {
	return &JwtUsecase{
		UserRepo: UserRepo,
	}
}
