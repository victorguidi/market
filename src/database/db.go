package database

import (
	"database/sql"
	// "reflect"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	db    *sql.DB
	store Storage
}

type Query struct {
	Table  string
	Fields *any
}

func NewDatabase(file string) (*Database, error) {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Database{
		db: db,
	}, nil
}

func (d *Database) Init() error {
	return nil
}

func (d *Database) createTables() error {
	return nil
}

func (d *Database) Get(T any) (value, err error) {
	query := "SELECT u.id, u.username, u.password, us.session_id FROM users u INNER JOIN user_session us ON u.id = us.user_id"
	rows, err := d.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	err = rows.Scan(&T)

	return nil, nil
}

func (d *Database) GetOne(T any) (value, err error) {
	return nil, nil
}

func (d *Database) Insert(T any) error {
	return nil
}

func (d *Database) Update(T any) error {
	return nil
}

func (d *Database) Delete(T any) error {
	return nil
}
