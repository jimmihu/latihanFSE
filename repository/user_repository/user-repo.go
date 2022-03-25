package user_repository

import (
	"latihanFSE/models/entity"
	"latihanFSE/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (repo *UserRepo) CreateUser(user *entity.User) *gorm.DB {
	user.Password, _ = utils.HashPassword(user.Password)

	role := GetViewerRole(repo)
	user.RoleID = role.ID
	result := repo.mysqlConn.Create(&user)
	return result
}

func GetViewerRole(repo *UserRepo) entity.Role {

	role := entity.Role{}
	repo.mysqlConn.Where("title = ?", "viewer").First(&role)
	return role
}

func (repo *UserRepo) GetUserList() ([]entity.UserList, *gorm.DB) {
	UserListResult := []entity.UserList{}
	Result := repo.mysqlConn.Model(&entity.User{}).Joins("Role").Find(&UserListResult)
	return UserListResult, Result
}

func (repo *UserRepo) GetUserDetail(Id uuid.UUID) (entity.User, *gorm.DB) {
	UserDetailResult := entity.User{}
	Result := repo.mysqlConn.Model(&entity.User{}).Joins("Role").First(&UserDetailResult, Id)
	return UserDetailResult, Result
}
