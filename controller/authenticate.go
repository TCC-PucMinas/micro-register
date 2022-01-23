package controller

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/TCC-PucMinas/micro-register/communicate"
	"github.com/TCC-PucMinas/micro-register/helpers"
	"github.com/TCC-PucMinas/micro-register/model"
)

type AuthenticateServer struct {
	communicate.AuthenticateCommunicateServer
}

func (s *AuthenticateServer) Authenticate(ctx context.Context, request *communicate.AuthenticateValidateRequest) (*communicate.AuthenticateValidateResponse, error) {

	res := &communicate.AuthenticateValidateResponse{}
	userResponse := &communicate.UserResponse{}

	auth := model.Authenticate{
		Email:    request.Email,
		Password: request.Password,
	}

	user, err := auth.GetOneUserByEmail()

	if err != nil {
		return res, err
	}

	if !helpers.CompareHashs(user.Password, auth.Password) {
		return res, errors.New("Login e senha incorreto!")
	}

	id := strconv.Itoa(int(user.Id))

	jwt, err := helpers.GenerateJwt(id)

	if err != nil {
		return res, err
	}

	permissions, err := model.GetAllRoles(user.Id)

	if err != nil {
		return res, err
	}

	res.AccessToken = jwt.AccessToken
	res.RefreshToken = jwt.RefreshToken

	userResponse.Id = id
	userResponse.Email = user.Email
	userResponse.Name = fmt.Sprintf("%v %v", user.FirstName, user.LastName)

	for _, v := range permissions {
		roles := communicate.Roles{}
		roles.Id = strconv.Itoa(int(v.Id))
		roles.Name = v.Name

		userResponse.Roles = append(userResponse.Roles, &roles)

		if err != nil {
			return nil, errors.New("Error in get menus!")
		}
	}

	if err := jwt.SetTokenCache(); err != nil {
		return res, nil
	}

	res.UserResponse = userResponse
	return res, nil
}

func (s *AuthenticateServer) ValidateToken(ctx context.Context, request *communicate.TokenRequest) (*communicate.TokenResponse, error) {
	res := &communicate.TokenResponse{}

	if check, err := helpers.CheckJwt(request.AccessToken, 0); err != nil || !check {
		return res, errors.New("Token invalid!")
	}

	jwt := helpers.JwtTokens{}

	if text, err := jwt.GetToken(request.AccessToken); err != nil || len(text) < 2 {
		return res, errors.New("Token invalid!")
	}

	route := model.Route{}
	route.Path = request.Path
	if claim, err := helpers.ExtractJwt(request.AccessToken); err != nil {
		return res, errors.New("Token invalid!")
	} else {
		if err := route.GetOneRouteByPathAndUserId(claim.Username); err != nil {
			return res, errors.New("Not permission!")
		}
	}

	res.Validate = true
	return res, nil
}

func (s *AuthenticateServer) RefreshToken(ctx context.Context, request *communicate.RefreshRequest) (*communicate.RefreshResponse, error) {

	res := &communicate.RefreshResponse{}

	if validate, err := helpers.CheckJwt(request.RefreshToken, 1); !validate || (err != nil) {
		return res, errors.New("Refresh token invalid!")
	}

	if validate, err := helpers.CheckJwt(request.AccessToken, 0); !validate || (err != nil) {
		return res, errors.New("Accesss token invalid!")
	}

	claim, err := helpers.ExtractJwt(request.RefreshToken)

	if err != nil {
		return res, err
	}

	if jwt, err := helpers.GenerateJwt(claim.Id); err != nil {
		return res, err
	} else {
		res.AccessToken = jwt.AccessToken
		res.RefreshToken = jwt.RefreshToken

		if err := jwt.SetTokenCache(); err != nil {
			return res, err
		}
	}
	return res, nil
}

func (s *AuthenticateServer) ValidateEmail(ctx context.Context, request *communicate.EmailRequest) (*communicate.EmailResponse, error) {
	res := &communicate.EmailResponse{}

	auth := model.Authenticate{
		Email: request.Email,
	}

	if _, err := auth.GetOneUserByEmail(); err != nil {
		res.Valid = false
		return res, nil
	}

	res.Valid = true

	return res, nil
}
