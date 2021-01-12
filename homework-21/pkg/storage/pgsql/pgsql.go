package pgsql

import (
	"context"
	"github.com/jackc/pgx/pgxpool"
	"go-core-own/homework-21/pkg/storage"
)

// Storage - служба работы с БД
type Storage struct {
	db *pgxpool.Pool
}

// New - фабрика, создаёт экземпляр хранилища
func New(connStr string) (*Storage, error) {
	s := new(Storage)
	var err error
	s.db, err = pgxpool.Connect(context.Background(), connStr)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// GetFilms возвращает фильмы
func (s *Storage) GetFilms(req storage.Request) (data []storage.Film, err error) {
	rows, err := s.db.Query(context.Background(), `
	SELECT
		films.id,
		films.title,
		films.year,
		films.fees,
		films.studio_id,
		studios.name,
		films.rating
	FROM
		films
	LEFT JOIN
		studios ON films.studio_id = studios.id
	WHERE $1 = 0 OR films.studio_id = $1;
	`,
		req.StudioID,
	)
	if err != nil {
		return data, err
	}
	defer rows.Close()

	for rows.Next() {
		var film storage.Film

		err = rows.Scan(
			&film.ID,
			&film.Title,
			&film.Year,
			&film.Fees,
			&film.StudioID,
			&film.StudioName,
			&film.Rating,
		)
		if err != nil {
			return data, err
		}

		data = append(data, film)
	}
	if rows.Err() != nil {
		return data, err
	}

	return data, nil
}

// NewFilms записывает новые фильмы
func (s *Storage) NewFilms(films []storage.Film) (err error) {
	tx, err := s.db.Begin(context.Background())
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background())

	for _, film := range films {
		_, err = tx.Exec(context.Background(), `
			INSERT INTO
				films (
					title,
					year,
					fees,
					studio_id,
					rating
				)
			VALUES ($1, $2, $3, $4, $5) RETURNING id;
			`,
			film.Title,
			film.Year,
			film.Fees,
			film.StudioID,
			film.Rating,
		)
		if err != nil {
			return err
		}
	}
	err = tx.Commit(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// UpdateFilm обновляет фильм
func (s *Storage) UpdateFilm(film storage.Film) (err error) {
	_, err = s.db.Exec(context.Background(), `
			UPDATE
				films
			SET
				title = $2,
				year = $3,
				fees = $4,
				studio_id = $5,
				rating = $6
			WHERE
				id = $1
			`,
		film.ID,
		film.Title,
		film.Year,
		film.Fees,
		film.StudioID,
		film.Rating,
	)

	return err
}

// DeleteFilm удаляет фильм из БД
func (s *Storage) DeleteFilm(film storage.Film) (err error) {
	_, err = s.db.Exec(context.Background(), `DELETE FROM films WHERE id = $1`, film.ID)

	return err
}
