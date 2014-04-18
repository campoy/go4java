package storage

import "fmt"

type Storage interface {
	Get(key string) (interface{}, error)
	Put(key string, value interface{}) error
}

var storage Storage = newMemcacheStorage()

func newMemcacheStorage() Storage {
	return nil
}

func Increment(key string) error {
	v, err := storage.Get(key) // HL
	if err != nil {
		return fmt.Errorf("get: %v", err)
	}
	count, ok := v.(int)
	if !ok {
		return fmt.Errorf("expected int, got %T", v)
	}
	return storage.Put(key, count+1) // HL
}
