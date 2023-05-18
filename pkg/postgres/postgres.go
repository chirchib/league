package postgres

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DB struct {
	DB     *sqlx.DB
	Config *DBConfig
}

type DBConfig struct {
	DBUser       string
	DBPassword   string
	DBHost       string
	DBPort       string
	DBName       string
	DBSchema     string
	DBSSLMode    string
	MaxIdleConns int
	MaxOpenConns int
	// в секундах
	ConnMaxLifetime int
	ApplicationName string
}

func NewDB(config *DBConfig) *DB {
	return &DB{
		Config: config,
	}
}

func (d *DB) Connect() (*sqlx.DB, error) {
	var (
		dataSourceName string
		err            error
	)

	dataSourceName = fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=%s search_path=%s application_name=%s",
		d.Config.DBUser,
		d.Config.DBPassword,
		d.Config.DBHost,
		d.Config.DBPort,
		d.Config.DBName,
		d.Config.DBSSLMode,
		d.Config.DBSchema,
		d.Config.ApplicationName,
	)

	if d.DB, err = sqlx.Connect("postgres", dataSourceName); err != nil {
		return d.DB, err
	}

	d.DB.SetMaxIdleConns(d.Config.MaxIdleConns)
	d.DB.SetMaxOpenConns(d.Config.MaxOpenConns)
	d.DB.SetConnMaxLifetime(time.Duration(d.Config.ConnMaxLifetime) * time.Second)

	return d.DB, nil
}
