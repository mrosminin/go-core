// Package storage объявляет интерфейсы для драйверов механизмов хранения данных
// ну и здесь же модели данных
package storage

type Film struct {
	ID         int
	Title      string
	Year       int
	Fees       int
	StudioID   int
	StudioName string
	Rating     string
}

type Request struct {
	StudioID int
}

type Interface interface {
	GetFilms(Request) []Film
	NewFilms([]Film) error
	UpdateFilm(Film) error
	DeleteFilm(Film) error
}
