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
	p := &Podcast{
		Title:       "Palenica borisa filana",
		Description: "Zábavný program pre každého, komu to páli.",
		Language:    "SK",
		Link:        "http://www.rtvs.sk/radio/relacie/detail/palenica-borisa-filana",
		Copyright:   "2013 RTVS - Rozhlas a televízia Slovenska",
	}

	settings, err := itunes.NewSettings(
		itunes.Author("Boris Filan"),
		itunes.Block,
		itunes.Explicit,
		itunes.Complete,
		itunes.NewFeedURL("http://www.rtvs.sk/radio/relacie/detail/palenica-borisa-filana"),
		itunes.Subtitle("Zábavný program pre každého, komu to páli."),
		itunes.Summary("Zábavný program pre každého, komu to páli."),
		itunes.Owner("Rozhlas a televízia Slovenska", "vsv@rtvs.sk"),
		itunes.Image("http://cdn.srv.rtvs.sk/a501/image/file/13/0006/wRe0.filan_boris_700.jpg"),
	)
	if err != nil {
		log.Fatal(err)
	}
	p.SetOptions(settings)

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
