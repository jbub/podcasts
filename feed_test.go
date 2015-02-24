package podcasts

import (
	"encoding/xml"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type FeedTestSuite struct {
	suite.Suite
}

func (s *FeedTestSuite) TestPubDateMarshalling() {
	t := time.Date(2015, time.January, 1, 0, 0, 0, 0, time.UTC)
	pd := PubDate{t}
	expected := "<PubDate>Thu, 01 Jan 2015 00:00:00 +0000</PubDate>"
	out, err := xml.Marshal(pd)

	s.Nil(err)
	s.Equal(expected, string(out))
}

func TestFeedTestSuite(t *testing.T) {
	suite.Run(t, new(FeedTestSuite))
}
