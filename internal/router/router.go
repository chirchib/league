package router

import (
	"league/configs"
	clubsHandler "league/internal/handlers/clubs"
	playersHandler "league/internal/handlers/players"
	seasonsHandler "league/internal/handlers/seasons"
	tablesHandler "league/internal/handlers/tables"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

func Router(db *sqlx.DB, logger *zap.SugaredLogger, config *configs.Config) (*gin.Engine, error) {
	g := gin.New()

	g.Use(gin.Recovery())

	// init repositories

	// init services

	// init handlers
	clubsHandler := clubsHandler.New(logger)
	playersHandler := playersHandler.New(logger)
	seasonsHandler := seasonsHandler.New(logger)
	tablesHandler := tablesHandler.New(logger)

	d := &routingDependencies{
		config: config,
		logger: logger,

		clubsHandler:   clubsHandler,
		playersHandler: playersHandler,
		seasonsHandler: seasonsHandler,
		tablesHandler:  tablesHandler,
	}

	routes(g, d)

	return g, nil
}
