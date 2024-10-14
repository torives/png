package repo

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

type PngDatabase interface {
	AddCategory(category string) error
	NextProjectId(category string) (int, error)
	Close()
}

type pngDb struct {
	db *sql.DB
}

func Open(path string) (PngDatabase, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	return &pngDb{db}, nil
}

// TODO: add implementation
func (pDb *pngDb) AddCategory(category string) error {
	return nil
}

// TODO: add implementation
func (pDb *pngDb) NextProjectId(category string) (int, error) {
	return 0, nil
}

func (pDb *pngDb) Close() {
	pDb.db.Close()
}
