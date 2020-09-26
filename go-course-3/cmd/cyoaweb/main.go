package main

import (
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
	cyoa.ErrHandle(err)

	story, err := cyoa.JsonStory(file)
	cyoa.ErrHandle(err)

	fmt.Printf("%+v\n", story)
}
