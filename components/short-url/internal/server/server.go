package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/isflavior/shorturl/components/short-url/internal/services"
)

// Server the http handler
type Server struct {
	Host            string
	ShortURLService services.ShortURLService
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		key := strings.TrimPrefix(r.URL.Path, "/")
		location := s.ShortURLService.GetLongURL(key)

		if location == "" {
			w.WriteHeader(404)
			return
		}

		w.Header().Set("Location", location)
		w.WriteHeader(302)
		return
	case "POST":
		body, err := ioutil.ReadAll(r.Body)

		longURL, err := url.ParseRequestURI(string(body))
		if err != nil {
			w.WriteHeader(400)
			return
		}

		shortID, err := s.ShortURLService.PutLongURL(longURL.String())
		if err != nil {
			w.WriteHeader(400)
			return
		}

		shortURL := fmt.Sprintf("http://%s/%s", s.Host, shortID)

		w.Write([]byte(shortURL))
	}
}
