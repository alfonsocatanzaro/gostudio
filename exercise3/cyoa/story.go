package cyoa

import (
	"encoding/json"
	"io"
)

// JSONStory load story from a reader object.
func JSONStory(r io.Reader) (Story, error) {
	d := json.NewDecoder(r)
	var story Story
	err := d.Decode(&story)
	if err != nil {
		return nil, err
	}
	return story, nil
}

// Story contain the story loaded by json
type Story map[string]Chapter

// Chapter contains a single chapter of the story
type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

// Option that user hate to choose
type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}
