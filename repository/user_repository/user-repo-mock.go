package user_repository

import (
	"latihanFSE/models/entity"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (repoMock *UserRepositoryMock) Ping() string {
	return "Pong"
}

func (repoMock *UserRepositoryMock) CreateUser(request *entity.User) *gorm.DB {
	args := repoMock.Called(request)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*gorm.DB)
}

func (repoMock *UserRepositoryMock) DeleteUser(ID uuid.UUID) *gorm.DB {
	return nil
}

func (repoMock *UserRepositoryMock) GetUserList() ([]entity.UserList, *gorm.DB) {
	return nil, nil
}

func (repoMock *UserRepositoryMock) GetRoleList() ([]entity.Role, *gorm.DB) {
	return nil, nil
}

func (repoMock *UserRepositoryMock) GetUserDetail(ID uuid.UUID) (entity.UserDetail, *gorm.DB) {
	return entity.UserDetail{}, nil
}

func (repoMock *UserRepositoryMock) UpdateUser(ID uuid.UUID, request *entity.User) *gorm.DB {
	return nil
}

func (repoMock *UserRepositoryMock) GetLoginUser(pn string) (entity.User, *gorm.DB) {
	return entity.User{}, nil
}
