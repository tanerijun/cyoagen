package cyoa

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"text/template"
)

func TestJSONStoryParser(t *testing.T) {
	_, err := JSONStoryParser(strings.NewReader("nonexistent"))
	if err == nil {
		t.Error("expected nil, got error")
	}

	f, err := os.Open("testdata/example.json")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	defer f.Close()

	_, err = JSONStoryParser(f)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestStoryHandler(t *testing.T) {
	f, err := os.Open("testdata/example.json")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	defer f.Close()

	story, err := JSONStoryParser(f)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	htmlBytes, err := os.ReadFile("testdata/template.html")
	if err != nil {
		log.Fatal(err)
	}
	tmpl := template.Must(template.New(filepath.Base(("testdata/template.html"))).Parse(string(htmlBytes)))

	sh := NewStoryHandler(story, tmpl)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	sh.ServeHTTP(w, req)
	resp := w.Result()

	if resp.StatusCode != 200 {
		t.Errorf("Expected 200 StatusCode, got %d", resp.StatusCode)
	}

	// should get back an html
	contentType := resp.Header["Content-Type"]
	if !strings.Contains(contentType[0], "html") {
		t.Errorf("Expected html, got %v", contentType[0])
	}

	// should redirect non-existent path to start page, no error 404 should happen
	req = httptest.NewRequest(http.MethodGet, "/nonexistent", nil)
	sh.ServeHTTP(w, req)
	resp = w.Result()

	if resp.StatusCode != 200 {
		t.Errorf("Expected 200 StatusCode, got %d", resp.StatusCode)
	}
}
