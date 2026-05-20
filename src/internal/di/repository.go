package di

import (
	"TicTacToe/internal/repository"

	"go.uber.org/fx"
)

var RepositoryModule = fx.Module("repository",
	fx.Provide(
		fx.Annotate(
			repository.NewGameRepositoryImpl,
			fx.As(new(repository.GameRepository)),
		),
	),
)
