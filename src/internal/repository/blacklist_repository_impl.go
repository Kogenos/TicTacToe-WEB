package repository

import (
	"context"
	_ "embed"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

//go:embed queries/create_blacklist_table.sql
var createBlacklistTableQuery string

//go:embed queries/add_to_blacklist.sql
var addToBlacklistQuery string

//go:embed queries/is_token_blacklisted.sql
var isTokenBlacklistedQuery string

type BlacklistRepositoryImpl struct {
	pool *pgxpool.Pool
}

func NewBlacklistRepositoryImpl(pool *pgxpool.Pool) (*BlacklistRepositoryImpl, error) {
	ctx := context.Background()
	if _, err := pool.Exec(ctx, createBlacklistTableQuery); err != nil {
		return nil, fmt.Errorf("failed to create blacklist table: %w", err)
	}
	return &BlacklistRepositoryImpl{pool: pool}, nil
}

func (r *BlacklistRepositoryImpl) Add(ctx context.Context, tokenHash string, expiresAt time.Time) error {
	_, err := r.pool.Exec(ctx, addToBlacklistQuery, tokenHash, expiresAt)
	return err
}

func (r *BlacklistRepositoryImpl) IsBlacklisted(ctx context.Context, tokenHash string) (bool, error) {
	var exists bool
	err := r.pool.QueryRow(ctx, isTokenBlacklistedQuery, tokenHash).Scan(&exists)
	return exists, err
}
