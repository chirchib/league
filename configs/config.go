package configs

import (
	"github.com/google/uuid"
)

type (
	Config struct {
		HTTP HTTP
		DB   DB
		Log  Log

		ServiceNodeID string
	}

	HTTP struct {
		Port int `env:"HTTP_PORT,required"`
	}

	DB struct {
		Host            string `env:"DB_HOST,required"`
		Port            string `env:"DB_PORT,required"`
		Name            string `env:"DB_NAME,required"`
		User            string `env:"DB_USER,required"`
		Password        string `env:"DB_PASSWORD,required"`
		SSLMode         string `env:"DB_SSL_MODE,required"`
		MaxIdleConns    int    `env:"DB_MAX_IDLE_CONNS"`
		MaxOpenConns    int    `env:"DB_MAX_OPEN_CONNS"`
		ConnMaxLifetime int    `env:"DB_MAX_LIFE_TIME"`
		Schema          string `env:"DB_SCHEMA" envDefault:"api_clients"`
	}

	Log struct {
		Level int    `env:"LOG_LEVEL"`
		UDP   string `env:"LOG_UDP"`
	}
)

func NewConfig() *Config {
	return &Config{
		ServiceNodeID: uuid.New().String(),
	}
}
