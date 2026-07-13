# База данных PostgreSQL

- Используется библиотека `jackc/pgx/v5` и пул соединений `pgxpool`.
- Строка подключения задаётся через переменную окружения `DATABASE_URL`.
- Таблицы создаются автоматически при старте через `CREATE TABLE IF NOT EXISTS`.
- Для JSON-полей (`game_board`, `history`) используется тип `JSONB`.

**Таблицы:**
- [Схема таблицы games](../04_Database/Схема таблицы games.md)
- [Схема таблицы users](../04_Database/Схема таблицы users.md)
- [Схема blacklist refresh-токенов](../04_Database/Схема blacklist refresh-токенов.md)

**Миграции:** Встроены в код через `//go:embed`. При инициализации репозитория выполняются соответствующие SQL-скрипты (см. [Использование embed для SQL](../05_Implementation/Использование%20embed%20для%20SQL.md)).
