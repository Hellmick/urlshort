package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/Hellmick/urlshort/urlshort"
)

func readYAMLFile(fileLocation *string) ([]byte, error) {

	file, err := os.ReadFile(*fileLocation)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func main() {

	fileLocation := flag.String("f", "config.yaml", "provide paths and urls config")
	flag.Parse()

	mux := defaultMux()

	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	yaml, err := readYAMLFile(fileLocation)
	if err != nil {
		panic(err)
	}

	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
