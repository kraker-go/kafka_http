package service

import (
	"context"
	"fmt"
	"kafka_http/internal/domain"
)

func (s *MovieService) GetMovies(ctx context.Context) ([]domain.Movie, error) {
	movies, err := s.repo.GetMovies(ctx)
	if err != nil {
		return []domain.Movie{}, fmt.Errorf("service: ошибка поиска %w", err)
	}

	return movies, nil
}
