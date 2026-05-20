package di

import (
	"TicTacToe/internal/repository"
	"TicTacToe/internal/service"
	"TicTacToe/internal/web"
	"os"

	"go.uber.org/fx"
)

var UserModule = fx.Module("user",

	fx.Provide(
		fx.Annotate(
			repository.NewUserRepositoryImpl,
			fx.As(new(repository.UserRepository)),
		),
	),

	fx.Provide(
		fx.Annotate(
			service.NewUserService,
			fx.As(new(service.UserService)),
		),
	),

	fx.Provide(web.NewUserHandler),
	fx.Provide(func() *service.JwtProvider {
		secret := os.Getenv("JWT_SECRET")
		if secret == "" {
			secret = "default-secret-key-change-me"
		}
		return service.NewJwtProvider(secret)
	}),

	fx.Provide(
		fx.Annotate(
			service.NewAuthService,
			fx.As(new(service.AuthService)),
		),
	),

	fx.Provide(web.NewJWTAuthenticator),
	fx.Provide(
		fx.Annotate(
			repository.NewBlacklistRepositoryImpl,
			fx.As(new(repository.BlacklistRepository)),
		),
	),
)
