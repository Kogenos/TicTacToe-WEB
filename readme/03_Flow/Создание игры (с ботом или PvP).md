# Создание игры (с ботом или PvP)

1. Клиент → `POST /game` с заголовком `Authorization: Bearer <access_token>` и JSON-телом:
   - `{"bot": true, "botside": "X"}` – игра с ботом, бот за X.
   - `{"bot": false}` – PvP-игра (ожидание игрока).
2. `GameHandler.CreateGame`:
   - Извлекает `userID` из контекста (установлен middleware).
   - Вызывает `GameServiceWithRepo.CreateGame`.
3. Сервис:
   - Создаёт игру через `domain.NewGame(userID)` или `domain.NewGameWithBot(userID, botSide)`.
   - Если игра с ботом и `botside == CellX` – сразу делает ход бота (через `logic.GetNextMove` и `logic.MakeMove`).
   - Сохраняет игру через репозиторий.
4. Возвращает `GameDTO` с полным состоянием игры.

**Особенности:** PvP-игра создаётся со статусом `waiting`. Бот-игра – сразу `playing`.

См. также [Присоединение к игре](Присоединение%20к%20игре.md), [Ход игрока (человек, бот)](Ход%20игрока%20(человек,%20бот).md).
