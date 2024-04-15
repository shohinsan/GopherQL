package test_helpers

import (
	"context"
	"testing"

	gopherql "github.com/shohinsan/GopherQL"
	"github.com/shohinsan/GopherQL/faker"
	"github.com/shohinsan/GopherQL/postgres"
	"github.com/stretchr/testify/require"
)

func TeardownDB(ctx context.Context, t *testing.T, db *postgres.DB) {
	t.Helper()

	err := db.Truncate(ctx)
	require.NoError(t, err)
}

func CreateUser(ctx context.Context, t *testing.T, userRepo gopherql.UserRepo) gopherql.User {
	t.Helper()

	user, err := userRepo.Create(ctx, gopherql.User{
		Username: faker.Username(),
		Email:    faker.Email(),
		Password: faker.Password,
	})
	require.NoError(t, err)

	return user
}

func CreateTweet(ctx context.Context, t *testing.T, tweetRepo gopherql.TweetRepo, forUser string) gopherql.Tweet {
	t.Helper()

	tweet, err := tweetRepo.Create(ctx, gopherql.Tweet{
		Body:   faker.RandStr(20),
		UserID: forUser,
	})
	require.NoError(t, err)

	return tweet
}

func LoginUser(ctx context.Context, t *testing.T, user gopherql.User) context.Context {
	t.Helper()

	return gopherql.PutUserIDIntoContext(ctx, user.ID)
}
