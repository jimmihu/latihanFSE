package app

import (
	"latihanFSE/delivery/product_delivery"
	"latihanFSE/delivery/user_delivery"
	"latihanFSE/middleware"
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
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
	}))

	router.POST("/users", UserDelivery.CreateUser)
	router.POST("/login", UserDelivery.LoginUser)

	routeGroupAdmin := router.Group("/")
	routeGroupAdmin.Use(middleware.JwtAuthRoles([]string{"admin"}, JwtUsecase))
	{
		routeGroupAdmin.GET("/users", UserDelivery.GetUserList)
		routeGroupAdmin.GET("/users/:id", UserDelivery.GetUserDetail)
		routeGroupAdmin.PUT("/users/:id", UserDelivery.UpdateUser)
		routeGroupAdmin.DELETE("/users/:id", UserDelivery.DeleteUser)
	}
	router.GET("/roles", UserDelivery.GetRoleList)

	routeGroupAuth := router.Group("/")
	routeGroupAuth.Use(middleware.JwtAuth(JwtUsecase))
	{
		routeGroupAuth.POST("/products", ProductDelivery.CreateProduct)
		routeGroupAuth.GET("/products", ProductDelivery.GetProductList)
		routeGroupAuth.GET("/products/:id", ProductDelivery.GetProductDetail)
		routeGroupAuth.DELETE("/products/:id", ProductDelivery.DeleteProduct)
		routeGroupAuth.PUT("/products/:id", ProductDelivery.UpdateProduct)
	}
	routeGroupChecker := router.Group("/")
	routeGroupChecker.Use(middleware.JwtAuthRoles([]string{"admin", "checker"}, JwtUsecase))
	{
		routeGroupChecker.PUT("/products/:id/checked", ProductDelivery.CheckProduct)
	}
	routeGroupSigner := router.Group("/")
	routeGroupSigner.Use(middleware.JwtAuthRoles([]string{"admin", "signer"}, JwtUsecase))
	{
		routeGroupSigner.PUT("/products/:id/published", ProductDelivery.PublishProduct)
	}

	return router
}
