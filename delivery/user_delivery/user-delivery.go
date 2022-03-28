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

func (u *UserDelivery) GetRoleList(c *gin.Context) {
	response := u.UserUsecase.GetRoleList()
	c.JSON(response.StatusCode, response)
}

func (u *UserDelivery) GetUserDetail(c *gin.Context) {
	ID := c.Param("id")
	response := u.UserUsecase.GetUserDetail(ID)
	c.JSON(response.StatusCode, response)
}

func (u *UserDelivery) DeleteUser(c *gin.Context) {
	ID := c.Param("id")
	response := u.UserUsecase.DeleteUser(ID)
	c.JSON(response.StatusCode, response)
}

func (u *UserDelivery) UpdateUser(c *gin.Context) {
	ID := c.Param("id")
	d := json.NewDecoder(c.Request.Body)
	d.DisallowUnknownFields()

	UpdateUserRequest := dto.UpdateUserRequest{}
	err := d.Decode(&UpdateUserRequest)
	if err != nil {
		dto.JsonRequestErrorResponse(c, err)
		return
	}

	response := u.UserUsecase.UpdateUser(ID, UpdateUserRequest)
	c.JSON(response.StatusCode, response)
}

func (u *UserDelivery) LoginUser(c *gin.Context) {
	d := json.NewDecoder(c.Request.Body)
	d.DisallowUnknownFields()

	LoginRequest := dto.LoginRequest{}
	err := d.Decode(&LoginRequest)
	if err != nil {
		dto.JsonRequestErrorResponse(c, err)
		return
	}

	response := u.UserUsecase.LoginUser(LoginRequest)
	c.JSON(response.StatusCode, response)
}
