package database

// TODO: implement all the methods for user

import (
	"database/sql"
	// "reflect"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	db *sql.DB
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

	if err := d.createTables(); err != nil {
		return err
	}

	return nil

}

func (d *Database) createTables() error {
	queryUser := `CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY AUTOINCREMENT, username VARCHAR(255), password TEXT NOT NULL, email VARCHAR(255), created_at DATETIME DEFAULT CURRENT_TIMESTAMP, updated_at DATETIME DEFAULT CURRENT_TIMESTAMP)`
	queryUserStocks := `CREATE TABLE IF NOT EXISTS user_stocks (user_id INT NOT NULL, stock_id INT NOT NULL, symbol VARCHAR(255))`

	if _, err := d.db.Exec(queryUser); err != nil {
		return err
	}

	if _, err := d.db.Exec(queryUserStocks); err != nil {
		return err
	}

	return nil
}

func (d *Database) GetUsers() (users []*User, err error) {

	query := "SELECT u.id, u.username FROM users u"
	rows, err := d.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	response := make([]*User, 0)
	for rows.Next() {
		resp := new(User)
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

	query := "INSERT INTO users (username, password, email) VALUES (?, ?, ?)"

	stmt, err := d.db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(user.Username, user.Password, user.Email)
	if err != nil {
		return err
	}

	return nil

}

func (d *Database) GetStocksFromUser(id int64) (value []Stock, err error) {
	query := "SELECT s.symbol FROM user_stocks s INNER JOIN users u ON s.user_id = u.id WHERE u.id = ?"
	stmt, err := d.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	var stocks []Stock
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var stock Stock
		err := rows.Scan(&stock.Symbol)
		if err != nil {
			return nil, err
		}
		stocks = append(stocks, stock)
	}
	return stocks, nil
}

func (d *Database) UpdateUser() error {
	return nil
}

func (d *Database) DeleteUser() error {
	return nil
}
