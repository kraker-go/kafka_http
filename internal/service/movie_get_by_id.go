package service

import (
	"context"
	"fmt"
	"kafka_http/internal/domain"
)

func (s *MovieService) GetMovieByID(ctx context.Context, id int) (*domain.Movie, error) {

	movieCache, err := s.c.Get(ctx, id)
	if err == nil {
		return movieCache, nil
	}

	movieRepo, err := s.repo.GetMovieByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("service: get movie repo: %w", err)
	}

	if err = s.c.Set(ctx, &movieRepo); err != nil {
		return nil, fmt.Errorf("service: cache set: %w", err)
	}

	return &movieRepo, nil
}
