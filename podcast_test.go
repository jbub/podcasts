package podcasts

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ValidFeed struct {
	Channel struct {
		Title string `xml:"title"`
	} `xml:"channel"`
}

type PodcastsTestSuite struct {
	suite.Suite
}

func (s *PodcastsTestSuite) TestPodcastFeed() {
	title := "test TITLE"
	p := &Podcast{
		Title:       title,
		Description: "Zábavný program pre každého, komu to páli.",
		Language:    "SK",
		Link:        "http://www.rtvs.sk/radio/relacie/detail/palenica-borisa-filana",
		Copyright:   "2013 RTVS - Rozhlas a televízia Slovenska",
	}

	feed, err := p.Feed()
	s.Nil(err)

	data, err := feed.XML()
	s.Nil(err)

	v := ValidFeed{}
	err = xml.Unmarshal([]byte(data), &v)

	s.Nil(err)
	s.Equal(title, v.Channel.Title)
}

func TestPodcastsTestSuite(t *testing.T) {
	suite.Run(t, new(PodcastsTestSuite))
}
