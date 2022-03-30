package jwt_usecase

import (
	"latihanFSE/models/entity"
	"latihanFSE/repository/user_repository"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type JwtUsecaseInterface interface {
	GenerateToken(UserID uuid.UUID, Name string) (string, error)
	ValidateToken(authHeader string) (*jwt.Token, error)
	ValidateTokenAndGetPayload(token string) (*entity.JwtPayload, error)
	UserHasAuthorization(userID uuid.UUID, roles []string) (bool, error)
}

type JwtUsecase struct {
	UserRepo user_repository.UserRepoInterface
}

func CreateJwtUseCase(UserRepo user_repository.UserRepoInterface) JwtUsecaseInterface {
	return &JwtUsecase{
		UserRepo: UserRepo,
	}
}
