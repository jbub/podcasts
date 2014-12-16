package podcasts

import (
	"errors"
	"net/url"
)

var (
	ErrInvalidURL   = errors.New("podcasts: invalid url")
	ErrInvalidImage = errors.New("podcasts: invalid image")
)

const (
	ValueYes = "yes"
)

func ItunesAuthor(author string) func(f *Feed) error {
	return func(f *Feed) error {
		f.Channel.Author = author
		return nil
	}
}

func ItunesBlock(f *Feed) error {
	f.Channel.Block = ValueYes
	return nil
}

func ItunesExplicit(f *Feed) error {
	f.Channel.Explicit = ValueYes
	return nil
}

func ItunesComplete(f *Feed) error {
	f.Channel.Complete = ValueYes
	return nil
}

func ItunesNewFeedURL(newURL string) func(f *Feed) error {
	return func(f *Feed) error {
		u, err := url.Parse(newURL)
		if err != nil {
			return err
		}
		if !u.IsAbs() {
			return ErrInvalidURL
		}
		f.Channel.NewFeedURL = newURL
		return nil
	}
}

func ItunesSubtitle(subtitle string) func(f *Feed) error {
	return func(f *Feed) error {
		f.Channel.Subtitle = subtitle
		return nil
	}
}

func ItunesSummary(summary string) func(f *Feed) error {
	return func(f *Feed) error {
		f.Channel.Summary = summary
		return nil
	}
}

func ItunesOwner(name string, email string) func(f *Feed) error {
	return func(f *Feed) error {
		f.Channel.Owner = &Owner{
			Name:  name,
			Email: email,
		}
		return nil
	}
}

func ItunesImage(href string) func(f *Feed) error {
	return func(f *Feed) error {
		u, err := url.Parse(href)
		if err != nil {
			return err
		}
		if !u.IsAbs() {
			return ErrInvalidImage
		}
		f.Channel.Image = &Image{
			Href: href,
		}
		return nil
	}
}
