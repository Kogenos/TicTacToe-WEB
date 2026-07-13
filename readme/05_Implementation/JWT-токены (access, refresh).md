# JWT-токены (access, refresh)
## Генерация
- Access token – время жизни **15 минут**.
- Refresh token – время жизни **7 дней**.
- Алгоритм подписи: HMAC‑SHA256.
- Секретный ключ читается из переменной окружения `JWT_SECRET`.
## Структура claims
- `sub` – UUID пользователя.
- `exp` – Unix timestamp.
- `type` – `"access"` или `"refresh"` (для разделения).
## Валидация
- Проверяется подпись, срок, тип.
- Для refresh-токенов дополнительно проверяется чёрный список (см. [[Чёрный список refresh-токенов]]).
## Использование
- Access token передаётся в заголовке `Authorization: Bearer <access_token>`.
- Refresh token передаётся только в теле запросов обновления (`/refresh-access`, `/refresh-refresh`).
## Методы `JwtProvider`
- `GenerateAccessToken(userID)`, `GenerateRefreshToken(userID)`
- `ValidateAccessToken(tokenString)`, `ValidateRefreshToken(tokenString)`
- `GetUserIDFromToken(token)`
- `HashToken(tokenString)` – для чёрного списка.

См. также [Регистрация и JWT-логин](../03_Flow/Регистрация%20и%20JWT-логин.md).
