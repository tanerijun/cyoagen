# CYOA-Generator

A generator for [Create Your Own Adventure](https://en.wikipedia.org/wiki/Choose_Your_Own_Adventure) books.

# Usage

Simply provide the books a path to the JSON file containing the story, and CYOA-Generator will generate a webpage that allows you to read the book.

Below is the format for the story JSON file:

```json
{
  // Each story arc will have a unique key that represents
  // the name of that particular arc.
  "story-arc": {
    "title": "A title for that story arc. Think of it like a chapter title.",
    "story": [
      "A series of paragraphs, each represented as a string in a slice.",
      "This is a new paragraph in this particular story arc."
    ],
    // Options will be empty if it is the end of that
    // particular story arc. Otherwise it will have one or
    // more JSON objects that represent an "option" that the
    // reader has at the end of a story arc.
    "options": [
      {
        "text": "the text to render for this option. eg 'venture down the dark passage'",
        "arc": "the name of the story arc to navigate to. This will match the story-arc key at the very root of the JSON document"
      }
    ]
  }
}
```

# Installation

1. Install Go
2. Clone repo
3. Build the app
   ```
   go build -o cyoagen
   ```
4. Run the binary with the path to the JSON story file as argument
   ```
   cyoagen stories/example.json
   ```
5. You can replace the example.json will your own version
6. Optionally, you can also pass the `-p` flag to host the server on custom port (default: 8080).
   ```
   cyoagen -p=3333 stories/example.json
   ```
7. Optionally, you can also pass the `-t` flag to provide a custom template file (default: "template/main.html").
   ```
   cyoagen -p=template/custom.html stories/example.json
   ```
8. Visit the link printed on your terminal and enjoy!
