package app

import (
	"latihanFSE/delivery/product_delivery"
	"latihanFSE/delivery/user_delivery"
	"latihanFSE/repository/product_repository"
	"latihanFSE/repository/user_repository"
	"latihanFSE/usecase/product_usecase"
	"latihanFSE/usecase/user_usecase"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRouter(mysqlConn *gorm.DB) *gin.Engine {

	UserRepo := user_repository.GetUserRepo(mysqlConn)
	UserUsecase := user_usecase.CreateUserUsecase(UserRepo)
	UserDelivery := user_delivery.CreateUserDelivery(UserUsecase)

	ProductRepo := product_repository.GetProductRepo(mysqlConn)
	ProductUsecase := product_usecase.CreateProductUsecase(ProductRepo)
	ProductDelivery := product_delivery.CreateProductDelivery(ProductUsecase)

	router := gin.Default()
	router.Use(cors.Default())

	router.POST("/users", UserDelivery.CreateUser)
	router.GET("/users", UserDelivery.GetUserList)
	router.GET("/users/:id", UserDelivery.GetUserDetail)
	router.PUT("/users/:id", UserDelivery.UpdateUser)
	router.DELETE("/users/:id", UserDelivery.DeleteUser)

	router.GET("/roles", UserDelivery.GetRoleList)

	router.POST("/products", ProductDelivery.CreateProduct)

	return router
}
