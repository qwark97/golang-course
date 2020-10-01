package cyoa

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
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

func defaultPathFunc(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)
	if path == "" || path == "/" {
		path = "/intro"
	}
	return path[1:]
}

// HandlerOpttion to allow optional parameters
type HandlerOpttion func(h *handler)

// WithTemplate allows to add custom template to story
func WithTemplate(t *template.Template) HandlerOpttion {
	return func(h *handler) {
		h.t = t
	}
}

// WithPathFunc allows to add custom paths handler
func WithPathFunc(pathFunc func(r *http.Request) string) HandlerOpttion {
	return func(h *handler) {
		h.pathFunc = pathFunc
	}
}

// NewHandler serves passed story
func NewHandler(s Story, opts ...HandlerOpttion) http.Handler {
	h := handler{s, tpl, defaultPathFunc}
	for _, opt := range opts {
		opt(&h)
	}
	return h
}

type handler struct {
	s        Story
	t        *template.Template
	pathFunc func(r *http.Request) string
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := h.pathFunc(r)

	if chapter, ok := h.s[path]; ok {
		err := h.t.Execute(w, chapter)
		HTTPErrHandle(w, err)
		return
	}
	http.Error(w, "Chapter not found.", http.StatusNotFound)
}

// JSONStory reads stroy from JSON file
func JSONStory(r io.Reader) (Story, error) {
	var story Story
	decoder := json.NewDecoder(r)

	if err := decoder.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil
}

// ErrHandle is helper function to handle error during runtime
func ErrHandle(err error) {
	if err != nil {
		log.Panic(err)
	}
}

// HTTPErrHandle is helper function to handle http internal server error
func HTTPErrHandle(w http.ResponseWriter, err error) {
	if err != nil {
		log.Println(err)
		http.Error(w, "Something went wrong...", http.StatusInternalServerError)
	}
}

// Story to Create Your Own Adwenture
type Story map[string]Chapter

// Chapter is a part of the story
type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

// Option to choose to move story forward
type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}
