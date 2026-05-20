# TicTacToe Project – навигация

Добро пожаловать в базу знаний по проекту «Крестики-нолики». Здесь задокументирована архитектура, компоненты, потоки данных и ключевые решения.

## Архитектура
- [Слои_приложения](01_Architecture/Слои_приложения.md)
- [Граф зависимостей (fx)](Граф_зависимостей_(fx).md)
- [База данных PostgreSQL](01_Architecture/База_данных_PostgreSQL.md)

## Компоненты
- [domain](02_Components/domain.md)
- [datasource](02_Components/datasource.md)
- [repository](02_Components/repository.md)
- [service](02_Components/service.md)
- [serviceWithRepos](02_Components/serviceWithRepos.md)
- [web](02_Components/web.md)
- [di](02_Components/di.md)

## Потоки (Use Cases)
- [Регистрация и JWT-логин](03_Flow/Регистрация_и_JWT-логин.md)
- [Создание игры (с ботом или PvP)](03_Flow/Создание_игры_(с_ботом_или_PvP).md)
- [Присоединение к игре](03_Flow/Присоединение_к_игре.md)
- [Ход игрока (человек, бот)](03_Flow/Ход_игрока_(человек,_бот).md)
- [Завершение игры и статусы](03_Flow/Завершение_игры_и_статусы.md)
- [История игр пользователя](03_Flow/История_игр_пользователя.md)
- [Лидерборд](03_Flow/Лидерборд.md)

## База данных
- [Схема таблицы games](04_Database/Схема_таблицы_games.md)
- [Схема таблицы users](04_Database/Схема_таблицы_users.md)
- [Схема blacklist refresh-токенов](04_Database/Схема_blacklist_refresh-токенов.md)
- [Важные SQL-запросы (лидерборд, история)](04_Database/Важные_SQL-запросы_(лидерборд,_история).md)

## Заметки по реализации
- [JWT-токены (access, refresh)](05_Implementation/JWT-токены_(access,_refresh).md)
- [Чёрный список refresh-токенов](05_Implementation/Чёрный_список_refresh-токенов.md)
- [Обработка ошибок](05_Implementation/Обработка_ошибок.md)
- [Graceful shutdown](05_Implementation/Graceful_shutdown.md)
- [Использование embed для SQL](05_Implementation/Использование_embed_для_SQL.md)

## Глоссарий
- [Минимакс](06_Glossary/Минимакс.md)
- [Bcrypt](06_Glossary/Bcrypt.md)
- [JSONB](06_Glossary/JSONB.md)
- [Singleton](06_Glossary/Singleton.md)
