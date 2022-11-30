package cyoa

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
	"text/template"
)

type Story map[string]Arc

type Arc struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type StoryHandler struct {
	s Story
	t *template.Template
}

func (sh StoryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSpace(r.URL.Path)[1:]

	if arc, ok := sh.s[path]; ok {
		err := sh.t.Execute(w, arc)
		if err != nil {
			log.Println(err)
			http.Error(w, "Something went wrong...", http.StatusInternalServerError)
		}
		return
	}

	err := sh.t.Execute(w, sh.s["intro"])
	if err != nil {
		log.Println(err)
		http.Error(w, "Something went wrong.\nThe provided JSON is probably not in the correct format.", http.StatusInternalServerError)
	}
}

func NewStoryHandler(s Story, t *template.Template) StoryHandler {
	return StoryHandler{s, t}
}

func JSONStoryParser(r io.Reader) (Story, error) {
	d := json.NewDecoder(r)
	var story Story
	if err := d.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil
}
