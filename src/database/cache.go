package database

import (
	"log"
	"os"
)

// TODO: implement all the cache

// The way the cache will work is: For Hourly, Daily, Weekly, Monthly timeframes, we will have a cache for each one of them.
// The cache will save a file for each timeframe, and it will store the data in bytes. When the user requests a timeframe, the Cache will check if the file exists, if it does, it will return the data from the file.
// If the file doesn't exist, it will call the api, get the data, save it in the file, and return the data to the user.
// The cache will have a method to check if the file exists, and a method to save the data in the file.
// The cache will have a method to get the data from the file, and a method to delete the file.
// Every 12 hours, the cache will delete the files, and it will only create a new one when the user requests the data.

type Cache struct {
	data  map[string][]byte
	store CStorage
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string][]byte),
	}
}

func (c *Cache) Init() error {
	return c.createFile()
}

func (c *Cache) Get(key string) (value any, err error) {
	file, err := os.Open("/run/media/victorguidi/Projects/market/src/databases/cache")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer file.Close()

	clen := len(c.data[key])
	buf := make([]byte, clen)
	_, err = file.Read(buf)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return buf, nil
}

func (c *Cache) GetOne(T any) (value, err error) {
	return nil, nil
}

func (c *Cache) Insert(data []byte) error {
	file, err := os.OpenFile("/run/media/victorguidi/Projects/market/src/databases/cache", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer file.Close()

	c.data["daily"] = data

	_, err = file.Write(c.data["daily"])
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (c *Cache) Update(T any) error {
	return nil
}

func (c *Cache) Delete(T any) error {
	return nil
}

func (c *Cache) createFile() error {

	cacheFile, err := os.Create("/run/media/victorguidi/Projects/market/src/databases/cache")
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer cacheFile.Close()

	return nil

}

func (c *Cache) DeleteFile() error {
	return nil
}
