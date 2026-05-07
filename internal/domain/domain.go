package domain

import "errors"

// ошибки
var (
	ErrorEmptyString   = errors.New("название не заполнено")
	ErrorAgeIncorrect  = errors.New("год выпуска слишком старый")
	ErrorMovieExist    = errors.New("такой фильм уже есть в базе")
	ErrorMovieNotFound = errors.New("фильм с таким ID не найден")
	ErrorIdIsEmpty     = errors.New("ID должен быть больше 0")
)

type Movie struct {
	ID    int
	Title string
	Year  int
	Genre string
}
