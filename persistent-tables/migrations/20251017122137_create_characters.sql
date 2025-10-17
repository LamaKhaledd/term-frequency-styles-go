-- +goose Up
CREATE TABLE characters (
    id INTEGER,
    word_id INTEGER NOT NULL,
    value TEXT NOT NULL
);

-- +goose Down
DROP TABLE characters;
