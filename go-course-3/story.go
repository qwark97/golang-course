package cyoa

import (
	"encoding/json"
	"io"
	"log"
)

func JsonStory(r io.Reader) (Story, error) {
	var story Story
	decoder := json.NewDecoder(r)

	if err := decoder.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil
}

func ErrHandle(err error) {
	if err != nil {
		log.Panic(err)
	}
}

type Story map[string]Chapter

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}
