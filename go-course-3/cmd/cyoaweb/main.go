package main

import (
	"flag"
	"fmt"
	cyoa "go-course/go-course-3"
	"log"
	"net/http"
	"os"
)

func main() {
	port := flag.Int("port", 3000, "the port to start the CYOA web application on")
	filename := flag.String("file", "gopher.json", "the JSON file file with CYOA story")
	flag.Parse()

	log.Printf("Using the story in %s.\n\n", *filename)

	file, err := os.Open(*filename)
	cyoa.ErrHandle(err)

	story, err := cyoa.JsonStory(file)
	cyoa.ErrHandle(err)

	h := cyoa.NewHandler(story)
	log.Printf("Starting the server on: http://localhost:%d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
