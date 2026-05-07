package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"os"
)

type Server struct {
	Port string
}

type Postgres struct {
	PORT     string
	HOST     string
	USER     string
	PASSWORD string
	DBNAME   string
	SSL_MODE string
}

func CfgPostgres(logg *zap.Logger) (*Postgres, error) {
	err := godotenv.Load()
	if err != nil {
		logg.Error("Error loading .env file", zap.Error(err))
		return &Postgres{}, fmt.Errorf("config: Ошибка загрузки env %w", err)
	}
	return &Postgres{
		PORT:     os.Getenv("DB_PORT"),
		HOST:     os.Getenv("DB_HOST"),
		USER:     os.Getenv("DB_USER"),
		PASSWORD: os.Getenv("DB_PASSWORD"),
		DBNAME:   os.Getenv("DB_DBNAME"),
		SSL_MODE: os.Getenv("DB_SSL_MODE"),
	}, nil

}

func CfgServer(logg *zap.Logger) (*Server, error) {
	err := godotenv.Load()
	if err != nil {
		return &Server{}, fmt.Errorf("config: Ошибка загрузки env %w", err)
	}

	logg.Info("env конфиг загружен")

	return &Server{
		Port: os.Getenv("PORT_SERVER"),
	}, nil
}
