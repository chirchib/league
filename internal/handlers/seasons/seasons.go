package seasons

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

func (h *Handler) GetSeasons(c *gin.Context) {
	return
}

func (h *Handler) GetSeason(c *gin.Context) {
	return
}
