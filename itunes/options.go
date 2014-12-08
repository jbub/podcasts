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

func Author(author string) func(s *Settings) error {
	return func(s *Settings) error {
		s.Author = author
		return nil
	}
}

func Block(s *Settings) error {
	s.Block = ValueYes
	return nil
}

func Explicit(s *Settings) error {
	s.Explicit = ValueYes
	return nil
}

func Complete(s *Settings) error {
	s.Complete = ValueYes
	return nil
}

func NewFeedURL(newURL string) func(s *Settings) error {
	return func(s *Settings) error {
		u, err := url.Parse(newURL)
		if err != nil {
			return err
		}
		if !u.IsAbs() {
			return ErrInvalidURL
		}
		s.NewFeedURL = newURL
		return nil
	}
}

func Subtitle(subtitle string) func(s *Settings) error {
	return func(s *Settings) error {
		s.Subtitle = subtitle
		return nil
	}
}

func Summary(summary string) func(s *Settings) error {
	return func(s *Settings) error {
		s.Summary = summary
		return nil
	}
}

func Owner(name string, email string) func(s *Settings) error {
	return func(s *Settings) error {
		s.Owner = name
		s.Email = email
		return nil
	}
}

func Image(href string) func(s *Settings) error {
	return func(s *Settings) error {
		u, err := url.Parse(href)
		if err != nil {
			return err
		}
		if !u.IsAbs() {
			return ErrInvalidImage
		}
		s.Image = href
		return nil
	}
}
