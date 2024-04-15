package domain

import (
	"context"

	gopherql "github.com/shohinsan/GopherQL"
	"github.com/shohinsan/GopherQL/uuid"
)

type TweetService struct {
	TweetRepo gopherql.TweetRepo
}

func NewTweetService(tr gopherql.TweetRepo) *TweetService {
	return &TweetService{
		TweetRepo: tr,
	}
}

func (ts *TweetService) All(ctx context.Context) ([]gopherql.Tweet, error) {
	return ts.TweetRepo.All(ctx)
}

func (ts *TweetService) GetByParentID(ctx context.Context, id string) ([]gopherql.Tweet, error) {
	return ts.TweetRepo.GetByParentID(ctx, id)
}

func (ts *TweetService) Create(ctx context.Context, input gopherql.CreateTweetInput) (gopherql.Tweet, error) {
	currentUserID, err := gopherql.GetUserIDFromContext(ctx)
	if err != nil {
		return gopherql.Tweet{}, gopherql.ErrUnauthenticated
	}

	input.Sanitize()

	if err := input.Validate(); err != nil {
		return gopherql.Tweet{}, err
	}

	tweet, err := ts.TweetRepo.Create(ctx, gopherql.Tweet{
		Body:   input.Body,
		UserID: currentUserID,
	})
	if err != nil {
		return gopherql.Tweet{}, err
	}

	return tweet, nil
}

func (ts *TweetService) GetByID(ctx context.Context, id string) (gopherql.Tweet, error) {
	if !uuid.Validate(id) {
		return gopherql.Tweet{}, gopherql.ErrInvalidUUID
	}

	return ts.TweetRepo.GetByID(ctx, id)
}

func (ts *TweetService) Delete(ctx context.Context, id string) error {
	currentUserID, err := gopherql.GetUserIDFromContext(ctx)
	if err != nil {
		return gopherql.ErrUnauthenticated
	}

	if !uuid.Validate(id) {
		return gopherql.ErrInvalidUUID
	}

	tweet, err := ts.TweetRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if !tweet.CanDelete(gopherql.User{ID: currentUserID}) {
		return gopherql.ErrForbidden
	}

	return ts.TweetRepo.Delete(ctx, id)
}

func (ts *TweetService) CreateReply(ctx context.Context, parentID string, input gopherql.CreateTweetInput) (gopherql.Tweet, error) {
	currentUserID, err := gopherql.GetUserIDFromContext(ctx)
	if err != nil {
		return gopherql.Tweet{}, gopherql.ErrUnauthenticated
	}

	input.Sanitize()

	if err := input.Validate(); err != nil {
		return gopherql.Tweet{}, err
	}

	if !uuid.Validate(parentID) {
		return gopherql.Tweet{}, gopherql.ErrInvalidUUID
	}

	if _, err := ts.TweetRepo.GetByID(ctx, parentID); err != nil {
		return gopherql.Tweet{}, gopherql.ErrNotFound
	}

	tweet, err := ts.TweetRepo.Create(ctx, gopherql.Tweet{
		Body:     input.Body,
		UserID:   currentUserID,
		ParentID: &parentID,
	})
	if err != nil {
		return gopherql.Tweet{}, err
	}

	return tweet, nil
}
