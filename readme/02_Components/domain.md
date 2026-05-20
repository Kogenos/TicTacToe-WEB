# domain

Пакет содержит **бизнес-сущности** и константы. Никаких зависимостей от других слоёв.

## Модели

- `Game` – игра (поля: UUID, ID игроков, статус, победитель, поле, история, флаг бота и т.д.).
- `User` – пользователь (UUID, логин, хеш пароля).
- `LeaderBoardEntry` – запись для лидерборда (ID, логин, победы, всего игр, win_ratio).
- `Board` – матрица 3×3 ячеек (`Cell`).
- `Move` – ход (координаты, игрок).

## Константы

- `CellEmpty`, `CellX`, `CellO`.
- `StatusWaiting`, `StatusPlaying`, `StatusFinished`.

## Конструкторы

- `NewGame(playerX_ID)`
- `NewGameWithBot(playerX_ID, botSide)`

См. также [service](service.md) (чистая логика) и [repository](repository.md) (сохранение).