package podcasts

type Podcast struct {
	Title       string
	Description string
	Link        string
	Language    string
	Copyright   string
	items       []*Item
}

func (p *Podcast) AddItem(item *Item) {
	p.items = append(p.items, item)
}

func (p *Podcast) Feed() *Feed {
	f := &Feed{
		Xmlns:   rssXmlns,
		Version: rssVersion,
		Channel: &Channel{
			Title:       p.Title,
			Description: p.Description,
			Link:        p.Link,
			Copyright:   p.Copyright,
			Language:    p.Language,
			Items:       p.items,
		},
	}
	return f
}
