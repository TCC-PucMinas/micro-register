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
