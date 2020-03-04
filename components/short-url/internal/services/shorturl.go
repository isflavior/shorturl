package services

import (
	"errors"

	"github.com/isflavior/shorturl/components/short-url/internal/stores"
)

type ShortURLService struct {
	Store        stores.Store
	KeyGenerator KeyGenerator
}

func (s *ShortURLService) GetLongURL(key string) string {
	longURL, err := s.Store.GetRecord(key)
	if err != nil {
		return ""
	}
	return longURL
}

func (s *ShortURLService) PutLongURL(url string) (string, error) {
	key, err := s.KeyGenerator.NewKey()
	if err != nil {
		return "", errors.New("key generation failed")
	}

	err = s.Store.PutRecord(key, url)
	if err != nil {
		return "", errors.New("record insert failed")
	}

	return key, err
}
