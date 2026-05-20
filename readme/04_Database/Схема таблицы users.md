
### 4_Database/Схема таблицы users

```sql
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    login TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL
);
```

## Описание

- `id` – уникальный идентификатор пользователя.
    
- `login` – уникальное имя пользователя.
    
- `password_hash` – хеш пароля, созданный bcrypt.
    

## Индексы

- Первичный ключ по `id`.
    
- Уникальный индекс по `login` для быстрого поиска и ограничения дублей.
    

См. также [Bcrypt](../06_Glossary/Bcrypt.md).