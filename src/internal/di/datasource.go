package di

import (
	"TicTacToe/internal/datasource"
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/fx"
)

var DatasourceModule = fx.Module("datasource",
	fx.Provide(func(lc fx.Lifecycle) (*pgxpool.Pool, error) {
		connString := os.Getenv("DATABASE_URL")
		if connString == "" {
			connString = "postgres://postgres:password@localhost:5432/tictactoe?sslmode=disable"
		}

		pool, err := datasource.NewPgxPool(context.Background(), connString)
		if err != nil {
			return nil, err
		}
		lc.Append(fx.Hook{
			OnStop: func(ctx context.Context) error {
				pool.Close()
				return nil
			},
		})
		return pool, nil
	}),
)
