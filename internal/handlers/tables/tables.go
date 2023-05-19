package tables

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Handler struct {
	// clubsService
	logger *zap.SugaredLogger
}

func New(
	logger *zap.SugaredLogger,
) *Handler {
	return &Handler{
		logger: logger,
	}
}

func (h *Handler) GetTable(c *gin.Context) {
	return
}

func (h *Handler) GetStat(c *gin.Context) {
	return
}

func (h *Handler) GetGames(c *gin.Context) {
	return
}

func (h *Handler) GetClubStats(c *gin.Context) {
	return
}

func (h *Handler) GetClubPlayers(c *gin.Context) {
	return
}
