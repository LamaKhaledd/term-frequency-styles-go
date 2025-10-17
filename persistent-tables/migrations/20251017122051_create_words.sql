-- +goose Up
CREATE TABLE words (
    id INTEGER,
    doc_id INTEGER NOT NULL,
    value TEXT NOT NULL
);

-- +goose Down
DROP TABLE words;
