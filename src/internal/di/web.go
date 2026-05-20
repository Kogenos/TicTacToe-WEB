package di

import (
	"TicTacToe/internal/web"
	"context"

	"go.uber.org/fx"
)

var WebModule = fx.Module("web",
	fx.Provide(web.NewGameHandler),
	fx.Provide(web.NewServer),
	fx.Invoke(RegisterServer))

func RegisterServer(lc fx.Lifecycle, srv *web.Server) {
	srv.SetupRoutes()
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go srv.Start()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			srv.Stop(ctx)
			return nil
		},
	})
}
