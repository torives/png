package repo

import (
	"database/sql"
	"fmt"
	"time"

	_ "modernc.org/sqlite"
)

type PngDatabase interface {
	AddCategory(category string) error
	NextProjectId(category string) (int64, error)
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

func (pDb *pngDb) NextProjectId(category string) (int64, error) {
	//FIXME: sanitize input
	query := fmt.Sprintf(`INSERT INTO %s VALUES (null, %d)`, category, time.Now().Unix())
	result, err := pDb.db.Exec(query)

	if err != nil {
		return -1, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}

	return id, nil
}

func (pDb *pngDb) Close() {
	pDb.db.Close()
}
