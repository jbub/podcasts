package podcasts

import (
	"bytes"
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/suite"
)

type validFeed struct {
	Channel struct {
		Title string `xml:"title"`
	} `xml:"channel"`
}

type PodcastsTestSuite struct {
	suite.Suite
	podcast *Podcast
}

func (s *PodcastsTestSuite) SetupTest() {
	title := "test TITLE"
	desc := "description"
	lang := "lang"
	link := "link"
	copyright := "copyright"

	s.podcast = &Podcast{
		Title:       title,
		Description: desc,
		Language:    lang,
		Link:        link,
		Copyright:   copyright,
	}
}

func (s *PodcastsTestSuite) TestPodcastFeedXml() {
	feed, err := s.podcast.Feed()
	s.Nil(err)

	data, err := feed.XML()
	s.Nil(err)

	vf := validFeed{}
	err = xml.Unmarshal([]byte(data), &vf)

	s.Nil(err)
	s.Equal(s.podcast.Title, vf.Channel.Title)
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
