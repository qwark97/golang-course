package cyoa

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Parse(defaultHandlerTmpl))
}

var defaultHandlerTmpl = `
<!DOCTYPE html>
<html>
    <head>
        <meta charset="utf-8">
        <title>Choose Your Own Adwenture</title>
    </head>
    <body>
        <h1>{{.Title}}</h1>
        {{range .Paragraphs}}
        <p>{{.}}</p>
        {{end}}
        <ul>
            {{range .Options}}
			<li><a href="/{{.Chapter}}">{{.Text}}</a></li>
            {{end}}
        </ul>
    </body>
</html>
`

func NewHandler(s Story) http.Handler {
	return handler{s}
}

type handler struct {
	s Story
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := tpl.Execute(w, h.s["intro"])
	ErrHandle(err)
}

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
