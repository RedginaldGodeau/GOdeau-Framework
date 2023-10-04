package Gobase

import (
	"GOdeau/godeau/Gobase/GoDriver"
	"database/sql"
	"errors"
	"fmt"
)

type Database struct {
	Conn *sql.DB
}

func AutoConnect() {}

func (db *Database) Close() error {
	return db.Conn.Close()
}

func (db *Database) Exec(sql string, params ...any) (sql.Result, error) {
	result, err := db.Conn.Exec(sql, params)
	closeErr := db.Close()

	if err != nil {
		return nil, err
	}

	if closeErr != nil {
		return nil, closeErr
	}

	return result, nil
}

func (db *Database) Rows(query string, buffer *struct{}, params ...interface{}) error {

	result := db.Conn.QueryRow(query, params)
	closeErr := db.Close()

	if closeErr != nil {
		return closeErr
	}

	err := result.Scan(&buffer)
	if !errors.Is(err, sql.ErrNoRows) {
		return err
	} else if err != nil {
		return err
	}

	return nil
}

func (db *Database) Query(sql string, tmp struct{}, buffer []*struct{}, params ...any) error {

	result, err := db.Conn.Query(sql, params)
	closeErr := db.Close()
	if err != nil {
		return err
	}

	if closeErr != nil {
		return closeErr
	}

	for result.Next() {
		err := result.Scan(&tmp)
		if err != nil {
			return err
		}

		buffer = append(buffer, &tmp)
	}

	return nil
}

func Connect(driver GoDriver.Driver, user string, password string, host string, database string) (*Database, error) {
	var url = fmt.Sprintf("%s://%s:%s@%s/%s?sslmode=disable", string(driver), user, password, host, database)

	db, err := sql.Open(string(driver), url)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Database{Conn: db}, nil
}
