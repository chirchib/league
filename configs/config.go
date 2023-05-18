package configs

import (
	"strings"
)

type (
	Config struct {
		ENV       string `env:"ENV,required"`
		HTTPPort  int    `env:"HTTP_PORT,required"`
		PprofPort int    `env:"PPROF_PORT,required"`
		// db
		DBHost            string `env:"DB_HOST,required"`
		DBPort            string `env:"DB_PORT,required"`
		DBName            string `env:"DB_NAME,required"`
		DBUser            string `env:"DB_USER,required"`
		DBPassword        string `env:"DB_PASSWORD,required"`
		DBSSLMode         string `env:"DB_SSL_MODE,required"`
		DBMaxIdleConns    int    `env:"DB_MAX_IDLE_CONNS"`
		DBMaxOpenConns    int    `env:"DB_MAX_OPEN_CONNS"`
		DBConnMaxLifetime int    `env:"DB_MAX_LIFE_TIME"`
		DBSchema          string `env:"DB_SCHEMA" envDefault:"api_clients"`
		DBApplicationName string `env:"DB_APPLICATION_NAME"`
		// Migrations
		DBMigrationsDir string `env:"DB_MIGRATIONS_DIR,required"`

		// LogLevel
		LOGLEVEL      int8   `env:"LOGLEVEL"`
		LoggerUDPAddr string `env:"LOGGER_UDP_ADDR" envDefault:"127.0.0.1:12345"`

		Limits struct {
			// Максимальное время обработки запроса клиента REST
			MaxRESTRequestProcessingTimeMs int `env:"LIMITS_MAX_REST_REQUEST_PROCESSING_TIME_MS,required"`
		}
	}
)

func (c *Config) IsDevEnv() bool {
	return strings.EqualFold(c.ENV, "DEV")
}

func NewConfig() *Config {
	return &Config{}
}
