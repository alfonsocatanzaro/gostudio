package cyoa

// Story contain the story loaded by json
type Story map[string]Chapter

// Chapet contains a single chapter of the story
type Chapter struct {
	Title      string   `json:"title"`
	paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

// Option that user hate to choose
type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}
