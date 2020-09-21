package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	urlshort "go-course/go-course-2"
)

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the fallback
	yaml, err := readYAML("yamlPaths.yaml")
	if err != nil {
		log.Panic("Error during reading file")
	}
	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		log.Panic(err)
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
	log.Printf("%s - someone hit the page", r.Host)
	fmt.Fprintln(w, "Hello, world!")
}

func readYAML(pathToFile string) ([]byte, error) {
	yamlBytes, err := ioutil.ReadFile(pathToFile)
	if err != nil {
		log.Printf("Cannot read provided file: %s", pathToFile)
		return nil, err
	}
	return yamlBytes, nil
}
