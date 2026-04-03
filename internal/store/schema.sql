CREATE TABLE IF NOT EXISTS entries (
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    title       TEXT,
    input       TEXT NOT NULL,
    output      TEXT,
    model       TEXT,
    created_at  DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    reveal_at   DATETIME,
    visible     BOOLEAN NOT NULL DEFAULT FALSE
);