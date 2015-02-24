package podcasts

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type PodcastsTestSuite struct {
	suite.Suite
	podcast *Podcast
}

func (s *PodcastsTestSuite) SetupTest() {
	s.podcast = &Podcast{
		Title:       "my podcast title",
		Description: "my podcast description",
		Language:    "my podcast lang",
		Link:        "my podcast link",
		Copyright:   "my podcast copyright",
	}
}

func (s *PodcastsTestSuite) TestContainsXmlHeader() {
	data, err := getPodcastXML(s.podcast)

	s.Nil(err)
	s.Contains(data, `<?xml version="1.0" encoding="UTF-8"?>`)
}

func (s *PodcastsTestSuite) TestContainsRssElement() {
	data, err := getPodcastXML(s.podcast)

	s.Nil(err)
	s.Contains(data, `<rss xmlns:itunes="http://www.itunes.com/dtds/podcast-1.0.dtd" version="2.0">`)
}

func (s *PodcastsTestSuite) TestContainsChannelElement() {
	data, err := getPodcastXML(s.podcast)

	s.Nil(err)
	s.Contains(data, `<channel>`)
	s.Contains(data, `</channel>`)
}

func (s *PodcastsTestSuite) TestContainsTitleElement() {
	data, err := getPodcastXML(s.podcast)

	s.Nil(err)
	s.Contains(data, fmt.Sprintf("<title>%s</title>", s.podcast.Title))
}

func (s *PodcastsTestSuite) TestContainsDescriptionElement() {
	data, err := getPodcastXML(s.podcast)

	s.Nil(err)
	s.Contains(data, fmt.Sprintf("<description>%s</description>", s.podcast.Description))
}

func (s *PodcastsTestSuite) TestContainsLanguageElement() {
	data, err := getPodcastXML(s.podcast)

	s.Nil(err)
	s.Contains(data, fmt.Sprintf("<language>%s</language>", s.podcast.Language))
}

func (s *PodcastsTestSuite) TestContainsLinkElement() {
	data, err := getPodcastXML(s.podcast)

	s.Nil(err)
	s.Contains(data, fmt.Sprintf("<link>%s</link>", s.podcast.Link))
}

func (s *PodcastsTestSuite) TestContainsCopyrightElement() {
	data, err := getPodcastXML(s.podcast)

	s.Nil(err)
	s.Contains(data, fmt.Sprintf("<copyright>%s</copyright>", s.podcast.Copyright))
}

func (s *PodcastsTestSuite) TestPodcastFeedWrite() {
	feed, err := s.podcast.Feed()
	s.Nil(err)

	var b bytes.Buffer
	err = feed.Write(&b)
	s.Nil(err)
}

func TestPodcastsTestSuite(t *testing.T) {
	suite.Run(t, new(PodcastsTestSuite))
}

func getPodcastXML(p *Podcast) (string, error) {
	feed, err := p.Feed()
	if err != nil {
		return "", err
	}
	return feed.XML()
}
