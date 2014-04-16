package storage

import (
	"fmt"
	"testing"
)

func init() {
	storage = make(fakeStorage)
}

type fakeStorage map[string]interface{}

func (s fakeStorage) Get(key string) (interface{}, error) {
	value, ok := s[key]
	if !ok {
		return nil, fmt.Errorf("%v not found", key)
	}
	return value, nil
}

func (s fakeStorage) Put(key string, value interface{}) error {
	s[key] = value
	return nil
}

// TEST OMIT

func TestIncrement(t *testing.T) {
	count := 1
	if err := storage.Put("count", count); err != nil { // HL
		t.Fatalf("put: %v", err)
	}
	if err := Increment("count"); err != nil { // HL
		t.Fatalf("increment: %v", err)
	}
	v, err := storage.Get("count") // HL
	if err != nil {
		t.Fatalf("get: %v", err)
	}
	got, ok := v.(int)
	if !ok {
		t.Fatalf("wanted int; got %T", v)
	}
	if got != count+1 {
		t.Fatalf("wanted %v; got %v", got, count+1)
	}
}
