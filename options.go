package podcasts

import (
	"errors"
	"net/url"
)

var (
	// ErrInvalidURL represents a error returned for invalid url.
	ErrInvalidURL = errors.New("podcasts: invalid url")

	// ErrInvalidImage represents a error returned for invalid image.
	ErrInvalidImage = errors.New("podcasts: invalid image")
)

const (
	// ValueYes represents positive value used in XML feed.
	ValueYes = "yes"
)

// Author sets itunes:author of given feed.
func Author(author string) func(f *Feed) error {
	return func(f *Feed) error {
		f.Channel.Author = author
		return nil
	}
}

// Block enables itunes:block of given feed.
func Block(f *Feed) error {
	f.Channel.Block = ValueYes
	return nil
}

// Explicit enables itunes:explicit of given feed.
func Explicit(f *Feed) error {
	f.Channel.Explicit = ValueYes
	return nil
}

// Complete enables itunes:complete of given feed.
func Complete(f *Feed) error {
	f.Channel.Complete = ValueYes
	return nil
}

// NewFeedURL sets itunes:new-feed-url of given feed.
func NewFeedURL(newURL string) func(f *Feed) error {
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

// Subtitle sets itunes:subtitle of given feed.
func Subtitle(subtitle string) func(f *Feed) error {
	return func(f *Feed) error {
		f.Channel.Subtitle = subtitle
		return nil
	}
}

// Summary sets itunes:summary of given feed.
func Summary(summary string) func(f *Feed) error {
	return func(f *Feed) error {
		f.Channel.Summary = &CDATAText{summary}
		return nil
	}
}

// Owner sets itunes:owner of given feed.
func Owner(name string, email string) func(f *Feed) error {
	return func(f *Feed) error {
		f.Channel.Owner = &ItunesOwner{
			Name:  name,
			Email: email,
		}
		return nil
	}
}

// Image sets itunes:image of given feed.
func Image(href string) func(f *Feed) error {
	return func(f *Feed) error {
		u, err := url.Parse(href)
		if err != nil {
			return err
		}
		if !u.IsAbs() {
			return ErrInvalidImage
		}
		f.Channel.Image = &ItunesImage{
			Href: href,
		}
		return nil
	}
}
