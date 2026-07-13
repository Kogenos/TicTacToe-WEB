# Использование embed для SQL

В проекте все SQL-запросы вынесены в отдельные `.sql` файлы в папке `internal/repository/queries/`. Эти файлы встраиваются в бинарник с помощью директивы `//go:embed`.

## Преимущества

- Подсветка синтаксиса SQL в IDE.
- Легко редактировать и тестировать запросы.
- SQL не смешивается с Go-кодом.
- Бинарный файл остаётся самодостаточным.

## Пример

```go
//go:embed queries/save_game.sql
var saveGameQuery string
```

## Список SQL-файлов

- `create_table.sql`, `create_users_table.sql`, `create_blacklist_table.sql`
    
- `save_game.sql`, `load_game.sql`, `list_games.sql`, `list_waiting_games.sql`
    
- `list_finished_games_by_user.sql`, `leaderboard.sql`
    
- `save_user.sql`, `find_user_by_login.sql`, `find_user_by_id.sql`, `list_users.sql`
    
- `add_to_blacklist.sql`, `is_token_blacklisted.sql`
    

См. также [Схема таблицы games](../04_Database/Схема%20таблицы%20games.md), [Важные SQL-запросы (лидерборд, история)](../04_Database/Важные%20SQL-запросы%20(лидерборд,%20история).md).
