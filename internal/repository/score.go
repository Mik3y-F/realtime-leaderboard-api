package repository

import "context"

type Score struct {
	PlayerID  string `json:"player_id"`
	Score     int32  `json:"score"`
	TimeStamp int64  `json:"timestamp,omitempty"`
}

type ScoreRepository interface {
	// GetTopNPlayers retrieves the top 'N' players from the Redis leaderboard.
	GetTopNPlayers(ctx context.Context, n int32) ([]*Score, error)

	// CreateScore creates a new score for a given player.
	CreateScore(ctx context.Context, playerID string, score int32) error
}
