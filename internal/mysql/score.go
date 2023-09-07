package mysql

import (
	"context"
	"time"

	"github.com/Mik3y-F/realtime-leaderboard-api/internal/mysql/generated"
	"github.com/Mik3y-F/realtime-leaderboard-api/internal/repository"
	"github.com/Mik3y-F/realtime-leaderboard-api/pkg"
)

var _ repository.ScoreRepository = (*ScoreRepository)(nil)

type ScoreRepository struct {
	db      *DB
	queries *generated.Queries
}

func NewScoreRepository(db *DB) *ScoreRepository {

	quries := generated.New(db.db)

	return &ScoreRepository{
		db:      db,
		queries: quries,
	}
}

// GetTopNPlayers retrieves the top 'N' players from the Redis leaderboard.
func (r *ScoreRepository) GetTopNPlayers(ctx context.Context, n int32) ([]*repository.Score, error) {
	results, err := r.queries.GetTopNPlayers(ctx, n)
	if err != nil {
		return nil, pkg.Errorf(pkg.INTERNAL_ERROR, "cannot get top N players: %v", err)
	}

	scores := make([]*repository.Score, len(results))
	for _, score := range results {

		t := score.TotalScore
		totalScore, ok := t.(int32)
		if !ok {
			return nil, pkg.Errorf(pkg.INTERNAL_ERROR, "cannot convert score to int64")
		}

		scores = append(scores, &repository.Score{
			PlayerID: score.PlayerID,
			Score:    totalScore,
		})
	}

	return scores, nil
}

// CreateScore creates a new score for a given player.
func (r *ScoreRepository) CreateScore(ctx context.Context, playerID string, score int32) error {

	timeNow := time.Now()

	_, err := r.queries.CreateScoreEntry(ctx, generated.CreateScoreEntryParams{
		PlayerID:  playerID,
		Score:     score,
		EntryTime: timeNow,
	})
	if err != nil {
		return pkg.Errorf(pkg.INTERNAL_ERROR, "cannot create score: %v", err)
	}

	return nil
}
