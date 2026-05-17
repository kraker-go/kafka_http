package app

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"kafka_http/config"
	"kafka_http/internal/cache"
	"kafka_http/internal/handler"
	"kafka_http/internal/kafka"
	"kafka_http/internal/logger"
	"kafka_http/internal/postgres"
	"kafka_http/internal/redis"
	"kafka_http/internal/repository"
	"kafka_http/internal/router"
	"kafka_http/internal/service"
	"kafka_http/server"
	"net/http"
)

func Run(logg *zap.Logger) error {
	conf, err := config.CfgPostgres(logg)
	if err != nil {
		return fmt.Errorf("App Ошибка загрузки сонфигурации %w", err)
	}

	postgresDB, err := postgres.NewPostgres(logg, *conf)
	if err != nil {
		return fmt.Errorf("App ошибка загрузки postgres %w", err)
	}

	logg.Info("Postgres подключен")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := postgres.MigrationDB(postgresDB); err != nil {
		return fmt.Errorf("App ошибка создания миграции %w", err)
	}

	logg.Info("миграции установлены")

	port, err := config.CfgServer(logg)
	if err != nil {
		return fmt.Errorf("App Ошибка чтения порта %w", err)
	}

	producer := kafka.NewKafkaProducer()

	go func() {
		consumer := kafka.NewKafkaConsumer()
		if err := consumer.Start(ctx, logg); err != nil {
			logg.Error("consumer stopped", zap.Error(err))
		}
	}()

	red := redis.InitRedis()
	_ = redis.Ping(red)
	if err != nil {
		fmt.Errorf("redis: связь отсутствует %w", err)
	}

	movieCache := cache.NewMovieCache(red)

	repo := repository.NewMovieRepo(postgresDB)
	serv := service.NewMovieService(repo, producer, movieCache)
	hand := handler.NewMovieHandler(serv, logg)

	router := router.NewRouter(hand)
	router.Use(logger.RequestIDMiddleware)
	router.Use(logger.LoggerMiddleware(logg))

	servUp := server.NewServer(port.Port, logg, router)

	go server.Shutdown(logg, servUp)

	logg.Info("Сервер запущен")

	if err := servUp.ListenAndServe(); err != nil &&
		err != http.ErrServerClosed {

		return fmt.Errorf(
			"ошибка запуска сервер %w",
			err,
		)
	}
	return nil
}
