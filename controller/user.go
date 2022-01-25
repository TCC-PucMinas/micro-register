package controller

import (
	"context"

	"github.com/TCC-PucMinas/micro-register/communicate"
	"github.com/TCC-PucMinas/micro-register/helpers"
	"github.com/TCC-PucMinas/micro-register/model"
)

// var senderNatsEmail = "email.user.forgot"

type UserCommunicate struct {
	communicate.UserCommunicateServer
}

func (s *UserCommunicate) CreateUser(ctx context.Context, request *communicate.CreateUserRequest) (*communicate.CreateUserResponse, error) {
	res := &communicate.CreateUserResponse{}

	user := model.User{}

	user.FirstName = request.FirstName
	user.LastName = request.LastName
	user.Email = request.Email
	user.Business = request.Busines
	user.CpfCnpj = request.CpfCpnj

	password, err := helpers.GenerateHashSalt(request.Password)

	if err != nil {
		return res, err
	}

	user.Password = password

	codeActive, err := helpers.GenerateHashSalt(user.Business + user.LastName)

	if err != nil {
		return res, err
	}
	user.CodeActive = codeActive
	user.Active = 1

	return res, nil
}
