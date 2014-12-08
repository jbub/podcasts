package podcasts

import "github.com/jbub/podcasts/itunes"

type Podcast struct {
	Title       string
	Description string
	Link        string
	Language    string
	Copyright   string
	options     *itunes.Settings
	items       []*Item
}

func (p *Podcast) AddItem(item *Item) {
	p.items = append(p.items, item)
}

func (p *Podcast) SetOptions(s *itunes.Settings) {
	p.options = s
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
	if p.options != nil {
		p.setupItunesFeed(f)
	}
	return f
}

func (p *Podcast) setupItunesFeed(f *Feed) {
	f.Channel.Author = p.options.Author
	f.Channel.Block = p.options.Block
	f.Channel.Explicit = p.options.Explicit
	f.Channel.Complete = p.options.Complete
	f.Channel.NewFeedURL = p.options.NewFeedURL
	f.Channel.Subtitle = p.options.Subtitle
	f.Channel.Summary = p.options.Summary

	for _, c := range p.options.Categories {
		f.Channel.Categories = append(f.Channel.Categories, &Category{Text: c.Text})
	}

	if p.options.Owner != "" && p.options.Email != "" {
		f.Channel.Owner = &Owner{
			Name:  p.options.Owner,
			Email: p.options.Email,
		}
	}

	if p.options.Image != "" {
		f.Channel.Image = &Image{
			Href: p.options.Image,
		}
	}
}
