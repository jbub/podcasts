package podcasts

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/jbub/podcasts/itunes"
	"github.com/stretchr/testify/suite"
)

type PodcastsTestSuite struct {
	suite.Suite
}

func (s *PodcastsTestSuite) TestPodcast() {
	itunes, err := itunes.New(
		itunes.Explicit("yes"),
	)
	if err != nil {
		log.Fatal(err)
	}
	p := &Podcast{
		Title:       "Palenica borisa filana",
		Description: "Zábavný program pre každého, komu to páli.",
		Language:    "SK",
		Link:        "http://www.rtvs.sk/radio/relacie/detail/palenica-borisa-filana",
		Copyright:   "2013 RTVS - Rozhlas a televízia Slovenska",
	}
	p.SetItunes(itunes)
	p.AddItem(&Item{
		Title:   "Epizoda 1",
		GUID:    "http://slovensko.rtvs.sk/clanok/ludia/experti",
		PubDate: &PubDate{time.Now()},
		Enclosure: &Enclosure{
			URL:    "http://static-media.rtvs.sk/items/223/546de29065c77.mp3",
			Length: "321",
			Type:   "MP3",
		},
	})
	p.Feed().Write(os.Stdout)
}

func TestPodcastsTestSuite(t *testing.T) {
	suite.Run(t, new(PodcastsTestSuite))
}
