package user_usecase

import (
	"latihanFSE/models/entity"
	"latihanFSE/repository/user_repository"
	"latihanFSE/usecase/jwt_usecase"
	"testing"
)

func getUsecase() UserUsecaseInterface {
	repoMock := new(user_repository.UserRepositoryMock)
	JwtUsecase := jwt_usecase.CreateJwtUseCase(repoMock)
	usecase := CreateUserUsecase(repoMock, JwtUsecase)

	return usecase
}

var reqUser = entity.User{}

func TestGivenEmptyUser_WhenCreateUser_ShouldProduceBadRequestError(t *testing.T) {

}
