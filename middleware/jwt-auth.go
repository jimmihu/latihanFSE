package middleware

import (
	"latihanFSE/models/dto"
	"latihanFSE/usecase/jwt_usecase"

	"github.com/gin-gonic/gin"
)

func JwtAuth(jwtUsecase jwt_usecase.JwtUsecaseInterface) gin.HandlerFunc {

	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")
		payload, err := jwtUsecase.ValidateTokenAndGetPayload(authHeader)
		if err != nil {
			res := dto.UnauthorizedResponse(err)
			c.AbortWithStatusJSON(res.StatusCode, res)
			return
		}

		c.Set("user_id", payload.UserID)
	}
}

func JwtAuthRoles(roles []string, JwtUsecase jwt_usecase.JwtUsecaseInterface) gin.HandlerFunc {

	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")
		payload, err := JwtUsecase.ValidateTokenAndGetPayload(authHeader)
		if err != nil {
			res := dto.UnauthorizedResponse(err)
			c.AbortWithStatusJSON(res.StatusCode, res)
			return
		}
		authorized, err2 := JwtUsecase.UserHasAuthorization(payload.UserID, roles)

		if err2 != nil {
			res := dto.DBErrorResponse(err)
			c.AbortWithStatusJSON(res.StatusCode, res)
			return
		}

		if !authorized {
			res := dto.UnauthorizedResponse(err)
			c.AbortWithStatusJSON(res.StatusCode, res)
			return
		}

		c.Set("user_id", payload.UserID)
	}
}
