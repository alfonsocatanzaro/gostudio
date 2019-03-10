package main

import (
	"flag"
	"fmt"
	"gostudio/exercise3/cyoa"
	"log"
	"net/http"
	"os"
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
	h := cyoa.NewHandler(story)
	fmt.Printf("starting the server at: %d\n", *port)
	addr := fmt.Sprintf(":%d", *port)
	log.Fatal(http.ListenAndServe(addr, h))
}
