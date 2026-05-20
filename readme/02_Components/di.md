# di

Содержит модули fx для сборки зависимостей.

- `container.go` – главный модуль `App`, объединяющий все подмодули.
- `datasource.go` – предоставляет `*pgxpool.Pool`.
- `repository.go` – предоставляет `GameRepository`.
- `service.go` – предоставляет `MinimaxGameService` и `GameServiceWithRepo`.
- `user.go` – предоставляет `UserRepository`, `UserService`, `JwtProvider`, `AuthService`, `JWTAuthenticator`, `BlacklistRepository`.
- `web.go` – предоставляет `GameHandler`, `Server` и регистрирует сервер через `fx.Invoke`.

**Использование аннотаций:** `fx.Annotate` + `fx.As` используется для приведения конкретных типов к интерфейсам (например, `NewUserRepositoryImpl` → `UserRepository`).

См. также [[Граф зависимостей (fx)]].