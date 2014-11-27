package podcasts

type Itunes struct {
	Author     string
	Subtitle   string
	Summary    string
	Owner      string
	Email      string
	Explicit   string
	Image      string
	categories []*Category
}

func (i *Itunes) AddCategory(category *Category) {
	i.categories = append(i.categories, category)
}

type Podcast struct {
	Title       string
	Description string
	Link        string
	Language    string
	Copyright   string
	Itunes      *Itunes
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
	if p.Itunes != nil {
		setupItunes(f.Channel, p.Itunes)
	}
	return f
}

func setupItunes(c *Channel, itunes *Itunes) {
	c.Author = itunes.Author
	c.Subtitle = itunes.Subtitle
	c.Summary = itunes.Summary
	c.Explicit = itunes.Explicit
	c.Categories = itunes.categories

	if itunes.Owner != "" && itunes.Email != "" {
		c.Owner = &Owner{
			Name:  itunes.Owner,
			Email: itunes.Email,
		}
	}

	if itunes.Image != "" {
		c.Image = &Image{
			Href: itunes.Image,
		}
	}
}
