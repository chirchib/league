package router

import (
	"go.uber.org/zap"
	"league/configs"
	"league/internal/handlers/clubs"
	"league/internal/handlers/players"
	"league/internal/handlers/seasons"
	"league/internal/handlers/tables"
)

type routingDependencies struct {
	config *configs.Config
	logger *zap.SugaredLogger
	// handlers
	clubsHandler   *clubs.Handler
	playersHandler *players.Handler
	seasonsHandler *seasons.Handler
	tablesHandler  *tables.Handler
}
