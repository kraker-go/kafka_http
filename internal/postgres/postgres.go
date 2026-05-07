package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"go.uber.org/zap"
	"kafka_http/config"
)

func NewPostgres(logg *zap.Logger, cfg config.Postgres) (*sql.DB, error) {
	str := fmt.Sprintf("port=%s host=%s user=%s password=%s dbname=%s sslmode=%s", cfg.PORT, cfg.HOST, cfg.USER, cfg.PASSWORD, cfg.DBNAME, cfg.SSL_MODE)
	postgres, err := sql.Open("postgres", str)
	if err != nil {
		return nil, fmt.Errorf("ошибка подключения postgres: %w", err)
	}
	println(str)

	err = postgres.Ping()
	if err != nil {
		return nil, fmt.Errorf("ошибка соединения с postgres: %w", err)
	}

	logg.Info("постгрес запущен")

	return postgres, nil
}

func MigrationDB(postgres *sql.DB) error {
	return goose.Up(postgres, "migrations")
}
