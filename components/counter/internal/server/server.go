package server

import (
	"io/ioutil"
	"net/http"
	"strconv"
)

// Server the http handler
type Server struct {
	FilePath string
	Counter  int
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Counter++
	current := strconv.Itoa(s.Counter)
	uid := []byte(current)

	err := ioutil.WriteFile(s.FilePath, uid, 0644)

	if err != nil {
		w.WriteHeader(400)
		return
	}

	w.Write(uid)
}
