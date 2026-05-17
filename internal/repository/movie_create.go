package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"kafka_http/internal/domain"
)

func (m *MovieRepo) CreateMovie(ctx context.Context, movie *domain.Movie) error {
	tx, err := m.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("repository: ошибка начала транзакции при создании фильма: %w", err)
	}
	defer tx.Rollback()

	var check domain.Movie

	err = tx.QueryRowContext(ctx, get_movie_title, movie.Title).Scan(&check.ID, &check.Title, &check.Year, &check.Genre)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = tx.QueryRowContext(ctx, create, movie.Title, movie.Year, movie.Genre).Scan(&movie.ID)
			if err != nil {
				return fmt.Errorf("repository: ошибка создания фильма: %w", err)
			}

			return nil
		}
		return fmt.Errorf("repository: ошибка запросы в бд: %w", err)
	}

	return domain.ErrorMovieExist
}
