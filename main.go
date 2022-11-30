package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"text/template"

	"github.com/tanerijun/cyoa-generator/cyoa"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: cyoagen <path_to_story_file>")
		os.Exit(1)
	}

	port := flag.Int("p", 8080, "The server port.")
	tpath := flag.String("t", "template/main.html", "The HTML template file.")
	flag.Parse()

	f, err := os.Open(os.Args[len(os.Args)-1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	story, err := cyoa.JSONStoryParser(f)
	if err != nil {
		log.Fatal(err)
	}

	htmlBytes, err := os.ReadFile(*tpath)
	if err != nil {
		log.Fatal(err)
	}

	tmpl := template.Must(template.New(filepath.Base((*tpath))).Parse(string(htmlBytes)))

	h := cyoa.NewStoryHandler(story, tmpl)

	fmt.Printf("Server is live on port %d.\n", *port)
	fmt.Printf("http://localhost:%d\n", *port)
	addr := fmt.Sprintf(":%d", *port)
	log.Fatal(http.ListenAndServe(addr, h))
}
