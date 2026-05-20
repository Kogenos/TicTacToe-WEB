# Схема таблицы games

```sql
CREATE TABLE IF NOT EXISTS games (
    id UUID PRIMARY KEY,
    player_x_id UUID NOT NULL,
    player_o_id UUID,
    status TEXT NOT NULL,
    winner_id UUID,
    game_board JSONB NOT NULL,
    current_player INT NOT NULL,
    history JSONB NOT NULL,
    bot BOOLEAN NOT NULL,
    bot_side INT NOT NULL,
    create_date TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);
```

## Описание полей

- `id` – уникальный идентификатор игры.
    
- `player_x_id` – ID пользователя, играющего за X.
    
- `player_o_id` – ID пользователя, играющего за O (может быть NULL для ожидающей игры или для игры с ботом).
    
- `status` – `waiting`, `playing`, `finished`.
    
- `winner_id` – ID победителя или NULL/`uuid.Nil` при ничьей.
    
- `game_board` – JSON-массив 3×3 (значения 0,1,2).
    
- `current_player` – 1 (X) или 2 (O).
    
- `history` – JSON-массив ходов `[{"row":0,"col":0,"player":1}, ...]`.
    
- `bot` – флаг, игра с ботом.
    
- `bot_side` – за кого бот (1 или 2).
    
- `create_date` – дата создания (используется для сортировки истории).
    

## Индексы

- Первичный ключ по `id`.
    
- Рекомендуется индекс на `status` для ускорения выборки игр в ожидании.
    

См. также [Важные SQL-запросы (лидерборд, история)](Важные SQL-запросы (лидерборд, история).md).