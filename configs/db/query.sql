-- name: GetPlayerByID :one
SELECT *
FROM players
WHERE id = ?
LIMIT 1;
-- name: CreatePlayer :execresult
INSERT INTO players (name)
VALUES (?);
-- name: UpdatePlayer :execresult
UPDATE players
SET name = ?,
    updated_at = CURRENT_TIMESTAMP
WHERE id = ?;
-- name: DeletePlayer :execresult
DELETE FROM players
WHERE id = ?;
-- name: ListPlayers :many
SELECT *
FROM players
