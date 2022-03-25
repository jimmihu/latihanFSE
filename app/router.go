package app

import (
	"latihanFSE/delivery/user_delivery"
	"latihanFSE/repository/user_repository"
	"latihanFSE/usecase/user_usecase"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRouter(mysqlConn *gorm.DB) *gin.Engine {

	UserRepo := user_repository.GetUserRepo(mysqlConn)
	UserUsecase := user_usecase.CreateUsercase(UserRepo)
	UserDelivery := user_delivery.CreateDelivery(UserUsecase)

	router := gin.Default()
	router.Use(cors.Default())
	router.POST("/users", UserDelivery.CreateUser)
	router.GET("/users", UserDelivery.GetUserList)
	router.GET("/roles", UserDelivery.GetRoleList)
	router.GET("/users/:id", UserDelivery.GetUserDetail)
	router.PUT("/users/:id", UserDelivery.UpdateUser)
	router.DELETE("/users/:id", UserDelivery.DeleteUser)
	return router
}
