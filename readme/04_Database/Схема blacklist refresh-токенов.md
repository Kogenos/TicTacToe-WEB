
### 4_Database/Схема blacklist refresh-токенов



```sql
CREATE TABLE IF NOT EXISTS refresh_token_blacklist (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    token_hash TEXT UNIQUE NOT NULL,
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);
```

## Описание

- `token_hash` – SHA‑256 хеш refresh-токена. Уникален.
    
- `expires_at` – дата истечения оригинального токена (берётся из `exp` в JWT).
    
- `created_at` – время добавления в чёрный список.
    

**Зачем:** При каждом обновлении токенов старый refresh-токен добавляется в чёрный список, чтобы его нельзя было использовать повторно.

**Использование:**

- `IsBlacklisted` проверяет наличие хеша и что `expires_at > NOW()`.
    
- Периодически можно очищать записи с истекшим сроком (но необязательно).
    

См. также [Чёрный список refresh-токенов](../05_Implementation/Чёрный%20список%20refresh-токенов.md).
