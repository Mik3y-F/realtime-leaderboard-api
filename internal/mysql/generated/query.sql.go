// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: query.sql

package generated

import (
	"context"
	"database/sql"
	"time"
)

const createPlayer = `-- name: CreatePlayer :execresult
INSERT INTO players (name, created_at, updated_at)
VALUES (?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
`

func (q *Queries) CreatePlayer(ctx context.Context, name string) (sql.Result, error) {
	return q.db.ExecContext(ctx, createPlayer, name)
}

const createScoreEntry = `-- name: CreateScoreEntry :execresult
INSERT INTO score_entries (player_id, score, entry_time)
VALUES (?, ?, ?)
`

type CreateScoreEntryParams struct {
	PlayerID  string
	Score     int32
	EntryTime time.Time
}

func (q *Queries) CreateScoreEntry(ctx context.Context, arg CreateScoreEntryParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createScoreEntry, arg.PlayerID, arg.Score, arg.EntryTime)
}

const deletePlayer = `-- name: DeletePlayer :execresult
DELETE FROM players
WHERE id = ?
`

func (q *Queries) DeletePlayer(ctx context.Context, id string) (sql.Result, error) {
	return q.db.ExecContext(ctx, deletePlayer, id)
}

const getPlayerByID = `-- name: GetPlayerByID :one
SELECT id, name, created_at, updated_at
FROM players
WHERE id = ?
LIMIT 1
`

func (q *Queries) GetPlayerByID(ctx context.Context, id string) (Player, error) {
	row := q.db.QueryRowContext(ctx, getPlayerByID, id)
	var i Player
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getPlayerScoreEntries = `-- name: GetPlayerScoreEntries :many
SELECT entry_id, player_id, score, entry_time
FROM score_entries
WHERE player_id = ?
`

func (q *Queries) GetPlayerScoreEntries(ctx context.Context, playerID string) ([]ScoreEntry, error) {
	rows, err := q.db.QueryContext(ctx, getPlayerScoreEntries, playerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ScoreEntry
	for rows.Next() {
		var i ScoreEntry
		if err := rows.Scan(
			&i.EntryID,
			&i.PlayerID,
			&i.Score,
			&i.EntryTime,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPlayerTotalScore = `-- name: GetPlayerTotalScore :one
SELECT SUM(score) AS total_score
FROM score_entries
WHERE player_id = ?
`

func (q *Queries) GetPlayerTotalScore(ctx context.Context, playerID string) (interface{}, error) {
	row := q.db.QueryRowContext(ctx, getPlayerTotalScore, playerID)
	var total_score interface{}
	err := row.Scan(&total_score)
	return total_score, err
}

const getTopNPlayers = `-- name: GetTopNPlayers :many
SELECT p.id AS player_id,
    p.name,
    SUM(se.score) AS total_score
FROM score_entries se
    JOIN players p ON se.player_id = p.id
GROUP BY se.player_id,
    p.name
ORDER BY total_score DESC
LIMIT ?
`

type GetTopNPlayersRow struct {
	PlayerID   string
	Name       string
	TotalScore interface{}
}

func (q *Queries) GetTopNPlayers(ctx context.Context, limit int32) ([]GetTopNPlayersRow, error) {
	rows, err := q.db.QueryContext(ctx, getTopNPlayers, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetTopNPlayersRow
	for rows.Next() {
		var i GetTopNPlayersRow
		if err := rows.Scan(&i.PlayerID, &i.Name, &i.TotalScore); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listPlayers = `-- name: ListPlayers :many
SELECT id, name, created_at, updated_at
FROM players
`

func (q *Queries) ListPlayers(ctx context.Context) ([]Player, error) {
	rows, err := q.db.QueryContext(ctx, listPlayers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Player
	for rows.Next() {
		var i Player
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updatePlayer = `-- name: UpdatePlayer :execresult
UPDATE players
SET name = ?,
    updated_at = CURRENT_TIMESTAMP
WHERE id = ?
`

type UpdatePlayerParams struct {
	Name string
	ID   string
}

func (q *Queries) UpdatePlayer(ctx context.Context, arg UpdatePlayerParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updatePlayer, arg.Name, arg.ID)
}
