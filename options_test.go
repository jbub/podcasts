package podcasts

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type OptionsTestSuite struct {
	suite.Suite
	feed *Feed
}

func (s *OptionsTestSuite) SetupTest() {
	s.feed = &Feed{
		Channel: &Channel{},
	}
}

func (s *OptionsTestSuite) TestAuthor() {
	author := "john"
	err := Author(author)(s.feed)

	s.Nil(err)
	s.Equal(author, s.feed.Channel.Author)
}

func (s *OptionsTestSuite) TestBlock() {
	err := Block(s.feed)

	s.Nil(err)
	s.Equal(ValueYes, s.feed.Channel.Block)
}

func (s *OptionsTestSuite) TestExplicit() {
	err := Explicit(s.feed)

	s.Nil(err)
	s.Equal(ValueYes, s.feed.Channel.Explicit)
}

func (s *OptionsTestSuite) TestComplete() {
	err := Complete(s.feed)

	s.Nil(err)
	s.Equal(ValueYes, s.feed.Channel.Complete)
}

func (s *OptionsTestSuite) TestNewFeedURL() {
	newURL := "http://example.com/test"
	err := NewFeedURL(newURL)(s.feed)

	s.Nil(err)
	s.Equal(newURL, s.feed.Channel.NewFeedURL)
}

func (s *OptionsTestSuite) TestSubtitle() {
	subtitle := "this is subtitle"
	err := Subtitle(subtitle)(s.feed)

	s.Nil(err)
	s.Equal(subtitle, s.feed.Channel.Subtitle)
}

func (s *OptionsTestSuite) TestSummary() {
	summary := "this is summary"
	err := Summary(summary)(s.feed)

	s.Nil(err)
	s.Equal(summary, s.feed.Channel.Summary)
}

func (s *OptionsTestSuite) TestOwner() {
	name := "anabelle"
	email := "test@test.com"
	err := Owner(name, email)(s.feed)

	s.Nil(err)
	s.Equal(name, s.feed.Channel.Owner.Name)
	s.Equal(email, s.feed.Channel.Owner.Email)
}

func (s *OptionsTestSuite) TestImage() {
	href := "http://example.com/test/image.jpg"
	err := Image(href)(s.feed)

	s.Nil(err)
	s.Equal(href, s.feed.Channel.Image.Href)
}

func TestOptionsTestSuite(t *testing.T) {
	suite.Run(t, new(OptionsTestSuite))
}
