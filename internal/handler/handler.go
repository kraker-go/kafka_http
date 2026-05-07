package handler

import (
	"go.uber.org/zap"
	"kafka_http/internal/service"
)

type MovieHandler struct {
	Hand *service.MovieService
	Logg *zap.Logger
}

func NewMovieHandler(s *service.MovieService, logg *zap.Logger) MovieHandler {
	return MovieHandler{s, logg}
}
