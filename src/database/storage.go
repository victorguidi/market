package database

type CStorage interface {
	Get() (value, err error)
	GetOne(T any) (value, err error)
	Insert(T any) error
	Update(T any) error
	Delete(T any) error
}

type DBStorage interface {
	GetUsers() (value any, err error)
	GetOneUserById() (value any, err error)
	InsertUser() error
	UpdateUser() error
	DeleteUser() error
}
