package router

import (
	"league/configs"
	clubsHandler "league/internal/handlers/clubs"
	playersHandler "league/internal/handlers/players"
	seasonsHandler "league/internal/handlers/seasons"
	tablesHandler "league/internal/handlers/tables"
	clubsService "league/internal/services/clubs"
	playersService "league/internal/services/players"
	seasonssService "league/internal/services/seasons"
	tablesService "league/internal/services/tables"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

func Router(db *sqlx.DB, logger *zap.SugaredLogger, config *configs.Config) (*gin.Engine, error) {
	g := gin.New()

	g.Use(gin.Recovery())

	// init repositories

	// init services
	clubsService := clubsService.New(logger)
	playersService := playersService.New(logger)
	seasonsService := seasonssService.New(logger)
	tablesService := tablesService.New(logger)

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
