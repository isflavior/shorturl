package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/isflavior/shorturl/components/short-url/internal/server"
	"github.com/isflavior/shorturl/components/short-url/internal/services"
	stores "github.com/isflavior/shorturl/components/short-url/internal/stores/leveldb"
	"github.com/syndtr/goleveldb/leveldb"
)

func main() {
	// settings
	host := flag.String("host", "localhost", "Host of the url shortener")
	port := flag.Int("port", 8080, "Port for the api server")
	db := flag.String("db", "", "Path to leveldb")
	endpoint := flag.String("counter", "", "Endpoint to the counter API")
	flag.Parse()

	if *db == "" {
		panic("LevelDB path is required")
	}

	database, err := leveldb.OpenFile(*db, nil)
	if err != nil {
		panic("LevelDB open failed")
	}
	defer database.Close()

	store := stores.LevelDBStore{DB: database}

	generator := services.KeyGenerator{Endpoint: *endpoint}

	service := services.ShortURLService{Store: &store, KeyGenerator: generator}

	server := server.Server{Host: *host, ShortURLService: service}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), &server))
}
