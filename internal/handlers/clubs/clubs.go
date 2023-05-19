package clubs

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

func (h *Handler) GetClubs(c *gin.Context) {
	return
}

func (h *Handler) GetClub(c *gin.Context) {
	return
}

func (h *Handler) GetClubPlayers(c *gin.Context) {
	return
}

func (h *Handler) GetClubTitles(c *gin.Context) {
	return
}

func (h *Handler) GetClubGames(c *gin.Context) {
	return
}
