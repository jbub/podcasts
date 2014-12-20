package podcasts

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type FeedTestSuite struct {
	suite.Suite
	feed *Feed
}

func (s *FeedTestSuite) SetupTest() {
	s.feed = &Feed{
		Xmlns:   rssXmlns,
		Version: rssVersion,
		Channel: &Channel{
			Title:       p.Title,
			Description: p.Description,
			Link:        p.Link,
			Copyright:   p.Copyright,
			Language:    p.Language,
			Items:       p.items,
		},
	}
}

func (s *FeedTestSuite) TestFeed() {

}

func TestFeedTestSuite(t *testing.T) {
	suite.Run(t, new(FeedTestSuite))
}
