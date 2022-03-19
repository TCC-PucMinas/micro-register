package controller

import (
	"context"
	"errors"

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
	user.Phone = request.Phone
	user.Business = request.Business
	user.CpfCnpj = request.CpfCnpj

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
	user.Active = 0

	idUser, err := user.CreateUser()
	if err != nil {
		return res, err
	}

	for _, v := range request.Address {
		address := model.Address{
			Street:       v.Street,
			State:        v.State,
			Number:       v.Number,
			Neighborhood: v.Neighborhood,
			Country:      v.Country,
			Complement:   v.Complement,
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
		From:       user.Email,
		CodeActive: user.CodeActive,
		Subject:    "Ative sua conta",
	}

	_ = nats.PublishAlertEmail(senderNatsActiveCode, emailSender)

	defer nats.Nats.Close()

	res.Created = true
	return res, nil
}

func (s *UserCommunicate) ValidateUserExist(ctx context.Context, request *communicate.ValidateUserExistRequest) (*communicate.ValidateUserExistResponse, error) {
	res := &communicate.ValidateUserExistResponse{}

	user := model.User{
		Email:   request.Email,
		CpfCnpj: request.CpfCnpj,
	}

	if err := user.GetOneUserByEmailAndCpf_Cpnj(); err != nil || user.Id != 0 {
		return res, errors.New("Email or cpf already registered!")
	}

	res.Valid = true
	return res, nil
}

func (s *UserCommunicate) ValidateCodeActive(ctx context.Context, request *communicate.ValidateCodeActiveRequest) (*communicate.ValidateCodeActiveResponse, error) {
	res := &communicate.ValidateCodeActiveResponse{}

	user := model.User{
		CodeActive: request.CodeActive,
		Active:     0,
	}

	if err := user.GetOneUserByCodeActive(); err != nil {
		return res, errors.New("Code invalid!")
	}

	res.Valid = true
	return res, nil
}

func (s *UserCommunicate) ActivatedUserCodeActive(ctxt context.Context, request *communicate.ActiveCodeUserRequest) (*communicate.ActiveCodeUserResponse, error) {
	res := &communicate.ActiveCodeUserResponse{}

	user := model.User{
		CodeActive: request.CodeActive,
		Active:     0,
	}

	if err := user.GetOneUserByCodeActive(); err != nil || user.Id == 0 {
		return res, errors.New("Code invalid!")
	}

	user.Active = 1

	if err := user.UpdateUserSetActiveUser(); err != nil {
		return res, err
	}

	res.Actived = true

	return res, nil
}

func (s *UserCommunicate) ListAllUserPaginateByName(ctx context.Context, request *communicate.ListAllUserPaginateByNameRequest) (*communicate.ListAllUserPaginateByNameResponse, error) {
	res := &communicate.ListAllUserPaginateByNameResponse{}

	user := model.User{}

	users, total, err := user.GetByNameLike(request.Name, request.Page, request.Limit)

	if err != nil {
		return res, err
	}

	data := &communicate.Data{}
	for _, c := range users {
		user := &communicate.User{}
		user.Id = c.Id
		user.FirstName = c.FirstName
		user.LastName = c.LastName
		user.Email = c.Email
		user.Phone = c.Phone
		user.Business = c.Business
		user.CpfCnpj = c.CpfCnpj
		user.Permission = c.Permission.Name
		data.Users = append(data.Users, user)
	}

	res.Total = total
	res.Page = request.Page
	res.Limit = request.Limit

	res.Data = data

	return res, nil
}
