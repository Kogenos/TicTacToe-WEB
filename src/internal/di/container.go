package di

import "go.uber.org/fx"

var App = fx.Module("tic-tac-toe",
	DatasourceModule,
	RepositoryModule,
	ServiceModule,
	WebModule,
	UserModule,
)
