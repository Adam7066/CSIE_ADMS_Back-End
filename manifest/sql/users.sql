CREATE TABLE users
(
    id         INTEGER PRIMARY KEY AUTOINCREMENT,
    email      TEXT      NOT NULL UNIQUE,
    password   TEXT      NOT NULL,
    username   TEXT      NOT NULL,
    role       TEXT               DEFAULT 'user' NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);