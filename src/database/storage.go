package database

type CStorage interface {
	Get(key string) (value any, err error)
	GetOne(T any) (value, err error)
	Insert(data []byte, key string) error
	Update(T any) error
	Delete(T any) error
}

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Stocks   any    `json:"stocks"`
}

type Stock struct {
	Symbol string `json:"symbol"`
}

type DBStorage interface {
	GetUsers() (users []*User, err error)
	GetOneUserById(id int64) (value any, err error)
	InsertUser(user *User) error
	UpdateUser() error
	DeleteUser() error
	GetStocksFromUser(id int64) (value []Stock, err error)
	// InsertStockToUserById() error
	// DeleteStockFromUserById() error
	// GetStocks() (value any, err error)
	// GetOneStockBySymbol() (value any, err error)
	// InsertStock() error
	// UpdateStock() error
	// DeleteStock() error
}
