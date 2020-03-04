package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/isflavior/shorturl/components/counter/internal/server"
)

func main() {
	// settings
	path := flag.String("path", "", "Path to the counter data file")
	port := flag.Int("port", 8181, "Port for the counter server")
	flag.Parse()

	if *path == "" {
		panic("Path is required")
	}

	init := 0

	data, err := ioutil.ReadFile(*path)
	if err != nil {
		err := ioutil.WriteFile(*path, []byte(strconv.Itoa(0)), 0644)
		if err != nil {
			panic("Counter file failed to open")
		}
	} else {
		init, _ = strconv.Atoi(string(data))
	}

	server := server.Server{FilePath: *path, Counter: init}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), &server))
}
