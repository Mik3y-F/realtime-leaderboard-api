-- name: GetPlayerByID :one
SELECT *
FROM players
WHERE id = ?
LIMIT 1;
-- name: CreatePlayer :execresult
INSERT INTO players (name, created_at, updated_at)
VALUES (?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
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
FROM players;
-- name: CreateScoreEntry :execresult
INSERT INTO score_entries (player_id, score, entry_time)
VALUES (?, ?, ?);
-- name: GetPlayerScoreEntries :many
SELECT *
FROM score_entries
WHERE player_id = ?;
-- name: GetPlayerTotalScore :one
SELECT SUM(score) AS total_score
FROM score_entries
WHERE player_id = ?;
-- name: GetTopNPlayers :many
SELECT p.id AS player_id,
    p.name,
    SUM(se.score) AS total_score
FROM score_entries se
    JOIN players p ON se.player_id = p.id
GROUP BY se.player_id,
    p.name
ORDER BY total_score DESC
LIMIT ?;
