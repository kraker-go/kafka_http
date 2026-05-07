package service

import (
	"context"
	"fmt"
	"kafka_http/internal/domain"
)

func (s *MovieService) CreateMovie(ctx context.Context, movie *domain.Movie) error {
	if movie.Title == "" {
		return domain.ErrorEmptyString
	}

	if movie.Year < 1900 || movie.Year > 2026 {
		return domain.ErrorAgeIncorrect
	}

	err := s.repo.CreateMovie(ctx, movie)
	if err != nil {

		return fmt.Errorf("service: ошибка создания фильма %w", err)
	}

	msg := []byte(fmt.Sprintf("Movie Created: %s", movie.Title))

	if err = s.producer.Send(ctx, msg); err != nil {
		return fmt.Errorf("service: kafka ошибка: %w", err)
	}

	return nil
}
