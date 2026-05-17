package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"kafka_http/internal/domain"
	"time"
)

type MovieCache struct {
	client *redis.Client
}

func NewMovieCache(client *redis.Client) *MovieCache {
	return &MovieCache{
		client: client,
	}
}

func (c *MovieCache) Set(ctx context.Context, movie *domain.Movie) error {

	data, err := json.Marshal(movie)
	if err != nil {
		return fmt.Errorf("json ошибка кодировки: %w", err)
	}

	key := fmt.Sprintf("movie:%d", movie.ID)

	return c.client.Set(
		ctx,
		key,
		data,
		time.Hour,
	).Err()
}

func (c *MovieCache) Get(ctx context.Context, id int) (*domain.Movie, error) {

	key := fmt.Sprintf("movie:%d", id)

	val, err := c.client.Get(ctx, key).Result()
	if err != nil {
		return nil, fmt.Errorf("Ошибка чтения кеша %w", err)
	}

	var movie domain.Movie

	if err = json.Unmarshal([]byte(val), &movie); err != nil {
		return nil, fmt.Errorf("json ошибка кодировки: %w", err)
	}

	return &movie, nil
}

func (c *MovieCache) Delete(ctx context.Context, id int) error {

	key := fmt.Sprintf("movie:%d", id)

	return c.client.Del(ctx, key).Err()
}
