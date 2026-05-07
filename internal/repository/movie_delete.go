package repository

import (
	"context"
	"database/sql"
	"fmt"
	"kafka_http/internal/domain"
)

func (m *MovieRepo) DeleteMovie(ctx context.Context, id int) error {
	var check domain.Movie
	err := m.db.QueryRowContext(ctx, delete, id).Scan(&check.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.ErrorMovieNotFound
		}
		return fmt.Errorf("ошибка запроса в бд %w", err)
	}

	result, err := m.db.ExecContext(ctx, delete, id)
	if err != nil {
		return fmt.Errorf("repository: ошибка удаления фильма %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("repository: ошибка проверки результата %w", err)
	}
	if rowsAffected == 0 {
		return domain.ErrorMovieNotFound
	}

	return nil
}
