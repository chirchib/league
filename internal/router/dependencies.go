package router

import (
	"go.uber.org/zap"
	"league/configs"
)

type routingDependencies struct {
	config *configs.Config
	logger *zap.SugaredLogger
	// handlers
}
