package storage

import (
	"sync"
)

// Storage constraints.
const (
	KeyMaxBytes   = 16
	ValueMaxBytes = 512

	StorageMaxSize = 1024
)

// Storage errors.
const (
	ErrKeyNotFound  = Error("key not found")
	ErrKeyTooLong   = Error("key too long")
	ErrKeyEmpty     = Error("key is empty")
	ErrValueTooLong = Error("value too long")
	ErrStorageFull  = Error("storage is full")
)

type Storage interface {
	Get(key string) (string, error)
	Set(key, val string) error
	Delete(key string) error
	Exists(key string) (bool, error)
}

// New returns new instance of the Storage.
func New() Storage {
	return &inMemmory{
		mux:  sync.Mutex{},
		data: make(map[string]string),
	}
}

type inMemmory struct {
	mux  sync.Mutex
	data map[string]string
}

func (im *inMemmory) Get(key string) (string, error) {
	if err := check(key); err != nil {
		return "", err
	}

	im.mux.Lock()
	defer im.mux.Unlock()

	val, ok := im.data[key]
	if !ok {
		return "", ErrKeyNotFound
	}

	return val, nil
}

func (im *inMemmory) Set(key, val string) error {
	if err := check(key); err != nil {
		return err
	}

	if err := valueCheck(val); err != nil {
		return err
	}

	im.mux.Lock()
	defer im.mux.Unlock()

	if im.isFull() && !im.exists(key) {
		return ErrStorageFull
	}

	im.data[key] = val

	return nil
}

func (im *inMemmory) Exists(key string) (bool, error) {
	if err := check(key); err != nil {
		return false, err
	}

	im.mux.Lock()
	defer im.mux.Unlock()

	return im.exists(key), nil
}

func (im *inMemmory) Delete(key string) error {
	if err := check(key); err != nil {
		return err
	}

	im.mux.Lock()
	defer im.mux.Unlock()

	if !im.exists(key) {
		return ErrKeyNotFound
	}

	delete(im.data, key)
	return nil
}

func (im *inMemmory) exists(key string) bool {
	_, ok := im.data[key]
	return ok
}

func (im *inMemmory) isFull() bool {
	return len(im.data) > StorageMaxSize
}

func check(key string) error {
	if len([]byte(key)) > KeyMaxBytes {
		return ErrKeyTooLong
	}

	if key == "" {
		return ErrKeyEmpty
	}
	return nil
}

func valueCheck(val string) error {
	if len(val) > ValueMaxBytes {
		return ErrValueTooLong
	}

	return nil
}
