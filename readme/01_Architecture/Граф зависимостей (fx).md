# Граф зависимостей (fx)

Проект использует **uber‑fx** для внедрения зависимостей. Модули определены в `internal/di`:

- `DatasourceModule` – предоставляет `*pgxpool.Pool`.
- `RepositoryModule` – предоставляет `GameRepository` (через `fx.Annotate`).
- `ServiceModule` – предоставляет `MinimaxGameService` и `GameServiceWithRepo`.
- `UserModule` – предоставляет `UserRepository`, `UserService`, `JwtProvider`, `AuthService`, `JWTAuthenticator`.
- `WebModule` – предоставляет `GameHandler`, `Server` и вызывает `RegisterServer`.

**Основной контейнер** – `di.App` (в `container.go`).

**Пример сборки:** `fx.New(di.App).Run()`.

**Ключевые принципы:**
- Каждый конструктор вызывается один раз (синглтон).
- `fx.Annotate` используется для привязки конкретной реализации к интерфейсу (`fx.As`).
- `fx.Invoke` запускает функции после построения всех зависимостей (например, регистрация маршрутов и старт сервера).

См. также [Слои приложения](Слои%20приложения.md).
