-- +goose Up
CREATE TABLE movies (
                        id SERIAL PRIMARY KEY,
                        title TEXT,
                        year INT,
                        genre TEXT
);

-- +goose Down
DROP TABLE movies;