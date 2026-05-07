package service

import (
	"context"
	"fmt"
	"kafka_http/internal/domain"
)

func (s *MovieService) DeleteMovie(ctx context.Context, id int) error {
	if id <= 0 {
		return domain.ErrorIdIsEmpty
	}
	err := s.repo.DeleteMovie(ctx, id)
	if err != nil {
		return fmt.Errorf("service: ошибка удаления файла %w", err)

	}

	return nil

}
