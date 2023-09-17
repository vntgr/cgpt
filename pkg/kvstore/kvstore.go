package kvstore

import (
	"fmt"

	"github.com/akrylysov/pogreb"
)

type KVInstance struct {
	DB *pogreb.DB
}

type KVOps interface {
	Put(key, val []byte) error
	Get(key []byte) (string, error)
}

func InitializeKVStore() (KVOps, error) {
	db, err := pogreb.Open("kily.store", nil)
	if err != nil {
		return nil, err
	}

	return &KVInstance{db}, nil
}

func (kv *KVInstance) Put(key, val []byte) error {
	fmt.Println(kv.Get(key))
	err := kv.DB.Put(key, val)
	if err != nil {
		return err
	}

	return nil
}

func (kv *KVInstance) Get(key []byte) (string, error) {
	val, err := kv.DB.Get(key)
	if err != nil {
		return "", err
	}

	return string(val), nil
}
