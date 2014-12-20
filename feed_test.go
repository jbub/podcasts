package podcasts

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type FeedTestSuite struct {
	suite.Suite
}

func (s *FeedTestSuite) TestFeed() {

}

func TestFeedTestSuite(t *testing.T) {
	suite.Run(t, new(FeedTestSuite))
}
