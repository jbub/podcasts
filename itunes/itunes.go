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

func Options(options ...func(i *ChannelOpts) error) (*ChannelOpts, error) {
	opts := ChannelOpts{}
	return &opts, opts.SetOption(options...)
}

type Category struct {
	Text       string
	Categories []*Category
}

type ChannelOpts struct {
	Author     string
	Block      string
	Explicit   string
	Complete   string
	NewFeedURL string
	Subtitle   string
	Summary    string
	Owner      string
	Email      string
	Image      string
	Categories []*Category
}

func (i *ChannelOpts) SetOption(options ...func(*ChannelOpts) error) error {
	for _, opt := range options {
		if err := opt(i); err != nil {
			return err
		}
	}
	return nil
}

func (i *ChannelOpts) AddCategory(c *Category) {
	i.Categories = append(i.Categories, c)
}

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
