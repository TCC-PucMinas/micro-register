package communicate

import (
	"context"
)

type Server struct {
	AuthenticateCommunicateServer
}

func (s *Server) ValidateCode(ctx context.Context, request *AuthenticateValidateRequest) (*AuthenticateValidateResponse, error) {

	res := &AuthenticateValidateResponse{}

	res.AccessToken = ""
	res.RefreshToken = ""

	return res, nil
}
