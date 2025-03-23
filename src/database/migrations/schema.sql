CREATE TABLE IF NOT EXISTS todos
(
    id        INTEGER PRIMARY KEY AUTOINCREMENT,
    title     TEXT    NOT NULL,
    completed BOOLEAN NOT NULL DEFAULT false
);
