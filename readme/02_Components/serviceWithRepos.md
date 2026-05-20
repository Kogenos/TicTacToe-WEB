# serviceWithRepos

Пакет содержит **декоратор** над `MinimaxGameService`, который добавляет работу с репозиторием (сохранение/загрузку).

## Структура `GameServiceWithRepo`

- Поля: `repo repository.GameRepository` и `logic *service.MinimaxGameService`.
- Методы:
  - `CreateGame` – создаёт игру (с ботом или PvP) и сохраняет.
  - `JoinGame` – присоединяет игрока к ожидающей игре.
  - `MakeMovePlayer` – ход человека (загружает игру, проверяет права, вызывает `logic.MakeMove`, сохраняет, обновляет статус).
  - `MakeMoveBot` – ход бота (аналогично, но без проверки игрока).
  - `GetAvailableGames` – список ожидающих игр (через репозиторий).
  - `GetFinishedGamesByUser` – завершённые игры игрока.
  - `GetLeaderBoard` – топ игроков.

**Зачем нужен:** отделяет сохранение от чистой логики, инкапсулирует управление статусами игры.

См. также [service](service.md), [repository](repository.md).