package service

import (
	"context"
	"fmt"
	"kafka_http/internal/domain"
)

func (s *MovieService) UpdateMovie(ctx context.Context, id int, movie *domain.Movie) (*domain.Movie, error) {
	if id <= 0 {
		return &domain.Movie{}, domain.ErrorIdIsEmpty
	}

	movie, err := s.repo.UpdateMovie(ctx, id, movie)
	if err != nil {
		return &domain.Movie{}, fmt.Errorf("service: ошибка обновления файла %w", err)
	}
	return movie, nil
}
