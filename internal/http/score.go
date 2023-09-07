package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Mik3y-F/realtime-leaderboard-api/internal/repository"
	"github.com/gin-gonic/gin"
)

func (s *HttpServer) registerScoreRoutes() {
	s.router.POST("/scores", s.handlePublishScore)
	s.router.GET("/scores/top/:n", s.handleGetTopNPlayers)
}

func (s *HttpServer) handlePublishScore(c *gin.Context) {
	var score *repository.Score

	if err := c.BindJSON(&score); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := s.ScoreRepository.CreateScore(c.Request.Context(), score.PlayerID, score.Score)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (s *HttpServer) handleGetTopNPlayers(c *gin.Context) {

	n, err := strconv.Atoi(c.Param("n"))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	players, err := s.ScoreRepository.GetTopNPlayers(c.Request.Context(), int32(n))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, players)
}
