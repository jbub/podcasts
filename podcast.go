package podcasts

import "github.com/jbub/podcasts/itunes"

type Podcast struct {
	Title       string
	Description string
	Link        string
	Language    string
	Copyright   string
	itunes      *itunes.Options
	items       []*Item
}

func (p *Podcast) AddItem(item *Item) {
	p.items = append(p.items, item)
}

func (p *Podcast) SetItunes(opts *itunes.Options) {
	p.itunes = opts
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
	if p.itunes != nil {
		p.setupItunesFeed(f)
	}
	return f
}

func (p *Podcast) setupItunesFeed(f *Feed) {
	f.Channel.Author = p.itunes.Author
	f.Channel.Block = p.itunes.Block
	f.Channel.Explicit = p.itunes.Explicit
	f.Channel.Complete = p.itunes.Complete
	f.Channel.NewFeedURL = p.itunes.NewFeedURL
	f.Channel.Subtitle = p.itunes.Subtitle
	f.Channel.Summary = p.itunes.Summary

	for _, c := range p.itunes.Categories {
		f.Channel.Categories = append(f.Channel.Categories, &Category{Text: c.Text})
	}

	if p.itunes.Owner != "" && p.itunes.Email != "" {
		f.Channel.Owner = &Owner{
			Name:  p.itunes.Owner,
			Email: p.itunes.Email,
		}
	}

	if p.itunes.Image != "" {
		f.Channel.Image = &Image{
			Href: p.itunes.Image,
		}
	}
}
