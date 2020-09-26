package main

import (
	"encoding/json"
	"flag"
	"fmt"
	cyoa "go-course/go-course-3"
	"os"
)

func main() {
	filename := flag.String("file", "gopher.json", "the JSON file file with CYOA story")
	flag.Parse()

	fmt.Printf("Using the story in %s.\n\n", *filename)

	file, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	decoder := json.NewDecoder(file)
	var story cyoa.Story

	if err := decoder.Decode(&story); err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", story)
}
