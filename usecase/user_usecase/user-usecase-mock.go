package user_usecase

import (
	"latihanFSE/models/dto"

	"github.com/stretchr/testify/mock"
)

type UserUsecaseMock struct {
	mock.Mock
}

func (usecaseMock *UserUsecaseMock) Ping() string {
	return "Pong"
}

func (useCaseMock *UserUsecaseMock) CreateUser(request dto.CreateUserRequest) dto.HttpResponse {
	args := useCaseMock.Called(request)
	if args.Get(0) == nil {
		return dto.DBErrorResponse(args.Error(0))
	}
	return args.Get(0).(dto.HttpResponse)
}

func (useCaseMock *UserUsecaseMock) DeleteUser(s string) dto.HttpResponse {
	args := useCaseMock.Called(s)
	if args.Get(0) == nil {
		return dto.DBErrorResponse(args.Error(0))
	}
	return args.Get(0).(dto.HttpResponse)
}

func (useCaseMock *UserUsecaseMock) GetUserList() dto.HttpResponse {
	args := useCaseMock.Called()
	if args.Get(0) == nil {
		return dto.DBErrorResponse(args.Error(0))
	}
	return args.Get(0).(dto.HttpResponse)
}

func (useCaseMock *UserUsecaseMock) GetRoleList() dto.HttpResponse {
	args := useCaseMock.Called()
	if args.Get(0) == nil {
		return dto.DBErrorResponse(args.Error(0))
	}
	return args.Get(0).(dto.HttpResponse)
}

func (useCaseMock *UserUsecaseMock) GetUserDetail(s string) dto.HttpResponse {
	args := useCaseMock.Called(s)
	if args.Get(0) == nil {
		return dto.DBErrorResponse(args.Error(0))
	}
	return args.Get(0).(dto.HttpResponse)
}

func (useCaseMock *UserUsecaseMock) UpdateUser(s string, request dto.UpdateUserRequest) dto.HttpResponse {
	args := useCaseMock.Called(s, request)
	if args.Get(0) == nil {
		return dto.DBErrorResponse(args.Error(0))
	}
	return args.Get(0).(dto.HttpResponse)
}

func (useCaseMock *UserUsecaseMock) LoginUser(request dto.LoginRequest) dto.HttpResponse {
	args := useCaseMock.Called(request)
	if args.Get(0) == nil {
		return dto.DBErrorResponse(args.Error(0))
	}
	return args.Get(0).(dto.HttpResponse)
}
