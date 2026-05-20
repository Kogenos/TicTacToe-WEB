# service

Пакет содержит **чистую бизнес-логику**, не зависящую от хранилища и HTTP.

## Ключевые компоненты

- `MinimaxGameService` – реализует алгоритм минимакса, ходы, проверку окончания игры.
  - `MakeMove` – применяет ход (проверяет, обновляет поле, переключает игрока).
  - `GetNextMove` – находит лучший ход для бота.
  - `GameFinished` – определяет победу/ничью.
  - `ValidateGame` – проверяет соответствие истории ходов полю (не используется в API, но оставлено).
- `JwtProvider` – генерация и валидация JWT (см. [JWT-токены (access, refresh)](../05_Implementation/JWT-токены (access, refresh).md)).
- `UserService` – регистрация и аутентификация (bcrypt).
- `AuthService` – объединяет `UserService` и `JwtProvider` для выдачи и обновления токенов.

**Принцип:** Методы `MinimaxGameService` работают с `domain.Game` и не сохраняют изменения. За сохранение отвечает [serviceWithRepos](serviceWithRepos.md).

См. также [Минимакс](../06_Glossary/Минимакс.md), [Bcrypt](../06_Glossary/Bcrypt.md).