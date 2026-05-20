DROP TABLE if EXISTS users;

create TABLE if not exists users (
    id UUID PRIMARY KEY,
    login TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL
);