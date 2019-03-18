package main

import (
	"flag"
	"fmt"
)

func main() {
	urlFlag := flag.String("url", "https://gophercises.com", "Url of the site to build sitemat from.")
	flag.Parse()

	fmt.Println(*urlFlag)

	/*
		1. GET the webage
		2. parse all the links on the page
		3. build proper urls with our links
		4. filter out any links w/ a diff domain
		5. Find all pages (BFS)
		6. print out XML
	*/
}
