package link

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

// Link represents a link (<a hfref="...">...<a>) in a HTML
// document.
type Link struct {
	Href string
	Text string
}

// Parse will take in an HTML document and will return a
// slice of links parsed from it
func Parse(r io.Reader) ([]Link, error) {
	//links := make(Link[])
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	dfs(doc, "")
	return nil, nil
}

func dfs(n *html.Node, padding string) {
	msg := n.Data
	if n.Type == html.ElementNode {
		if n.Data == "a" {
			if href, ok := getAttributeValue(n, "href"); ok {
				msg = msg + ` href="` + href + `"`
			}
		}
		msg = "<" + msg + ">"
	}

	fmt.Println(padding, msg)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		dfs(c, padding+"  ")
	}
}

func getAttributeValue(n *html.Node, name string) (string, bool) {
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			return attr.Val, true
		}
	}
	return "", false
}
