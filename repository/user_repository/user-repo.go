package user_repository

import (
	"latihanFSE/models/entity"
	"latihanFSE/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (repo *UserRepo) CreateUser(user *entity.User) *gorm.DB {
	user.Password, _ = utils.HashPassword(user.Password)

	role := entity.Role{}
	repo.mysqlConn.Where("title = ?", "viewer").First(&role)
	user.RoleID = role.ID
	result := repo.mysqlConn.Create(&user)
	return result
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

func (repo *UserRepo) GetLoginUser(pn string) (entity.User, *gorm.DB) {
	UserResult := entity.User{}
	Result := repo.mysqlConn.Model(&entity.User{}).Joins("Role").Where("personal_number = ?", pn).First(&UserResult)
	return UserResult, Result
}
