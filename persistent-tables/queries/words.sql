-- name: InsertDocument :exec
INSERT INTO documents (name) VALUES (?);

-- name: GetDocumentByName :one
SELECT id, name FROM documents WHERE name = ?;

-- name: InsertWord :exec
INSERT INTO words (id, doc_id, value) VALUES (?, ?, ?);

-- name: InsertCharacter :exec
INSERT INTO characters (id, word_id, value) VALUES (?, ?, ?);

-- name: TopWords :many
SELECT value, COUNT(*) as count
FROM words
GROUP BY value
ORDER BY count DESC
LIMIT 25;
