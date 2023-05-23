package database

// TODO: implement all the methods for user

import (
	"database/sql"
	"log"
	"time"

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
	queryStocks := `CREATE TABLE IF NOT EXISTS stocks (id INTEGER PRIMARY KEY AUTOINCREMENT, Symbol VARCHAR(255), Name VARCHAR(255), Description TEXT, CIK VARCHAR(255), Exchange VARCHAR(255), Currency VARCHAR(255), Country VARCHAR(255), Sector VARCHAR(255), Industry VARCHAR(255), Address VARCHAR(255), FiscalYearEnd VARCHAR(255), LatestQuarter VARCHAR(255), MarketCapitalization VARCHAR(255), EBITDA VARCHAR(255), PERatio VARCHAR(255), PEGRatio VARCHAR(255), BookValue VARCHAR(255), DividendPerShare VARCHAR(255), DividendYield VARCHAR(255), EPS VARCHAR(255), RevenuePerShareTTM VARCHAR(255), ProfitMargin VARCHAR(255), OperatingMarginTTM VARCHAR(255), ReturnOnAssetsTTM VARCHAR(255), ReturnOnEquityTTM VARCHAR(255), RevenueTTM VARCHAR(255), GrossProfitTTM VARCHAR(255), DilutedEPSTTM VARCHAR(255), QuarterlyEarningsGrowthYOY VARCHAR(255), QuarterlyRevenueGrowthYOY VARCHAR(255), AnalystTargetPrice VARCHAR(255), TrailingPE VARCHAR(255), ForwardPE VARCHAR(255), PriceToSalesRatioTTM VARCHAR(255), PriceToBookRatio VARCHAR(255), EVToRevenue VARCHAR(255), EVToEBITDA VARCHAR(255), Beta VARCHAR(255), SharesOutstanding VARCHAR(255), DividendDate VARCHAR(255), ExDividendDate VARCHAR(255), updated_at DATETIME DEFAULT CURRENT_TIMESTAMP)`
	queryRss := `CREATE TABLE IF NOT EXISTS feed (id INTEGER PRIMARY KEY AUTOINCREMENT, link VARCHAR(255), created_at DATETIME DEFAULT CURRENT_TIMESTAMP)`
	queryRssUser := `CREATE TABLE IF NOT EXISTS user_feed (user_id INT NOT NULL, feed_id INT NOT NULL)`

	if _, err := d.db.Exec(queryUser); err != nil {
		return err
	}

	if _, err := d.db.Exec(queryUserStocks); err != nil {
		return err
	}

	if _, err := d.db.Exec(queryStocks); err != nil {
		return err
	}

	if _, err := d.db.Exec(queryRss); err != nil {
		return err
	}

	if _, err := d.db.Exec(queryRssUser); err != nil {
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

func (d *Database) UpdateUser() error {
	return nil
}

func (d *Database) DeleteUser() error {
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
	log.Println(stocks)
	return stocks, nil
}

func (d *Database) InsertStockToUser(id int64, stock *Stock) error {
	query := "INSERT INTO user_stocks (user_id, stock_id, symbol) VALUES (?, ?, ?)"
	stmt, err := d.db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id, stock.ID, stock.Symbol)
	if err != nil {
		return err
	}
	return nil
}

func (d *Database) InsertNewStockInfo(stock *Stock) error {
	query := "INSERT INTO stocks (Symbol) VALUES (?)"
	stmt, err := d.db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(stock.Symbol)
	if err != nil {
		return err
	}
	return nil
}

func (d *Database) UpdateStockInfo(symbol string, stock *Stock) error {
	query := "UPDATE stocks SET Name = ?, Description = ?, CIK = ?, Exchange = ?, Currency = ?, Country = ?, Sector = ?, Industry = ?, Address = ?, FiscalYearEnd = ?, LatestQuarter = ?, MarketCapitalization = ?, EBITDA = ?, PERatio = ?, PEGRatio = ?, BookValue = ?, DividendPerShare = ?, DividendYield = ?, EPS = ?, RevenuePerShareTTM = ?, ProfitMargin = ?, OperatingMarginTTM = ?, ReturnOnAssetsTTM = ?, ReturnOnEquityTTM = ?, RevenueTTM = ?, GrossProfitTTM = ?, DilutedEPSTTM = ?, QuarterlyEarningsGrowthYOY = ?, QuarterlyRevenueGrowthYOY = ?, AnalystTargetPrice = ?, TrailingPE = ?, ForwardPE = ?, PriceToSalesRatioTTM = ?, PriceToBookRatio = ?, EVToRevenue = ?, EVToEBITDA = ?, Beta = ?, SharesOutstanding = ?, DividendDate = ?, ExDividendDate = ?, updated_at = ? WHERE Symbol = ?"
	stmt, err := d.db.Prepare(query)
	if err != nil {
		return err
	}

	current_time := time.Now().UTC()

	_, err = stmt.Exec(stock.Name, stock.Description, stock.CIK, stock.Exchange, stock.Currency, stock.Country, stock.Sector, stock.Industry, stock.Address, stock.FiscalYearEnd, stock.LatestQuarter, stock.MarketCapitalization, stock.EBITDA, stock.PERatio, stock.PEGRatio, stock.BookValue, stock.DividendPerShare, stock.DividendYield, stock.EPS, stock.RevenuePerShareTTM, stock.ProfitMargin, stock.OperatingMarginTTM, stock.ReturnOnAssetsTTM, stock.ReturnOnEquityTTM, stock.RevenueTTM, stock.GrossProfitTTM, stock.DilutedEPSTTM, stock.QuarterlyEarningsGrowthYOY, stock.QuarterlyRevenueGrowthYOY, stock.AnalystTargetPrice, stock.TrailingPE, stock.ForwardPE, stock.PriceToSalesRatioTTM, stock.PriceToBookRatio, stock.EVToRevenue, stock.EVToEBITDA, stock.Beta, stock.SharesOutstanding, stock.DividendDate, stock.ExDividendDate, current_time, symbol)
	if err != nil {
		return err
	}
	return nil
}

func (d *Database) GetStockAndCheckLastUpdate(symbol string) (bool, error) {
	query := "SELECT updated_at FROM stocks WHERE Symbol = ?"
	stmt, err := d.db.Prepare(query)
	if err != nil {
		return false, err
	}
	var updatedAt time.Time
	err = stmt.QueryRow(symbol).Scan(&updatedAt)
	if err != nil {
		return false, err
	}
	log.Println(updatedAt)
	if time.Since(updatedAt).Hours() > 24 {
		return true, nil
	}
	return false, nil
}

func (d *Database) GetStockInfo(symbol string) (value *Stock, err error) {
	query := "SELECT * FROM stocks WHERE Symbol = ?"
	stmt, err := d.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	var stock Stock
	err = stmt.QueryRow(symbol).Scan(&stock.ID, &stock.Symbol, &stock.Name, &stock.Description, &stock.CIK, &stock.Exchange, &stock.Currency, &stock.Country, &stock.Sector, &stock.Industry, &stock.Address, &stock.FiscalYearEnd, &stock.LatestQuarter, &stock.MarketCapitalization, &stock.EBITDA, &stock.PERatio, &stock.PEGRatio, &stock.BookValue, &stock.DividendPerShare, &stock.DividendYield, &stock.EPS, &stock.RevenuePerShareTTM, &stock.ProfitMargin, &stock.OperatingMarginTTM, &stock.ReturnOnAssetsTTM, &stock.ReturnOnEquityTTM, &stock.RevenueTTM, &stock.GrossProfitTTM, &stock.DilutedEPSTTM, &stock.QuarterlyEarningsGrowthYOY, &stock.QuarterlyRevenueGrowthYOY, &stock.AnalystTargetPrice, &stock.TrailingPE, &stock.ForwardPE, &stock.PriceToSalesRatioTTM, &stock.PriceToBookRatio, &stock.EVToRevenue, &stock.EVToEBITDA, &stock.Beta, &stock.SharesOutstanding, &stock.DividendDate, &stock.ExDividendDate, &stock.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &stock, nil
}

func (d *Database) InsertNewLinkRss(link string) error {
	query := "INSERT INTO rss SELECT ? WHERE NOT EXISTS (SELECT * FROM rss WHERE link = ?)"
	stmt, err := d.db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(link, link)
	if err != nil {
		return err
	}
	return nil
}

func (d *Database) GetAllLinkRss() (value []string, err error) {
	query := "SELECT link FROM rss"
	rows, err := d.db.Query(query)
	if err != nil {
		return nil, err
	}
	var links []string
	for rows.Next() {
		var link string
		err = rows.Scan(&link)
		if err != nil {
			return nil, err
		}
		links = append(links, link)
	}
	return links, nil
}

func (d *Database) DeleteLinkRss(link string) error {
	return nil
}
func (d *Database) DeleteStockFromUserById() error {
	return nil
}
func (d *Database) DeleteStock() error {
	return nil
}
