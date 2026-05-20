package di

import (
	"TicTacToe/internal/service"
	"TicTacToe/internal/serviceWithRepos"

	"go.uber.org/fx"
)

var ServiceModule = fx.Module("service",
	fx.Provide(service.NewMinimaxGameService),
	fx.Provide(serviceWithRepos.NewGameServiceWithRepo))
