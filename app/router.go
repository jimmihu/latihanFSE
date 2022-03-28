package app

import (
	"latihanFSE/delivery/product_delivery"
	"latihanFSE/delivery/user_delivery"
	"latihanFSE/repository/product_repository"
	"latihanFSE/repository/user_repository"
	"latihanFSE/usecase/jwt_usecase"
	"latihanFSE/usecase/product_usecase"
	"latihanFSE/usecase/user_usecase"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRouter(mysqlConn *gorm.DB) *gin.Engine {

	UserRepo := user_repository.GetUserRepo(mysqlConn)
	JwtUsecase := jwt_usecase.CreateJwtUseCase(UserRepo)
	UserUsecase := user_usecase.CreateUserUsecase(UserRepo, JwtUsecase)
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
	router.GET("/products", ProductDelivery.GetProductList)
	router.GET("/products/:id", ProductDelivery.GetProductDetail)
	router.DELETE("/products/:id", ProductDelivery.DeleteProduct)
	router.PUT("/products/:id", ProductDelivery.UpdateProduct)
	router.PUT("/products/:id/checked", ProductDelivery.CheckProduct)
	router.PUT("/products/:id/published", ProductDelivery.PublishProduct)

	router.POST("/login", UserDelivery.LoginUser)

	return router
}
