# Регистрация и JWT-логин

## Регистрация

1. Клиент → `POST /register` с JSON `{"login":"alice","password":"secret"}`.
2. `UserHandler.Register` читает `SignUpRequest`.
3. `UserService.Register`:
   - Проверяет длину (login≥3, password≥6).
   - Проверяет уникальность логина (через `UserRepository.FindByLogin`).
   - Хеширует пароль bcrypt.
   - Создаёт `domain.User` с новым UUID.
   - Сохраняет через `UserRepository.Save`.
4. Возвращает `{"user_id":"..."}`.

## Логин (получение JWT)

1. Клиент → `POST /login` с JSON `{"login":"alice","password":"secret"}`.
2. `UserHandler.Login` вызывает `AuthService.Authenticate`.
3. `AuthService.Authenticate`:
   - Вызывает `UserService.Authenticate` (проверяет пароль).
   - Генерирует access-токен (15 мин) через `JwtProvider.GenerateAccessToken`.
   - Генерирует refresh-токен (7 дней) через `JwtProvider.GenerateRefreshToken`.
4. Возвращает `JwtResponse` (`type: Bearer`, access_token, refresh_token).

См. также [JWT-токены (access, refresh)](../05_Implementation/JWT-токены (access, refresh).md), [[Чёрный список refresh-токенов]].