package repo

import (
	"database/sql"
	"fmt"

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

func (pDb *pngDb) AddCategory(category string) error {
	//FIXME: sanitize input
	query := fmt.Sprintf(
		`CREATE TABLE %s(id INTEGER PRIMARY KEY NOT NULL, date INTEGER)`,
		category,
	)

	_, err := pDb.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

// TODO: add implementation
func (pDb *pngDb) NextProjectId(category string) (int, error) {
	return 0, nil
}

func (pDb *pngDb) Close() {
	pDb.db.Close()
}
