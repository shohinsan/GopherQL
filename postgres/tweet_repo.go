package postgres

import (
	"context"
	"fmt"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	gopherql "github.com/shohinsan/GopherQL"
)

type TweetRepo struct {
	DB *DB
}

func NewTweetRepo(db *DB) *TweetRepo {
	return &TweetRepo{
		DB: db,
	}
}

func (tr *TweetRepo) All(ctx context.Context) ([]gopherql.Tweet, error) {
	return getAllTweets(ctx, tr.DB.Pool)
}

func getAllTweets(ctx context.Context, q pgxscan.Querier) ([]gopherql.Tweet, error) {
	query := `SELECT * FROM tweets WHERE parent_id IS NULL ORDER BY created_at DESC;`

	var tweets []gopherql.Tweet

	if err := pgxscan.Select(ctx, q, &tweets, query); err != nil {
		return nil, fmt.Errorf("error get all tweets %+v", err)
	}

	return tweets, nil
}

func (tr *TweetRepo) Create(ctx context.Context, tweet gopherql.Tweet) (gopherql.Tweet, error) {
	tx, err := tr.DB.Pool.Begin(ctx)
	if err != nil {
		return gopherql.Tweet{}, fmt.Errorf("error starting transaction: %v", err)
	}
	defer tx.Rollback(ctx)

	tweet, err = createTweet(ctx, tx, tweet)
	if err != nil {
		return gopherql.Tweet{}, err
	}

	if err := tx.Commit(ctx); err != nil {
		return gopherql.Tweet{}, fmt.Errorf("error commiting: %v", err)
	}

	return tweet, nil
}

func createTweet(ctx context.Context, tx pgx.Tx, tweet gopherql.Tweet) (gopherql.Tweet, error) {
	query := `INSERT INTO tweets (body, user_id, parent_id) VALUES ($1, $2, $3) RETURNING *;`

	t := gopherql.Tweet{}

	if err := pgxscan.Get(ctx, tx, &t, query, tweet.Body, tweet.UserID, tweet.ParentID); err != nil {
		return gopherql.Tweet{}, fmt.Errorf("error insert: %v", err)
	}

	return t, nil
}

func (tr *TweetRepo) GetByID(ctx context.Context, id string) (gopherql.Tweet, error) {
	return getTweetByID(ctx, tr.DB.Pool, id)
}

func getTweetByID(ctx context.Context, q pgxscan.Querier, id string) (gopherql.Tweet, error) {
	query := `SELECT * FROM tweets WHERE id = $1 LIMIT 1;`

	t := gopherql.Tweet{}

	if err := pgxscan.Get(ctx, q, &t, query, id); err != nil {
		if pgxscan.NotFound(err) {
			return gopherql.Tweet{}, gopherql.ErrNotFound
		}

		return gopherql.Tweet{}, fmt.Errorf("error get tweet: %+v", err)
	}

	return t, nil
}

func (tr *TweetRepo) Delete(ctx context.Context, id string) error {
	tx, err := tr.DB.Pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("error starting transaction: %v", err)
	}
	defer tx.Rollback(ctx)

	if err := deleteTweet(ctx, tx, id); err != nil {
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("error commiting: %v", err)
	}

	return nil
}

func deleteTweet(ctx context.Context, tx pgx.Tx, id string) error {
	query := `DELETE FROM tweets WHERE id = $1;`

	if _, err := tx.Exec(ctx, query, id); err != nil {
		return fmt.Errorf("error insert: %v", err)
	}

	return nil
}

func (tr *TweetRepo) GetByParentID(ctx context.Context, id string) ([]gopherql.Tweet, error) {
	return getTweetsByParentID(ctx, tr.DB.Pool, id)
}

func getTweetsByParentID(ctx context.Context, q pgxscan.Querier, id string) ([]gopherql.Tweet, error) {
	query := `SELECT * FROM tweets WHERE parent_id = $1;`

	var tweets []gopherql.Tweet

	if err := pgxscan.Select(ctx, q, &tweets, query, id); err != nil {
		return nil, fmt.Errorf("error get all tweets by parent id %+v", err)
	}

	return tweets, nil
}
