package user_delivery

import (
	"latihanFSE/usecase/user_usecase"

	"github.com/gin-gonic/gin"
)

type UserDeliveryInterface interface {
	CreateUser(*gin.Context)
}

type UserDelivery struct {
	UserUsecase user_usecase.UserUsecaseInterface
}

func CreateDelivery(UserUsecase user_usecase.UserUsecaseInterface) UserDeliveryInterface {
	return &UserDelivery{
		UserUsecase: UserUsecase,
	}
}
