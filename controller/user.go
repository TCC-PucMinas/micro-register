package controller

import (
	"context"

	"github.com/TCC-PucMinas/micro-register/communicate"
)

// var senderNatsEmail = "email.user.forgot"

type UserCommunicate struct {
	communicate.UserCommunicateServer
}

func (s *UserCommunicate) CreateUser(ctx context.Context, request *communicate.CreateUserRequest) (*communicate.CreateUserResponse, error) {
	res := &communicate.CreateUserResponse{}

	return res, nil
}
