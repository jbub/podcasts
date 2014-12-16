podcasts
========

```go
package main

import (
    "log"
    "os"
    "time"

    "github.com/jbub/podcasts"
)

func main() {
    p := &podcasts.Podcast{
        Title:       "Palenica borisa filana",
        Description: "Zábavný program pre každého, komu to páli.",
        Language:    "SK",
        Link:        "http://www.rtvs.sk/radio/relacie/detail/palenica-borisa-filana",
        Copyright:   "2013 RTVS - Rozhlas a televízia Slovenska",
    }

    p.AddItem(&podcasts.Item{
        Title:   "Epizoda 1",
        GUID:    "http://slovensko.rtvs.sk/clanok/ludia/experti",
        PubDate: &podcasts.PubDate{time.Now()},
        Enclosure: &podcasts.Enclosure{
            URL:    "http://static-media.rtvs.sk/items/223/546de29065c77.mp3",
            Length: "321",
            Type:   "MP3",
        },
    })

    feed := p.Feed()

    err := feed.SetOptions(
        podcasts.ItunesAuthor("Boris Filan"),
        podcasts.ItunesBlock,
        podcasts.ItunesExplicit,
        podcasts.ItunesComplete,
        podcasts.ItunesNewFeedURL("http://www.rtvs.sk/radio/relacie/detail/palenica-borisa-filana"),
        podcasts.ItunesSubtitle("Zábavný program pre každého, komu to páli."),
        podcasts.ItunesSummary("Zábavný program pre každého, komu to páli."),
        podcasts.ItunesOwner("Rozhlas a televízia Slovenska", "vsv@rtvs.sk"),
        podcasts.ItunesImage("http://cdn.srv.rtvs.sk/a501/image/file/13/0006/wRe0.filan_boris_700.jpg"),
    )

    if err != nil {
        log.Fatal(err)
    }

    feed.Write(os.Stdout)
}

```
