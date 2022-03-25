package user_delivery

import (
	"encoding/json"
	"latihanFSE/models/dto"

	"github.com/gin-gonic/gin"
)

func (u *UserDelivery) CreateUser(c *gin.Context) {
	d := json.NewDecoder(c.Request.Body)
	d.DisallowUnknownFields()

	CreateUserRequest := dto.CreateUserRequest{}
	err := d.Decode(&CreateUserRequest)
	if err != nil {
		dto.JsonRequestErrorResponse(c, err)
		return
	}

	response := u.UserUsecase.CreateUser(CreateUserRequest)
	c.JSON(response.StatusCode, response)
}

func (u *UserDelivery) GetUserList(c *gin.Context) {
	response := u.UserUsecase.GetUserList()
	c.JSON(response.StatusCode, response)
}

func (u *UserDelivery) GetUserDetail(c *gin.Context) {
	ID := c.Param("id")
	response := u.UserUsecase.GetUserDetail(ID)
	c.JSON(response.StatusCode, response)
}
