package stores

import (
	"github.com/syndtr/goleveldb/leveldb"
)

type LevelDBStore struct {
	DB *leveldb.DB
}

func (s *LevelDBStore) GetRecord(key string) (string, error) {
	data, err := s.DB.Get([]byte(key), nil)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (s *LevelDBStore) PutRecord(key string, value string) error {
	err := s.DB.Put([]byte(key), []byte(value), nil)
	return err
}
