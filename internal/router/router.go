package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"league/configs"
)

func Router(db *sqlx.DB, logger *zap.SugaredLogger, config *configs.Config) (*gin.Engine, error) {
	g := gin.New()

	g.Use(gin.Recovery())

	// init repositories

	// init services

	// init handlers

	d := &routingDependencies{
		config: config,
		logger: logger,
	}

	routes(g, d)

	return g, nil
}
