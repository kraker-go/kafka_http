package repository

import (
	"context"
	"fmt"
	"kafka_http/internal/domain"
)

func (m *MovieRepo) GetMovies(ctx context.Context) ([]domain.Movie, error) {
	var movies []domain.Movie

	rows, err := m.db.QueryContext(ctx, get_all)
	if err != nil {
		return nil, fmt.Errorf("repository: ошибка запроса поиска всех фильмов %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var movie domain.Movie
		err := rows.Scan(&movie.ID, &movie.Title, &movie.Year, &movie.Genre)
		if err != nil {
			return nil, fmt.Errorf("repository: ошибка получения всех фильмов %w", err)
		}
		movies = append(movies, movie)
	}

	return movies, nil
}
