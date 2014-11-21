package podcasts

type Podcast struct {
	Title       string
	Description string
	URL         string
	Copyright   string
	Language    string
	categories  []*Category
	items       []*Item
}

func (p *Podcast) AddItem(item *Item) {
	p.items = append(p.items, item)
}

func (p *Podcast) AddCategory(category *Category) {
	p.categories = append(p.categories, category)
}

func (p *Podcast) Feed() *Feed {
	return &Feed{
		Xmlns:   rssXmlns,
		Version: rssVersion,
		Channel: &Channel{
			Title:       p.Title,
			Description: p.Description,
			Link:        p.URL,
			Copyright:   p.Copyright,
			Language:    p.Language,
			Items:       p.items,
			Categories:  p.categories,
		},
	}
}
