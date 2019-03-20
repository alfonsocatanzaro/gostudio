package main

import (
	"flag"
	"fmt"
	"gostudio/exercise4/link"
	"net/http"
	"net/url"
	"strings"
)

/*
	1. GET the webage
	2. parse all the links on the page
	3. build proper urls with our links
	4. filter out any links w/ a diff domain
	5. Find all pages (BFS)
	6. print out XML
*/

func main() {
	urlFlag := flag.String("url", "https://gophercises.com/", "Url of the site to build sitemat from.")
	flag.Parse()

	resp, err := http.Get(*urlFlag)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	reqURL := resp.Request.URL
	baseURL := &url.URL{
		Scheme: reqURL.Scheme,
		Host:   reqURL.Host,
	}

	base := baseURL.String()

	links, _ := link.Parse(resp.Body)
	var hrefs []string
	for _, l := range links {
		switch {
		case strings.HasPrefix(l.Href, "/"):
			hrefs = append(hrefs, base+l.Href)
		case strings.HasPrefix(l.Href, "http"):
			hrefs = append(hrefs, l.Href)
		}
	}

	for _, l := range hrefs {
		fmt.Println(l)
	}

	/*
		/some-path
		https://gophercises.com/some-path
		http://gophercises.com/some-path
		#fragments
		mailto:mail@domain.com
	*/

}
