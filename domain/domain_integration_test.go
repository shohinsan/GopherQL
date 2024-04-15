//go:build integration
// +build integration

package domain

import (
	"context"
	"log"
	"os"
	"testing"

	gopherql "github.com/shohinsan/GopherQL"
	"github.com/shohinsan/GopherQL/config"
	"github.com/shohinsan/GopherQL/jwt"
	"github.com/shohinsan/GopherQL/postgres"
	"golang.org/x/crypto/bcrypt"
)

var (
	conf             *config.Config
	db               *postgres.DB
	authTokenService gopherql.AuthTokenService
	authService      gopherql.AuthService
	tweetService     gopherql.TweetService
	userRepo         gopherql.UserRepo
	tweetRepo        gopherql.TweetRepo
)

func TestMain(m *testing.M) {
	ctx := context.Background()

	config.LoadEnv(".env.test")

	passwordCost = bcrypt.MinCost

	conf = config.New()

	db = postgres.New(ctx, conf)
	defer db.Close()

	if err := db.Drop(); err != nil {
		log.Fatal(err)
	}

	if err := db.Migrate(); err != nil {
		log.Fatal(err)
	}

	userRepo = postgres.NewUserRepo(db)
	tweetRepo = postgres.NewTweetRepo(db)

	authTokenService = jwt.NewTokenService(conf)

	authService = NewAuthService(userRepo, authTokenService)
	tweetService = NewTweetService(tweetRepo)

	os.Exit(m.Run())
}
