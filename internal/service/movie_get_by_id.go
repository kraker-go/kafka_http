package service

import (
	"context"
	"fmt"
	"kafka_http/internal/domain"
)

func (s *MovieService) GetMovieByID(ctx context.Context, id int) (domain.Movie, error) {
	if id <= 0 {
		return domain.Movie{}, domain.ErrorIdIsEmpty
	}
	movie, err := s.repo.GetMovieByID(ctx, id)
	if err != nil {
		return domain.Movie{}, fmt.Errorf("service: ошибка поиска %w", err)
	}

	return movie, nil
}
