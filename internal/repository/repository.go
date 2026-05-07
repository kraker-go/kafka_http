package repository

import (
	"database/sql"
)

type MovieRepo struct {
	db *sql.DB
}

func NewMovieRepo(db *sql.DB) *MovieRepo {
	return &MovieRepo{db: db}
}

// константы запросов
const create = "INSERT INTO movies (title, year, genre) VALUES ($1, $2, $3) returning id"
const get_movie_title = "SELECT id, title, year, genre FROM movies WHERE title = $1"
const get_all = "SELECT id, title, year, genre FROM movies"
const get_movie_id = "SELECT id, title, year, genre FROM movies WHERE id = $1"
const delete = "DELETE FROM movies WHERE id = $1"
const update = "UPDATE movies SET title=$1, year=$2, genre=$3 WHERE id=$4 returning id, title, year, genre"
