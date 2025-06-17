package db

import (
	"context"
	"errors"
	"time"
)

var (
	ErrArticleExists = errors.New("article already exists")
)

type Article struct {
	ID          int64
	GUID        string
	Title       string
	Link        string
	Description string
	PubDate     time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (a *Article) Save(ctx context.Context) error {
	// First check if article exists
	exists, err := a.Exists(ctx)
	if err != nil {
		return err
	}
	if exists {
		return ErrArticleExists
	}

	query := `
		INSERT INTO articles (guid, title, link, description, pub_date)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at, updated_at`

	return DB.QueryRow(ctx, query,
		a.GUID, a.Title, a.Link, a.Description, a.PubDate,
	).Scan(&a.ID, &a.CreatedAt, &a.UpdatedAt)
}

func (a *Article) Exists(ctx context.Context) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM articles WHERE guid = $1)`
	var exists bool
	err := DB.QueryRow(ctx, query, a.GUID).Scan(&exists)
	return exists, err
}

func GetArticleByGUID(ctx context.Context, guid string) (*Article, error) {
	query := `
		SELECT id, guid, title, link, description, pub_date, created_at, updated_at
		FROM articles
		WHERE guid = $1`

	article := &Article{}
	err := DB.QueryRow(ctx, query, guid).Scan(
		&article.ID,
		&article.GUID,
		&article.Title,
		&article.Link,
		&article.Description,
		&article.PubDate,
		&article.CreatedAt,
		&article.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return article, nil
}
