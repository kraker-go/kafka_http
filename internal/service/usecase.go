package service

import (
	"context"
	"kafka_http/internal/cache"
	"kafka_http/internal/domain"
	"kafka_http/internal/kafka"
)

type MovieCRUD interface {
	CreateMovie(context.Context, *domain.Movie) error
	GetMovies(context.Context) ([]domain.Movie, error)
	GetMovieByID(context.Context, int) (domain.Movie, error)
	DeleteMovie(context.Context, int) error
	UpdateMovie(context.Context, int, *domain.Movie) (*domain.Movie, error)
}

type MovieService struct {
	repo     MovieCRUD
	producer *kafka.KafkaProducer
	c        *cache.MovieCache
}

func NewMovieService(repo MovieCRUD, prod *kafka.KafkaProducer, cache *cache.MovieCache) *MovieService {
	return &MovieService{repo: repo, producer: prod, c: cache}
}
