package graph

import (
	"context"
	"errors"

	gopherql "github.com/shohinsan/GopherQL"
)

func mapAuthResponse(a gopherql.AuthResponse) *AuthResponse {
	return &AuthResponse{
		AccessToken: a.AccessToken,
		User:        mapUser(a.User),
	}
}

func (m *mutationResolver) Register(ctx context.Context, input RegisterInput) (*AuthResponse, error) {
	res, err := m.AuthService.Register(ctx, gopherql.RegisterInput{
		Email:           input.Email,
		Username:        input.Username,
		Password:        input.Password,
		ConfirmPassword: input.ConfirmPassword,
	})
	if err != nil {
		switch {
		case errors.Is(err, gopherql.ErrValidation) ||
			errors.Is(err, gopherql.ErrEmailTaken) ||
			errors.Is(err, gopherql.ErrUsernameTaken):
			return nil, buildBadRequestError(ctx, err)
		default:
			return nil, err
		}
	}

	return mapAuthResponse(res), nil
}

func (m *mutationResolver) Login(ctx context.Context, input LoginInput) (*AuthResponse, error) {
	res, err := m.AuthService.Login(ctx, gopherql.LoginInput{
		Email:    input.Email,
		Password: input.Password,
	})
	if err != nil {
		switch {
		case errors.Is(err, gopherql.ErrValidation) ||
			errors.Is(err, gopherql.ErrBadCredentials):
			return nil, buildBadRequestError(ctx, err)
		default:
			return nil, err
		}
	}

	return mapAuthResponse(res), nil
}
