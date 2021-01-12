package pgsql

import (
	"context"
	"go-core-own/homework-21/pkg/storage"
	"log"
	"os"
	"testing"
)

var s *Storage

func TestMain(m *testing.M) {
	var err error
	// БД для тетов.
	s, err = New("postgres://postgres:qwerty123@localhost/films")
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer s.db.Close()
	os.Exit(m.Run())
}

func TestStorage_CRUD(t *testing.T) {
	_, err := s.db.Exec(context.Background(), `DELETE FROM films WHERE id > 0`)
	if err != nil {
		t.Fatalf("Unable to truncate table films: %v\n", err)
	}
	// Create test
	films := []storage.Film{
		{
			Title:    "Film1",
			Year:     2020,
			Fees:     1_000_000,
			StudioID: 1,
			Rating:   "PG-18",
		},
		{
			Title:    "Film2",
			Year:     2021,
			Fees:     500_000,
			StudioID: 2,
			Rating:   "PG-13",
		},
	}
	err = s.NewFilms(films)
	if err != nil {
		t.Fatalf("NewFilms() error: %v\n", err)
	}

	// Read tests
	tests := []struct {
		name    string
		req     storage.Request
		wantLen int
	}{
		{
			name:    "ReadTest1",
			req:     storage.Request{},
			wantLen: 3,
		},
		{
			name:    "ReadTest2",
			req:     storage.Request{StudioID: 1},
			wantLen: 1,
		},
		{
			name:    "ReadTest3",
			req:     storage.Request{StudioID: 3},
			wantLen: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := s.GetFilms(tt.req)
			if err != nil {
				t.Errorf("GetFilms() error: %v\n", err)
			}
			if gotLen := len(data); tt.wantLen != gotLen {
				t.Errorf("GetFilms() error: %v\n", err)
			}
		})
	}

	// Update test
	data, err := s.GetFilms(storage.Request{})
	if err != nil {
		t.Fatalf("GetFilms() error: %v\n", err)
	}
	if len(data) < 2 {
		t.Fatalf("No new films: cannot proceed with tests")
	}

	film := data[len(data)-1]
	film.Title = "New title"

	err = s.UpdateFilm(film)
	if err != nil {
		t.Errorf("UpdateFilm() error: %v\n", err)
	}
	data, err = s.GetFilms(storage.Request{})
	if err != nil {
		t.Fatalf("GetFilms() error: %v\n", err)
	}

	var updatedFilm storage.Film
	for _, item := range data {
		if item.ID == film.ID {
			updatedFilm = item
		}
	}
	if updatedFilm.Title != film.Title {
		t.Errorf("UpdateFilm() error: wanted title = '%s', got = '%s'\n", film.Title, updatedFilm.Title)
	}

	// Delete test
	err = s.DeleteFilm(film)
	if err != nil {
		t.Errorf("DeleteFilm() error: %v\n", err)
	}
	data, err = s.GetFilms(storage.Request{})
	if err != nil {
		t.Fatalf("GetFilms() error: %v\n", err)
	}
	found := false
	for _, item := range data {
		if item.ID == film.ID {
			found = true
		}
	}
	if found {
		t.Errorf("DeleteFilm() error: film wasn't deleted")
	}

}
