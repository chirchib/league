package players

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Handler struct {
	logger *zap.SugaredLogger
}

func New(
	logger *zap.SugaredLogger,
) *Handler {
	return &Handler{
		logger: logger,
	}
}

func (h *Handler) GetPlayers(c *gin.Context) {
	return
}

func (h *Handler) GetPlayer(c *gin.Context) {
	return
}

func (h *Handler) GetPlayerStats(c *gin.Context) {
	return
}
