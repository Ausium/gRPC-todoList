package handler

import (
	"context"
	"user/internal/repository"
	"user/internal/service"
	pb "user/internal/service"
	"user/pkg/e"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

// 用户注册
func (*UserService) UserRegister(ctx context.Context, req *service.UserRequest) (resp *service.UserDetailResponse, err error) {
	var user repository.User
	resp = new(service.UserDetailResponse)
	resp.Code = e.SUCCESS
	err = user.UserCreate(req)
	if err != nil {
		resp.Code = e.ERROR
		return resp, nil
	}
	resp.UserDetail = repository.BuildUser(user)
	return resp, nil
}

// 用户登录 token
func (*UserService) UserLogin(ctx context.Context, req *service.UserRequest) (resp *service.UserDetailResponse, err error) {
	var user repository.User
	resp = new(service.UserDetailResponse)
	resp.Code = e.SUCCESS
	err = user.ShowUserInfo(req)
	if err != nil {
		resp.Code = e.ERROR
		return resp, nil
	}
	resp.UserDetail = repository.BuildUser(user)
	return resp, nil
}
