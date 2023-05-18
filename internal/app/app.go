package app

import (
	"fmt"
	"time"
	
	"league/configs"
	"league/pkg/logger"
	"league/pkg/postgres"

	"github.com/caarlos0/env/v6"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type App struct {
	Logger *zap.SugaredLogger
	DB     *sqlx.DB
	Config *configs.Config
}

func NewApp(config *configs.Config) *App {
	return &App{
		Config: config,
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

	return a, nil
}

func (a *App) InitLogger() error {
	var err error

	loggerConfig := logger.LoggerConfig{
		Level: zapcore.Level(a.Config.LOGLEVEL),
	}

	if a.Logger, err = logger.NewLogger(&loggerConfig); err != nil {
		return err
	}

	defer a.Logger.Sync()

	return nil
}

func (a *App) ParseENV() error {
	err := env.Parse(a.Config)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) InitDB() error {
	var (
		dbConfig = &postgres.DBConfig{
			DBUser:          a.Config.DBUser,
			DBPassword:      a.Config.DBPassword,
			DBHost:          a.Config.DBHost,
			DBPort:          a.Config.DBPort,
			DBName:          a.Config.DBName,
			DBSchema:        a.Config.DBSchema,
			DBSSLMode:       a.Config.DBSSLMode,
			MaxIdleConns:    a.Config.DBMaxIdleConns,
			MaxOpenConns:    a.Config.DBMaxOpenConns,
			ConnMaxLifetime: a.Config.DBConnMaxLifetime,
			ApplicationName: a.Config.DBApplicationName,
		}
		err error
	)

	a.Logger.Named("DB").Infof(
		"MaxOpenConns: %d; MaxIdleConns: %d; ConnMaxLifetime: %d seconds",
		a.Config.DBMaxOpenConns,
		a.Config.DBMaxIdleConns,
		a.Config.DBConnMaxLifetime,
	)

	if a.DB, err = postgres.NewDB(dbConfig).Connect(); err != nil {
		return err
	}

	a.DB.SetConnMaxIdleTime(60 * time.Second)

	return nil
}
