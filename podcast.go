package podcasts

// Podcast represents a web podcast.
type Podcast struct {
	Title       string
	Description string
	Link        string
	Language    string
	Copyright   string
	items       []*Item
}

// AddItem adds an item to the podcast.
func (p *Podcast) AddItem(item *Item) {
	p.items = append(p.items, item)
}

// Feed creates a new feed for current podcast.
func (p *Podcast) Feed(options ...func(f *Feed) error) (*Feed, error) {
	f := &Feed{
		ItunesXMLNS:  itunesXMLNS,
		ContentXMLNS: contentXMLNS
		Version:      rssVersion,
		Channel:      &Channel{
			Title:       p.Title,
			Description: p.Description,
			Link:        p.Link,
			Copyright:   p.Copyright,
			Language:    p.Language,
			Items:       p.items,
		},
	}
	return f, f.SetOptions(options...)
}
