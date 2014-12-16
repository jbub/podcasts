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

func Author(author string) func(f *Feed) error {
	return func(f *Feed) error {
		f.Channel.Author = author
		return nil
	}
}

func Block(f *Feed) error {
	f.Channel.Block = ValueYes
	return nil
}

func Explicit(f *Feed) error {
	f.Channel.Explicit = ValueYes
	return nil
}

func Complete(f *Feed) error {
	f.Channel.Complete = ValueYes
	return nil
}

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

func Subtitle(subtitle string) func(f *Feed) error {
	return func(f *Feed) error {
		f.Channel.Subtitle = subtitle
		return nil
	}
}

func Summary(summary string) func(f *Feed) error {
	return func(f *Feed) error {
		f.Channel.Summary = summary
		return nil
	}
}

func Owner(name string, email string) func(f *Feed) error {
	return func(f *Feed) error {
		f.Channel.Owner = &ItunesOwner{
			Name:  name,
			Email: email,
		}
		return nil
	}
}

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
