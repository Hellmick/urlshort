package urlshort

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func generateHandler(url string) func(http.ResponseWriter, *http.Request) {

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("There is an issue with HTTP request")
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Failed to read response body")
	}

	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, string(body))
	}
}

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {

	handlerFunc := func(w http.ResponseWriter, req *http.Request) {
		for path, url := range pathsToUrls {
			handler := generateHandler(url)
			http.HandleFunc(path, handler)
		}
	}

	return handlerFunc
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//   - path: /some-path
//     url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// TODO: Implement this...
	return nil, nil
}
