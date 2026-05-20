# web

HTTP-слой: хендлеры, DTO, middleware, сервер.

## Структуры

- `GameHandler` – обработчики игровых эндпоинтов.
- `UserHandler` – обработчики регистрации, логина, получения данных пользователя.
- `JWTAuthenticator` – middleware для Bearer-авторизации.
- `Server` – инкапсулирует все хендлеры, middleware, порт и маршруты.

## DTO (`models.go`)

- `GameDTO`, `MoveDTO`, `SignUpRequest`, `JwtRequest`, `JwtResponse`, `RefreshJwtRequest`, `ErrorResponse`.

## Мапперы (`mapper.go`)

- `ToGameDTO` – преобразует `domain.Game` в `GameDTO`.
- `ToDomainGame` – обратное (не используется, но оставлено).

## Эндпоинты (открытые)

- `POST /register` – регистрация.
- `POST /login` – получение JWT.
- `POST /refresh-access`, `POST /refresh-refresh` – обновление токенов.

## Эндпоинты (защищённые)

- `POST /game` – создать игру.
- `GET /game/{id}` – получить игру.
- `POST /game/{id}/move` – сделать ход.
- `POST /game/{id}/join` – присоединиться к игре.
- `GET /games/available` – список доступных игр.
- `GET /games` – все игры.
- `GET /games/finished` – завершённые игры пользователя.
- `GET /leaderboard` – таблица лидеров.
- `GET /user/me` – информация о текущем пользователе.
- `GET /user/{id}` – информация о другом пользователе.
- `GET /users` – список всех пользователей.

См. также [Регистрация и JWT-логин](../03_Flow/Регистрация и JWT-логин.md), [Создание игры (с ботом или PvP)](../03_Flow/Создание игры (с ботом или PvP).md).