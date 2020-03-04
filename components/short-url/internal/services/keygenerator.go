package services

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/catinello/base62"
)

type KeyGenerator struct {
	Endpoint string
}

func (k *KeyGenerator) NewKey() (string, error) {
	response, err := http.Get(k.Endpoint)
	if err != nil {
		return "", errors.New("uid request failed")
	}

	data, _ := ioutil.ReadAll(response.Body)
	uid, _ := strconv.Atoi(string(data))

	key := base62.Encode(uid)
	return key, nil
}
