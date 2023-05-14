package database

// TODO: implement all the methods for user

import (
	"database/sql"
	// "reflect"

	_ "github.com/mattn/go-sqlite3"
)

type User struct{}

type Database struct {
	db    *sql.DB
	store DBStorage
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

func (d *Database) GetUsers() (value any, err error) {
	query := "SELECT u.id, u.username FROM users u INNER JOIN user_session us ON u.id = us.user_id"
	rows, err := d.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	type RUsers struct {
		ID       int
		Username string
	}

	response := make([]*RUsers, 0)
	for rows.Next() {
		resp := new(RUsers)
		err := rows.Scan(&resp.ID, &resp.Username)
		if err != nil {
			return nil, err
		}
		response = append(response, resp)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return response, nil
}

func (d *Database) GetOneUserById(id int64) (value any, err error) {
	return nil, nil
}

func (d *Database) InsertUser(user *User) error {
	return nil
}

func (d *Database) UpdateUser() error {
	return nil
}

func (d *Database) DeleteUser() error {
	return nil
}
