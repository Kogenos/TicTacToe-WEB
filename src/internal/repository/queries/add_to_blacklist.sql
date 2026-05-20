INSERT INTO refresh_token_blacklist (token_hash, expires_at)
VALUES ($1, $2)
ON CONFLICT (token_hash) DO NOTHING;