package http

import (
	"net/http"

	"github.com/Mik3y-F/realtime-leaderboard-api/internal/repository"
	"github.com/gin-gonic/gin"
)

func (s *HttpServer) registerPlayerRoutes() {
	s.router.POST("/players", s.handleCreatePlayer)
	s.router.GET("/players", s.handleListPlayers)

	s.router.GET("/players/:id", s.handleGetPlayerByID)
	s.router.DELETE("/players/:id", s.handleDeletePlayer)

	s.router.GET("/players/:id/score", s.handleGetPlayerScore)

	s.router.PUT("/players/:id", s.handleUpdatePlayerDetails)

}

func (s *HttpServer) handleCreatePlayer(c *gin.Context) {
	var player *repository.Player

	if err := c.BindJSON(&player); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	playerID, err := s.PlayerRepository.CreatePlayer(c.Request.Context(), player)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": playerID})
}

func (s *HttpServer) handleGetPlayerByID(c *gin.Context) {
	playerID := c.Param("id")

	player, err := s.PlayerRepository.GetPlayerByID(c.Request.Context(), playerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, player)
}

func (s *HttpServer) handleListPlayers(c *gin.Context) {
	players, err := s.PlayerRepository.ListPlayers(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, players)
}

func (s *HttpServer) handleDeletePlayer(c *gin.Context) {
	playerID := c.Param("id")

	err := s.PlayerRepository.DeletePlayer(c.Request.Context(), playerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (s *HttpServer) handleUpdatePlayerDetails(c *gin.Context) {
	playerID := c.Param("id")

	var update repository.PlayerUpdate
	if err := c.BindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	player, err := s.PlayerRepository.UpdatePlayerDetails(c.Request.Context(), playerID, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, player)
}

func (s *HttpServer) handleGetPlayerScore(c *gin.Context) {
	playerID := c.Param("id")

	score, err := s.PlayerRepository.GetPlayerTotalScore(c.Request.Context(), playerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, score)
}
