package main

import (
	"flag"
	"fmt"
	"gostudio/exercise4/link"
	"io"
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

	fmt.Println("Get -> " + *urlFlag)

	pages := get(urlFlag)
	for _, page := range pages {
		fmt.Println(page)
	}

}

func get(urlStr *string) []string {
	resp, err := http.Get(*urlStr)
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

	return filter(
		hrefs(resp.Body, base),
		withPrefex(base),
	)

}

type filterFunc func(string) bool

func hrefs(r io.Reader, base string) []string {
	links, _ := link.Parse(r)
	var ret []string
	for _, l := range links {
		switch {
		case strings.HasPrefix(l.Href, "/"):
			ret = append(ret, base+l.Href)
		case strings.HasPrefix(l.Href, "http"):
			ret = append(ret, l.Href)
		}
	}

	return ret
}

func filter(links []string, keepFns ...filterFunc) []string {
	var ret []string
LINK:
	for _, link := range links {
		for _, keepFn := range keepFns {
			if !keepFn(link) {
				continue LINK
			}
		}
		ret = append(ret, link)

	}
	return ret
}

func withPrefex(pfx string) func(string) bool {
	return func(link string) bool {
		return strings.HasPrefix(link, pfx)
	}
}
