package service

import "context"

type Player struct {
	ID    string
	Name  string
	Score int
}

type CachingService interface {
	// GetTopNPlayers retrieves the top 'N' players from the Redis leaderboard.
	GetTopNPlayers(ctx context.Context, n int) ([]*Player, error)

	// CachePlayerScore caches the score of a given player.
	CachePlayerScore(ctx context.Context, playerID string, score int) error
}
