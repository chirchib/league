package clubs

import "go.uber.org/zap"

type Service struct {
	logger *zap.SugaredLogger
}

func New(
	logger *zap.SugaredLogger,
) *Service {
	return &Service{
		logger: logger}
}
