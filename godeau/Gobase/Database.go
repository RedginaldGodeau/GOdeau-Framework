package Gobase

import (
	"GOdeau/godeau/Gobase/GoDriver"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	SQL *sql.DB
	ORM *gorm.DB
}

func (db *Database) Close() error {
	return db.SQL.Close()
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

	orm, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Database{SQL: db, ORM: orm}, nil
}
