package Database

import (
	sql2 "database/sql"
	"errors"
)

type Database interface {
	Connection() error
	Close() error
	Exec(sql string, param ...any) (sql2.Result, error)
	QueryRow(sql string, data *struct{}, param ...any) error
	Query(sql string, buffer *struct{}, data []*struct{}, param ...any) error
}

type Db struct {
	Driver string
	Source string
	Conn   *sql2.DB

	Database
}

// type parameters interface{} @! Pas important pour l'instant

func (db *Db) Connection() error {

	conn, err := sql2.Open(db.Driver, db.Source)

	if err != nil {
		return err
	}

	err = conn.Ping()
	if err != nil {
		return err
	}

	db.Conn = conn

	return nil
}

func (db *Db) Close() error {
	err := db.Conn.Close()
	return err
}

func (db *Db) Exec(sql string, param ...any) (sql2.Result, error) {
	if err := db.Connection(); err != nil {
		return nil, err
	}

	result, err := db.Conn.Exec(sql, param)
	closeErr := db.Close()

	if err != nil {
		return nil, err
	}

	if closeErr != nil {
		return nil, closeErr
	}

	return result, nil
}

func (db *Db) QueryRow(sql string, data *struct{}, param ...any) error {
	if err := db.Connection(); err != nil {
		return err
	}

	result := db.Conn.QueryRow(sql, param)
	closeErr := db.Close()

	if closeErr != nil {
		return closeErr
	}

	err := result.Scan(&data)
	if !errors.Is(err, sql2.ErrNoRows) {
		return err
	} else if err != nil {
		return err
	}

	return nil
}

func (db *Db) Query(sql string, buffer *struct{}, data []*struct{}, param ...any) error {
	if err := db.Connection(); err != nil {
		return err
	}

	result, err := db.Conn.Query(sql, param)
	closeErr := db.Close()
	if err != nil {
		return err
	}

	if closeErr != nil {
		return closeErr
	}

	for result.Next() {
		err := result.Scan(&buffer)
		if err != nil {
			return err
		}

		data = append(data, buffer)
	}

	return nil
}
