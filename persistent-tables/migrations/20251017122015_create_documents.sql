-- +goose Up
CREATE TABLE documents (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL
);

-- +goose Down
DROP TABLE documents;
