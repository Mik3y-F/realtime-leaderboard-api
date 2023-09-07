package repository

type Score struct {
	PlayerID  string
	Score     int
	TimeStamp int64
}

type ScoreRepository interface {
	// GetTopNPlayers retrieves the top 'N' players from the Redis leaderboard.
	GetTopNPlayers(n int) ([]*Score, error)

	// GetPlayerScore retrieves the score of a given player.
	GetPlayerScore(playerID string) (*Score, error)

	// CreateScore creates a new score for a given player.
	CreateScore(playerID string, score int) error
}
