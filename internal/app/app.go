package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"league/configs"
	"league/internal/router"
	"league/pkg/logger"
	"league/pkg/postgres"
	"os"
	"os/signal"
	"syscall"

	"github.com/caarlos0/env/v6"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type App struct {
	logger *zap.SugaredLogger
	db     *sqlx.DB
	config *configs.Config
}

func NewApp(config *configs.Config) *App {
	return &App{
		config: config,
	}
}

func (a *App) Start() (*App, error) {
	var err error

	if err = a.ParseENV(); err != nil {
		return a, fmt.Errorf("can't parse ENV: %s", err)
	}

	if err = a.InitLogger(); err != nil {
		return a, fmt.Errorf("can't start logger: %s", err)
	}

	if err = a.InitDB(); err != nil {
		return a, fmt.Errorf("can't connect to db: %s", err)
	}

	if err = a.InitRouter(); err != nil {
		return a, fmt.Errorf("can't init router: %s", err)
	}

	a.gracefulDown()

	return a, nil
}

func (a *App) ParseENV() error {
	if err := env.Parse(a.config); err != nil {
		return err
	}

	return nil
}

func (a *App) InitLogger() error {
	var err error

	loggerConfig := logger.LoggerConfig{
		Level: zapcore.Level(a.config.Log.Level),
	}

	if a.logger, err = logger.NewLogger(&loggerConfig); err != nil {
		return err
	}

	defer a.logger.Sync()

	return nil
}

func (a *App) InitDB() error {
	var (
		dbConfig = &postgres.DBConfig{
			DBUser:          a.config.DB.User,
			DBPassword:      a.config.DB.Password,
			DBHost:          a.config.DB.Host,
			DBPort:          a.config.DB.Port,
			DBName:          a.config.DB.Name,
			DBSchema:        a.config.DB.Schema,
			DBSSLMode:       a.config.DB.SSLMode,
			MaxIdleConns:    a.config.DB.MaxIdleConns,
			MaxOpenConns:    a.config.DB.MaxOpenConns,
			ConnMaxLifetime: a.config.DB.ConnMaxLifetime,
		}
		err error
	)

	if a.db, err = postgres.NewDB(dbConfig).Connect(); err != nil {
		return err
	}

	return nil
}

func (a *App) InitRouter() error {
	var (
		g   *gin.Engine
		err error
	)

	g, err = router.Router(a.db, a.logger, a.config)
	if err != nil {
		return err
	}

	if err = g.Run(fmt.Sprint(":", a.config.HTTP.Port)); err != nil {
		return err
	}

	return nil
}

func (a *App) gracefulDown() {
	interrupt := make(chan os.Signal, 1)

	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGABRT)

	<-interrupt
}
