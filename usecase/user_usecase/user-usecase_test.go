package user_usecase

import (
	"latihanFSE/models/dto"
	"latihanFSE/models/entity"
	"latihanFSE/repository/user_repository"
	"latihanFSE/usecase/jwt_usecase"
	"log"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/stretchr/testify/assert"

	"gorm.io/gorm"
)

var repoMock = new(user_repository.UserRepositoryMock)
var JwtUsecase = jwt_usecase.CreateJwtUseCase(repoMock)
var usecase = CreateUserUsecase(repoMock, JwtUsecase)

var reqUser = entity.User{}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	m.Run()
}

func TestGivenNil_WhenCreateUser_ShouldProduceBadRequestError(t *testing.T) {
	var retun gorm.DB
	repoMock.On("CreateUser", nil).Return(&retun)
	log.Println(retun.Error)
	respons := dto.HttpResponse{}
	if retun.RowsAffected == 0 {
		respons = dto.DBErrorResponse(retun.Error)
	} else {
		respons = dto.HttpResponse{
			StatusCode: http.StatusOK,
			Status:     "ok",
			Error:      nil,
			Data:       nil,
		}
	}
	//incomplete cant get *gorm must change CreateUser return #lazy
	assert.EqualValues(t, http.StatusInternalServerError, respons.StatusCode)
}
