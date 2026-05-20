package repository

import (
	"context"
	"time"
)

type BlacklistRepository interface {
	Add(ctx context.Context, tokenHash string, expiresAt time.Time) error
	IsBlacklisted(ctx context.Context, tokenHash string) (bool, error)
}
