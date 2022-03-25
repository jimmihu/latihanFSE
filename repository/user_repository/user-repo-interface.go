package user_repository

import (
	"latihanFSE/models/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepoInterface interface {
	CreateUser(user *entity.User) *gorm.DB
	GetUserList() ([]entity.UserList, *gorm.DB)
	GetRoleList() ([]entity.Role, *gorm.DB)
	GetUserDetail(ID uuid.UUID) (entity.UserDetail, *gorm.DB)
	DeleteUser(ID uuid.UUID) *gorm.DB
}

type UserRepo struct {
	mysqlConn *gorm.DB
}

func GetUserRepo(mysqlConn *gorm.DB) UserRepoInterface {
	return &UserRepo{
		mysqlConn: mysqlConn,
	}
}
