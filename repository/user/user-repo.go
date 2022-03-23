package repository

import (
	"latihanFSE/models/entity"
	"latihanFSE/utils"

	"github.com/jinzhu/gorm"
)

func (repo *UserRepo) CreateUser(user *entity.User) *gorm.DB {
	user.Password, _ = utils.HashPassword(user.Password)

	role := getViewerRole(repo)
	user.RoleID = role.ID
	result := repo.mysqlConn.Create(&user)
	return result
}

func getViewerRole(repo *UserRepo) entity.Role {

	role := entity.Role{}
	repo.mysqlConn.Where("title = ?", "viewer").First(&role)
	return role
}
