package graph

import (
	"context"

	gopherql "github.com/shohinsan/GopherQL"
)

func mapUser(u gopherql.User) *User {
	return &User{
		ID:        u.ID,
		Email:     u.Email,
		Username:  u.Username,
		CreatedAt: u.CreatedAt,
	}
}

func (q *queryResolver) Me(ctx context.Context) (*User, error) {
	userID, err := gopherql.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, gopherql.ErrUnauthenticated
	}

	return mapUser(gopherql.User{
		ID: userID,
	}), nil
}
