package communicate

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/TCC-PucMinas/micro-register/helpers"
	"github.com/TCC-PucMinas/micro-register/model"
)

type Server struct {
	AuthenticateCommunicateServer
}

func (s *Server) Authenticate(ctx context.Context, request *AuthenticateValidateRequest) (*AuthenticateValidateResponse, error) {

	res := &AuthenticateValidateResponse{}
	userResponse := &UserResponse{}

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
		roles := Roles{}
		roles.Id = strconv.Itoa(int(v.Id))
		roles.Name = v.Name

		userResponse.Roles = append(userResponse.Roles, &roles)

		menus, err := model.GetAllMenus(v.Id)

		if err != nil {
			return nil, errors.New("Error in get menus!")
		}

		for _, v := range menus {
			menus := Menus{}
			menus.Id = strconv.Itoa(int(v.Id))
			menus.Name = v.Name
			menus.IconMenu = v.IconMenu
			menus.UrlMenu = v.UrlMenu
			userResponse.Menus = append(userResponse.Menus, &menus)
		}
	}

	res.UserResponse = userResponse
	return res, nil
}

func (s *Server) ValidateToken(ctx context.Context, request *TokenRequest) (*TokenResponse, error) {
	res := &TokenResponse{}

	if check, err := helpers.CheckJwt(request.AccessToken, 0); err != nil || !check {
		return res, errors.New("Token invalid!")
	}
	res.Validate = true
	return res, nil
}

func (s *Server) RefreshToken(ctx context.Context, request *RefreshRequest) (*RefreshResponse, error) {

	res := &RefreshResponse{}

	if validate, err := helpers.CheckJwt(request.RefreshToken, 1); !validate || (err != nil) {
		return res, errors.New("Refresh token invalid!")
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
		return res, nil
	}
}
