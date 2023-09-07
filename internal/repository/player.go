package repository

import (
	"context"

	"github.com/Mik3y-F/realtime-leaderboard-api/pkg"
)

type Player struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func (p *Player) Validate() error {
	if p.Name == "" {
		return pkg.Errorf(pkg.INVALID_ERROR, "name is empty")
	}

	return nil
}

type PlayerUpdate struct {
	Name *string `json:"name"`
}

type PlayerRepository interface {
	// CreatePlayer adds a new player to the database.
	CreatePlayer(ctx context.Context, player *Player) (string, error)

	// GetPlayerByID retrieves a player based on their ID.
	GetPlayerByID(ctx context.Context, playerID string) (*Player, error)

	// DeletePlayer removes a player from the database based on their ID.
	DeletePlayer(ctx context.Context, playerID string) error

	// ListPlayers retrieves a list of players from the database.
	ListPlayers(ctx context.Context) ([]*Player, error)

	// UpdatePlayerDetails updates the details of a player.
	UpdatePlayerDetails(ctx context.Context, playerID string, player PlayerUpdate) (*Player, error)

	// GetPlayerScore retrieves the score of a given player.
	GetPlayerTotalScore(ctx context.Context, playerID string) (*Score, error)
}
