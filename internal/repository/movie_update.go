package repository

import (
	"context"
	"database/sql"
	"fmt"
	"kafka_http/internal/domain"
)

func (m *MovieRepo) UpdateMovie(ctx context.Context, ID int, movie *domain.Movie) (*domain.Movie, error) {
	tx, err := m.db.BeginTx(ctx, nil)
	if err != nil {
		return &domain.Movie{}, fmt.Errorf("repository: ошибка начала транзакции при обновлении %w", err)
	}
	defer tx.Rollback()

	var check domain.Movie

	err = tx.QueryRowContext(ctx, get_movie_id, ID).Scan(&check.ID, &check.Title, &check.Year, &check.Genre)
	if err != nil {
		if err == sql.ErrNoRows {
			return &domain.Movie{}, domain.ErrorMovieNotFound
		}
		return &domain.Movie{}, fmt.Errorf("repository: ошибка проверки обновления данных %w", err)
	}

	err = tx.QueryRowContext(ctx, update, movie.Title, movie.Year, movie.Genre, ID).Scan(&check.ID, &check.Title, &check.Year, &check.Genre)
	if err != nil {
		return &domain.Movie{}, fmt.Errorf("repository: ошибка обновления данных %w", err)
	}

	tx.Commit()

	return &check, nil
}
