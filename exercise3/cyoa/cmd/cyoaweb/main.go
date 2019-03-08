package main

import (
	"flag"
	"fmt"
	"gostudio/exercise3/cyoa"
	"os"
)

func main() {
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
	fmt.Printf("%+v\n", story)
}