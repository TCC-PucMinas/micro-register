package controller

import (
	"context"

	"github.com/TCC-PucMinas/micro-register/communicate"
	"github.com/TCC-PucMinas/micro-register/helpers"
	"github.com/TCC-PucMinas/micro-register/model"
	"github.com/TCC-PucMinas/micro-register/service"
)

var senderNatsActiveCode = "email.user.active"

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

	idUser, err := user.CreateUser()
	if err != nil {
		return res, err
	}

	for _, v := range request.Address {
		address := model.Address{
			Street:     v.Street,
			State:      v.State,
			Number:     v.Number,
			Country:    v.Country,
			Complement: v.Complement,
		}
		idAddress, err := address.CreateAddress()
		if err != nil {
			return res, err
		}

		addressUser := model.AddressUser{
			UserId:    idUser,
			AddressId: idAddress,
		}
		if _, err := addressUser.CreateAddressUser(); err != nil {
			return res, err
		}
	}

	nats := service.Nats{}

	if err := nats.Connect(); err != nil {
		return res, err
	}

	emailSender := &service.EmailCommunicate{
		From:    user.Email,
		Forgot:  user.CodeActive,
		Subject: "Ative sua conta",
	}

	nats.PublishAlertEmail(senderNatsEmail, emailSender)

	defer nats.Nats.Close()

	res.Created = true
	return res, nil
}
