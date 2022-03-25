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

func (repo *UserRepo) GetRoleList() ([]entity.Role, *gorm.DB) {
	RoleListResult := []entity.Role{}
	result := repo.mysqlConn.Model(&entity.Role{}).Find(&RoleListResult)
	return RoleListResult, result
}

func (repo *UserRepo) GetUserDetail(ID uuid.UUID) (entity.UserDetail, *gorm.DB) {
	UserDetailResult := entity.UserDetail{}
	Result := repo.mysqlConn.Model(&entity.User{}).Joins("Role").First(&UserDetailResult, ID)
	return UserDetailResult, Result
}

func (repo *UserRepo) DeleteUser(ID uuid.UUID) *gorm.DB {
	Result := repo.mysqlConn.Delete(&entity.User{}, ID)
	return Result
}

func (repo *UserRepo) UpdateUser(ID uuid.UUID, user *entity.User) *gorm.DB {
	if user.Password == "" {

	} else {
		user.Password, _ = utils.HashPassword(user.Password)
	}
	Result := repo.mysqlConn.Model(entity.User{ID: ID}).Updates(&user)
	return Result
}
