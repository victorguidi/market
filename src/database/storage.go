package database

type Storage interface {
	Get() (value, err error)
	GetOne(T any) (value, err error)
	Insert(T any) error
	Update(T any) error
	Delete(T any) error
}
