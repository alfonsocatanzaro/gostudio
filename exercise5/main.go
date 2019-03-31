package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"gostudio/exercise4/link"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

/*
	1. GET the webpage
	2. parse all the links on the page
	3. build proper urls with our links
	4. filter out any links w/ a diff domain
	5. Find all pages (BFS)
	6. print out XML
*/

type loc struct {
	Value string `xml:"loc"`
}
type urlset struct {
	Urls  []loc  `xml:"url"`
	Xmlns string `xml:"xmlns,attr"`
}

const xmlns = "http://www.sitemaps.org/schemas/sitemap/0.9"

func main() {
	urlFlag := flag.String("url", "https://gophercises.com/", "Url of the site to build sitemap from.")
	maxDepth := flag.Int("depth", 10, "the max number of links deep to traverse.")
	flag.Parse()

	pages := bfs(*urlFlag, *maxDepth)
	toXml := urlset{
		Xmlns: xmlns,
	}
	for _, page := range pages {
		toXml.Urls = append(toXml.Urls, loc{page})
	}

	fmt.Printf(xml.Header)
	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("", "    ")
	if err := enc.Encode(toXml); err != nil {
		panic(err)
	}

}

func bfs(urlStr string, maxDepth int) []string {
	seen := make(map[string]struct{})
	var q map[string]struct{}
	nq := map[string]struct{}{
		urlStr: {},
	}

	for i := 0; i <= maxDepth; i++ {
		q, nq = nq, make(map[string]struct{})
		if len(q) == 0 {
			break
		}
		for u := range q {
			if _, ok := seen[u]; ok {
				continue
			}
			seen[u] = struct{}{}

			for _, l := range get(&u) {
				if _, ok := seen[l]; !ok {
					nq[l] = struct{}{}
				}
			}
		}
	}
	var ret = make([]string, 0, len(seen))
	for u := range seen {
		ret = append(ret, u)
	}
	return ret
}

func get(urlStr *string) []string {
	resp, err := http.Get(*urlStr)
	if err != nil {
		return []string{}
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
		withPrefix(base),
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
	for _, l := range links {
		for _, keepFn := range keepFns {
			if !keepFn(l) {
				continue LINK
			}
		}
		ret = append(ret, l)

	}
	return ret
}

func withPrefix(pfx string) func(string) bool {
	return func(link string) bool {
		return strings.HasPrefix(link, pfx)
	}
}
