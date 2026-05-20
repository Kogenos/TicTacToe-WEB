# TicTacToe Project – навигация

Добро пожаловать в базу знаний по проекту «Крестики-нолики». Здесь задокументирована архитектура, компоненты, потоки данных и ключевые решения.

## Архитектура
- [Слои приложения](01_Architecture/Слои%20приложения.md)
- [Граф зависимостей (fx)](01_Architecture/Граф%20зависимостей%20(fx).md)
- [База данных PostgreSQL](01_Architecture/База%20данных%20PostgreSQL.md)

## Компоненты
- [domain](02_Components/domain.md)
- [datasource](02_Components/datasource.md)
- [repository](02_Components/repository.md)
- [service](02_Components/service.md)
- [serviceWithRepos](02_Components/serviceWithRepos.md)
- [web](02_Components/web.md)
- [di](02_Components/di.md)

## Потоки (Use Cases)
- [Регистрация и JWT-логин](03_Flow/Регистрация%20и%20JWT-логин.md)
- [Создание игры (с ботом или PvP)](03_Flow/Создание%20игры%20(с%20ботом%20или%20PvP).md)
- [Присоединение к игре](03_Flow/Присоединение%20к%20игре.md)
- [Ход игрока (человек, бот)](03_Flow/Ход%20игрока%20(человек,%20бот).md)
- [Завершение игры и статусы](03_Flow/Завершение%20игры%20и%20статусы.md)
- [История игр пользователя](03_Flow/История%20игр%20пользователя.md)
- [Лидерборд](03_Flow/Лидерборд.md)

## База данных
- [Схема таблицы games](04_Database/Схема%20таблицы%20games.md)
- [Схема таблицы users](04_Database/Схема%20таблицы%20users.md)
- [Схема blacklist refresh-токенов](04_Database/Схема%20blacklist%20refresh-токенов.md)
- [Важные SQL-запросы (лидерборд, история)](04_Database/Важные%20SQL-запросы%20(лидерборд,%20история).md)

## Заметки по реализации
- [JWT-токены (access, refresh)](05_Implementation/JWT-токены%20(access,%20refresh).md)
- [Чёрный список refresh-токенов](05_Implementation/Чёрный%20список%20refresh-токенов.md)
- [Обработка ошибок](05_Implementation/Обработка%20ошибок.md)
- [Graceful shutdown](05_Implementation/Graceful%20shutdown.md)
- [Использование embed для SQL](05_Implementation/Использование%20embed%20для%20SQL.md)

## Глоссарий
- [Минимакс](06_Glossary/Минимакс.md)
- [Bcrypt](06_Glossary/Bcrypt.md)
- [JSONB](06_Glossary/JSONB.md)
- [Singleton](06_Glossary/Singleton.md)
