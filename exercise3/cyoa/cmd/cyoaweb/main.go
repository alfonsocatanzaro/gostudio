package main

import (
	"flag"
	"fmt"
	"gostudio/exercise3/cyoa"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	port := flag.Int("port", 3000, "The port to start the CYOA web application on")
	filename := flag.String("file", "gopher.json", "The JSON file with the CYOA story")
	flag.Parse()

	fmt.Printf("Using the story in %s.\n", *filename)
	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	story, err := cyoa.JSONStory(f)
	if err != nil {
		panic(err)
	}

	t := template.Must(template.New("").Parse(storyTemplate))
	h := cyoa.NewHandler(story,
		cyoa.WithPathFunc(pathFn),
		cyoa.WithTemplate(t),
	)
	mux := http.NewServeMux()
	mux.Handle("/story/", h)
	fmt.Printf("starting the server at: %d\n", *port)
	addr := fmt.Sprintf(":%d", *port)
	log.Fatal(http.ListenAndServe(addr, h))
}

func pathFn(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)
	if path == "/story" || path == "/story/" {
		path = "/story/intro"
	}
	return path[len("/story/"):]
}

var storyTemplate = `<!DOCTYPE html>
<html>

<head>
  <meta charset="utf-8">
  <title>Choose You Own Adventure</title>
</head>

<body>
  <section class="page">
    <h1>{{.Title}}</h1>
    {{range .Paragraphs}}
    <p>{{.}}</p>
    {{end}}
    <ul>
      {{range .Options}}
      <li><a href="/story/{{.Arc}}">{{.Text}}</a></li>
      {{end}}
    </ul>
  </section>
  <style>
    body {
      font-family: helvetica, arial;
    }

    h1 {
      text-align: center;
      position: relative;
    }

    .page {
      width: 80%;
      max-width: 500px;
      margin: auto;
      margin-top: 40px;
      margin-bottom: 40px;
      padding: 80px;
      background: #FFFCF6;
      border: 1px solid #eee;
      box-shadow: 0 10px 6px -6px #777;
    }

    ul {
      border-top: 1px dotted #ccc;
      padding: 10px 0 0 0;
      -webkit-padding-start: 0;
    }

    li {
      padding-top: 10px;
    }

    a,
    a:visited {
      text-decoration: none;
      color: #6295b5;
    }

    a:active,
    a:hover {
      color: #7792a2;
    }

    p {
      text-indent: 1em;
    }
  </style>
</body>`
