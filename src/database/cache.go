package database

type Cache struct {
	store Storage
}

func NewCache() (*Cache, error) {
	return &Cache{}, nil
}

func (c *Cache) Init() error {
	return nil
}

func (c *Cache) Get() (value, err error) {
	return nil, nil
}

func (c *Cache) GetOne(T any) (value, err error) {
	return nil, nil
}

func (c *Cache) Insert(T any) error {
	return nil
}

func (c *Cache) Update(T any) error {
	return nil
}

func (c *Cache) Delete(T any) error {
	return nil
}
