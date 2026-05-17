package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"kafka_http/internal/domain"
)

func (m *MovieRepo) GetMovieByID(ctx context.Context, id int) (domain.Movie, error) {
	var movie domain.Movie

	err := m.db.QueryRowContext(ctx, get_movie_id, id).Scan(&movie.ID, &movie.Title, &movie.Year, &movie.Genre)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Movie{}, domain.ErrorMovieNotFound
		}

		return domain.Movie{}, fmt.Errorf("repository: ошибка получения фильма %w", err)
	}

	return movie, nil
}
