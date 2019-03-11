package main

import (
	"flag"
	"fmt"
	"gostudio/exercise3/cyoa"
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	port := flag.Int("port", 3000, "The port to start the CYOA web application on")
	filename := flag.String("file", "gopher.json", "The JSON file with the CYOA story")
	templ := flag.String("template", "", "HTML Template to use")
	flag.Parse()

	fmt.Printf("Using the story in %s.\n", *filename)
	fmt.Printf("Using the template in %s.\n", *templ)
	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	story, err := cyoa.JSONStory(f)
	if err != nil {
		panic(err)
	}

	var t *template.Template

	if *templ != "" {
		fmt.Printf("eccolo %s.\n", *templ)
		t = template.Must(template.New("").ParseFiles(*templ))
	}

	h := cyoa.NewHandler(story, t)
	fmt.Printf("starting the server at: %d\n", *port)
	addr := fmt.Sprintf(":%d", *port)
	log.Fatal(http.ListenAndServe(addr, h))
}
