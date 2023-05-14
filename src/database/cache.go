package database

// TODO: implement all the cache

// The way the cache will work is: For Hourly, Daily, Weekly, Monthly timeframes, we will have a cache for each one of them.
// The cache will save a file for each timeframe, and it will store the data in bytes. When the user requests a timeframe, the Cache will check if the file exists, if it does, it will return the data from the file.
// If the file doesn't exist, it will call the api, get the data, save it in the file, and return the data to the user.
// The cache will have a method to check if the file exists, and a method to save the data in the file.
// The cache will have a method to get the data from the file, and a method to delete the file.
// Every 12 hours, the cache will delete the files, and it will only create a new one when the user requests the data.

type Cache struct {
	store CStorage
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
