package mysql

import (
	"context"
	"fmt"

	"github.com/Mik3y-F/realtime-leaderboard-api/internal/mysql/generated"
	"github.com/Mik3y-F/realtime-leaderboard-api/internal/repository"
	"github.com/Mik3y-F/realtime-leaderboard-api/pkg"
)

var _ repository.PlayerRepository = (*PlayerRepository)(nil)

type PlayerRepository struct {
	db      *DB
	queries *generated.Queries
}

func NewPlayerRepository(db *DB) *PlayerRepository {

	// Create a new instance of the generated queries.
	queries := generated.New(db.db)

	return &PlayerRepository{
		db:      db,
		queries: queries,
	}
}

// CreatePlayer adds a new player to the database.
func (r *PlayerRepository) CreatePlayer(ctx context.Context, player *repository.Player) (string, error) {

	err := player.Validate()
	if err != nil {
		return "", pkg.Errorf(pkg.INVALID_ERROR, "invalid player details: %v", err)
	}

	result, err := r.queries.CreatePlayer(ctx, player.Name)
	if err != nil {
		return "", pkg.Errorf(pkg.INTERNAL_ERROR, "cannot create player: %v", err)
	}

	insertedPlayerID, err := result.LastInsertId()
	if err != nil {
		return "", pkg.Errorf(pkg.INTERNAL_ERROR, "cannot get inserted player ID: %v", err)
	}

	return fmt.Sprintf("%d", insertedPlayerID), nil
}

// GetPlayerByID retrieves a player based on their ID.
func (r *PlayerRepository) GetPlayerByID(ctx context.Context, playerID string) (*repository.Player, error) {
	result, err := r.queries.GetPlayerByID(ctx, playerID)
	if err != nil {
		return nil, pkg.Errorf(pkg.INTERNAL_ERROR, "cannot get player: %v", err)
	}

	return &repository.Player{
		Id:   result.ID,
		Name: result.Name,
	}, nil
}

// ListPlayers retrieves a list of players from the database.
func (r *PlayerRepository) ListPlayers(ctx context.Context) ([]*repository.Player, error) {
	results, err := r.queries.ListPlayers(ctx)
	if err != nil {
		return nil, pkg.Errorf(pkg.INTERNAL_ERROR, "cannot list players: %v", err)
	}

	players := make([]*repository.Player, len(results))
	for i, result := range results {
		players[i] = &repository.Player{
			Id:   result.ID,
			Name: result.Name,
		}
	}

	return players, nil
}

// UpdatePlayerDetails updates the details of a player.whose id has been passed as a parameter
func (r *PlayerRepository) UpdatePlayerDetails(
	ctx context.Context, playerID string, update repository.PlayerUpdate) (*repository.Player, error) {

	player, err := r.GetPlayerByID(ctx, playerID)
	if err != nil {
		return nil, err
	}

	if n := update.Name; n != nil {
		player.Name = *n
	}

	_, err = r.queries.UpdatePlayer(ctx, generated.UpdatePlayerParams{
		ID:   player.Id,
		Name: player.Name,
	})
	if err != nil {
		return nil, pkg.Errorf(pkg.INTERNAL_ERROR, "cannot update player: %v", err)
	}

	return &repository.Player{
		Id:   player.Id,
		Name: player.Name,
	}, nil
}

// DeletePlayer removes a player from the database based on their ID.
func (r *PlayerRepository) DeletePlayer(ctx context.Context, playerID string) error {
	_, err := r.queries.DeletePlayer(ctx, playerID)
	if err != nil {
		return pkg.Errorf(pkg.INTERNAL_ERROR, "cannot delete player: %v", err)
	}

	return nil
}
