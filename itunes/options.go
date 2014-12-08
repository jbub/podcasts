package itunes

import (
	"errors"
	"net/url"
)

var (
	ErrInvalidURL   = errors.New("itunes: invalid url")
	ErrInvalidImage = errors.New("itunes: invalid image")
)

const (
	ValueYes = "yes"
)

func Author(author string) func(i *ChannelOpts) error {
	return func(i *ChannelOpts) error {
		i.Author = author
		return nil
	}
}

func Block(i *ChannelOpts) error {
	i.Block = ValueYes
	return nil
}

func Explicit(i *ChannelOpts) error {
	i.Explicit = ValueYes
	return nil
}

func Complete(i *ChannelOpts) error {
	i.Complete = ValueYes
	return nil
}

func NewFeedURL(newURL string) func(i *ChannelOpts) error {
	return func(i *ChannelOpts) error {
		u, err := url.Parse(newURL)
		if err != nil {
			return err
		}
		if !u.IsAbs() {
			return ErrInvalidURL
		}
		i.NewFeedURL = newURL
		return nil
	}
}

func Subtitle(subtitle string) func(i *ChannelOpts) error {
	return func(i *ChannelOpts) error {
		i.Subtitle = subtitle
		return nil
	}
}

func Summary(summary string) func(i *ChannelOpts) error {
	return func(i *ChannelOpts) error {
		i.Summary = summary
		return nil
	}
}

func Owner(name string, email string) func(i *ChannelOpts) error {
	return func(i *ChannelOpts) error {
		i.Owner = name
		i.Email = email
		return nil
	}
}

func Image(href string) func(i *ChannelOpts) error {
	return func(i *ChannelOpts) error {
		u, err := url.Parse(href)
		if err != nil {
			return err
		}
		if !u.IsAbs() {
			return ErrInvalidImage
		}
		i.Image = href
		return nil
	}
}
