package podcasts

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

var (
	validItems = []struct {
		Title           string
		GUID            string
		Time            time.Time
		ExpectedPubDate string
		EnclosureURL    string
		EnclosureLength string
		EnclosureType   string
	}{
		{
			Title:           "Item 1",
			GUID:            "http://www.example-podcast.com/my-podcast/1/episode",
			Time:            time.Date(2015, time.January, 1, 0, 0, 0, 0, time.UTC),
			ExpectedPubDate: "Thu, 01 Jan 2015 00:00:00 +0000",
			EnclosureURL:    "http://www.example-podcast.com/my-podcast/1/episode-one",
			EnclosureLength: "1234",
			EnclosureType:   "MP3",
		},
		{
			Title:           "Item 2",
			GUID:            "http://www.example-podcast.com/my-podcast/2/episode",
			Time:            time.Date(2015, time.January, 2, 0, 0, 0, 0, time.UTC),
			ExpectedPubDate: "Fri, 02 Jan 2015 00:00:00 +0000",
			EnclosureURL:    "http://www.example-podcast.com/my-podcast/2/episode-two",
			EnclosureLength: "56445",
			EnclosureType:   "WAV",
		},
	}
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

func (s *PodcastsTestSuite) TestContainsBlockElement() {
	data, err := getPodcastXML(s.podcast, Block)

	s.Nil(err)
	s.Contains(data, fmt.Sprintf("<itunes:block>%s</itunes:block>", ValueYes))
}

func (s *PodcastsTestSuite) TestContainsExplicitElement() {
	data, err := getPodcastXML(s.podcast, Explicit)

	s.Nil(err)
	s.Contains(data, fmt.Sprintf("<itunes:explicit>%s</itunes:explicit>", ValueYes))
}

func (s *PodcastsTestSuite) TestContainsCompleteElement() {
	data, err := getPodcastXML(s.podcast, Complete)

	s.Nil(err)
	s.Contains(data, fmt.Sprintf("<itunes:complete>%s</itunes:complete>", ValueYes))
}

func (s *PodcastsTestSuite) TestContainsAuthorElement() {
	author := "Test Author"
	data, err := getPodcastXML(s.podcast, Author(author))

	s.Nil(err)
	s.Contains(data, fmt.Sprintf("<itunes:author>%s</itunes:author>", author))
}

func (s *PodcastsTestSuite) TestContainsNewFeedURLElement() {
	url := "http://localhost/my-test-url"
	data, err := getPodcastXML(s.podcast, NewFeedURL(url))

	s.Nil(err)
	s.Contains(data, fmt.Sprintf("<itunes:new-feed-url>%s</itunes:new-feed-url>", url))
}

func (s *PodcastsTestSuite) TestContainsSubtitleElement() {
	subtitle := "Test Subtitle"
	data, err := getPodcastXML(s.podcast, Subtitle(subtitle))

	s.Nil(err)
	s.Contains(data, fmt.Sprintf("<itunes:subtitle>%s</itunes:subtitle>", subtitle))
}

func (s *PodcastsTestSuite) TestContainsSummaryElement() {
	summary := "Test Subtitle"
	data, err := getPodcastXML(s.podcast, Summary(summary))

	s.Nil(err)
	s.Contains(data, fmt.Sprintf("<itunes:summary>%s</itunes:summary>", summary))
}

func (s *PodcastsTestSuite) TestContainsOwnerElement() {
	name := "Test Name"
	email := "test@name.com"
	data, err := getPodcastXML(s.podcast, Owner(name, email))

	s.Nil(err)
	s.Contains(data, "<itunes:owner>")
	s.Contains(data, fmt.Sprintf("<itunes:name>%s</itunes:name>", name))
	s.Contains(data, fmt.Sprintf("<itunes:email>%s</itunes:email>", email))
	s.Contains(data, "</itunes:owner>")
}

func (s *PodcastsTestSuite) TestContainsImageElement() {
	image := "http://localhost/myimage.jpg"
	data, err := getPodcastXML(s.podcast, Image(image))

	s.Nil(err)
	s.Contains(data, fmt.Sprintf(`<itunes:image href="%s"></itunes:image>`, image))
}

type testItem struct {
	Item            *Item
	ExpectedPubDate string
}

func (s *PodcastsTestSuite) TestContainsItemElements() {
	for _, item := range validItems {
		s.podcast.AddItem(&Item{
			Title:   item.Title,
			GUID:    item.GUID,
			PubDate: &PubDate{item.Time},
			Enclosure: &Enclosure{
				URL:    item.EnclosureURL,
				Length: item.EnclosureLength,
				Type:   item.EnclosureType,
			},
		})
	}

	feed, err := s.podcast.Feed()
	s.Nil(err)

	data, err := feed.XML()
	s.Nil(err)

	for _, item := range validItems {
		s.Contains(data, "<item>")
		s.Contains(data, fmt.Sprintf("<title>%s</title>", item.Title))
		s.Contains(data, fmt.Sprintf("<guid>%s</guid>", item.GUID))
		s.Contains(data, fmt.Sprintf("<pubDate>%s</pubDate>", item.ExpectedPubDate))
		s.Contains(data, fmt.Sprintf(`<enclosure url="%s" length="%s" type="%s"></enclosure>`, item.EnclosureURL, item.EnclosureLength, item.EnclosureType))
		s.Contains(data, "</item>")
	}
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

func getPodcastXML(p *Podcast, options ...func(f *Feed) error) (string, error) {
	feed, err := p.Feed(options...)
	if err != nil {
		return "", err
	}
	return feed.XML()
}
