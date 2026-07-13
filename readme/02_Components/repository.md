# repository

Слой доступа к данным. Содержит интерфейсы и их реализации на PostgreSQL.

## Интерфейсы

- `GameRepository` – методы `Save`, `Load`, `ListGames`, `ListWaitingGames`, `ListFinishedGamesByUser`, `GetLeaderBoard`.
- `UserRepository` – `Save`, `FindByLogin`, `FindByID`, `ListUsers`.
- `BlacklistRepository` – `Add`, `IsBlacklisted`.

## Реализации

- `GameRepositoryImpl` – использует `*pgxpool.Pool` и SQL-запросы (см. [Использование embed для SQL](../05_Implementation/Использование embed для SQL.md)).
- `UserRepositoryImpl` – аналогично.
- `BlacklistRepositoryImpl` – работа с таблицей `refresh_token_blacklist`.

**Важные особенности:**
- В `Save` игры используется `ON CONFLICT` (upsert).
- JSON-поля сериализуются/десериализуются через `json.Marshal/Unmarshal`.
- Все методы принимают `context.Context`.

См. также [База данных PostgreSQL](../01_Architecture/База%20данных%20PostgreSQL.md).
