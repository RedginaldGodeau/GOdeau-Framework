package Gobase

import (
	"database/sql"
	"errors"
)

func (db *Database) Exec(sql string, params ...any) (sql.Result, error) {
	result, err := db.SQL.Exec(sql, params)
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

	result := db.SQL.QueryRow(query, params)
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

	result, err := db.SQL.Query(sql, params)
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
