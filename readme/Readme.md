# TicTacToe Project – навигация

Добро пожаловать в базу знаний по проекту «Крестики-нолики». Здесь задокументирована архитектура, компоненты, потоки данных и ключевые решения.

## Архитектура
- [Слои приложения](01_Architecture/Слои приложения.md)
- [[Граф зависимостей (fx)]]
- [База данных PostgreSQL](01_Architecture/База данных PostgreSQL.md)

## Компоненты
- [domain](02_Components/domain.md)
- [datasource](02_Components/datasource.md)
- [repository](02_Components/repository.md)
- [service](02_Components/service.md)
- [serviceWithRepos](02_Components/serviceWithRepos.md)
- [web](02_Components/web.md)
- [di](02_Components/di.md)

## Потоки (Use Cases)
- [Регистрация и JWT-логин](03_Flow/Регистрация и JWT-логин.md)
- [Создание игры (с ботом или PvP)](03_Flow/Создание игры (с ботом или PvP).md)
- [Присоединение к игре](03_Flow/Присоединение к игре.md)
- [Ход игрока (человек, бот)](03_Flow/Ход игрока (человек, бот).md)
- [Завершение игры и статусы](03_Flow/Завершение игры и статусы.md)
- [История игр пользователя](03_Flow/История игр пользователя.md)
- [Лидерборд](03_Flow/Лидерборд.md)

## База данных
- [Схема таблицы games](04_Database/Схема таблицы games.md)
- [Схема таблицы users](04_Database/Схема таблицы users.md)
- [Схема blacklist refresh-токенов](04_Database/Схема blacklist refresh-токенов.md)
- [Важные SQL-запросы (лидерборд, история)](04_Database/Важные SQL-запросы (лидерборд, история).md)

## Заметки по реализации
- [JWT-токены (access, refresh)](05_Implementation/JWT-токены (access, refresh).md)
- [[Чёрный список refresh-токенов]]
- [Обработка ошибок](05_Implementation/Обработка ошибок.md)
- [Graceful shutdown](05_Implementation/Graceful shutdown.md)
- [Использование embed для SQL](05_Implementation/Использование embed для SQL.md)

## Глоссарий
- [Минимакс](06_Glossary/Минимакс.md)
- [Bcrypt](06_Glossary/Bcrypt.md)
- [JSONB](06_Glossary/JSONB.md)
- [Singleton](06_Glossary/Singleton.md)