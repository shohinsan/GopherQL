package domain

import (
	"context"
	"errors"
	"fmt"

	gopherql "github.com/shohinsan/GopherQL"
	"golang.org/x/crypto/bcrypt"
)

var passwordCost = bcrypt.DefaultCost

type AuthService struct {
	AuthTokenService gopherql.AuthTokenService
	UserRepo         gopherql.UserRepo
}

func NewAuthService(ur gopherql.UserRepo, ats gopherql.AuthTokenService) *AuthService {
	return &AuthService{
		AuthTokenService: ats,
		UserRepo:         ur,
	}
}

func (as *AuthService) Register(ctx context.Context, input gopherql.RegisterInput) (gopherql.AuthResponse, error) {
	input.Sanitize()

	if err := input.Validate(); err != nil {
		return gopherql.AuthResponse{}, err
	}

	// check if username is already taken
	if _, err := as.UserRepo.GetByUsername(ctx, input.Username); !errors.Is(err, gopherql.ErrNotFound) {
		return gopherql.AuthResponse{}, gopherql.ErrUsernameTaken
	}

	// check if email is already taken
	if _, err := as.UserRepo.GetByEmail(ctx, input.Email); !errors.Is(err, gopherql.ErrNotFound) {
		return gopherql.AuthResponse{}, gopherql.ErrEmailTaken
	}

	user := gopherql.User{
		Email:    input.Email,
		Username: input.Username,
	}

	// hash the password
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), passwordCost)
	if err != nil {
		return gopherql.AuthResponse{}, fmt.Errorf("error hashing password: %v", err)
	}

	user.Password = string(hashPassword)

	// create the user
	user, err = as.UserRepo.Create(ctx, user)
	if err != nil {
		return gopherql.AuthResponse{}, fmt.Errorf("error creating user: %v", err)
	}

	accessToken, err := as.AuthTokenService.CreateAccessToken(ctx, user)
	if err != nil {
		return gopherql.AuthResponse{}, gopherql.ErrGenAccessToken
	}

	// return accessToken and user
	return gopherql.AuthResponse{
		AccessToken: accessToken,
		User:        user,
	}, nil
}

func (as *AuthService) Login(ctx context.Context, input gopherql.LoginInput) (gopherql.AuthResponse, error) {
	input.Sanitize()

	if err := input.Validate(); err != nil {
		return gopherql.AuthResponse{}, err
	}

	user, err := as.UserRepo.GetByEmail(ctx, input.Email)
	if err != nil {
		switch {
		case errors.Is(err, gopherql.ErrNotFound):
			return gopherql.AuthResponse{}, gopherql.ErrBadCredentials
		default:
			return gopherql.AuthResponse{}, err
		}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return gopherql.AuthResponse{}, gopherql.ErrBadCredentials
	}

	accessToken, err := as.AuthTokenService.CreateAccessToken(ctx, user)
	if err != nil {
		return gopherql.AuthResponse{}, gopherql.ErrGenAccessToken
	}

	return gopherql.AuthResponse{
		AccessToken: accessToken,
		User:        user,
	}, nil
}
