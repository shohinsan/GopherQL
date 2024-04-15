package domain

import (
	"context"

	gopherql "github.com/shohinsan/GopherQL"
	"github.com/shohinsan/GopherQL/uuid"
)

type UserService struct {
	UserRepo gopherql.UserRepo
}

func NewUserService(ur gopherql.UserRepo) *UserService {
	return &UserService{
		UserRepo: ur,
	}
}

func (u *UserService) GetByID(ctx context.Context, id string) (gopherql.User, error) {
	if !uuid.Validate(id) {
		return gopherql.User{}, gopherql.ErrInvalidUUID
	}

	return u.UserRepo.GetByID(ctx, id)
}
