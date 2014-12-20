podcasts
========

Podcast generator written in Go.

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

    feed, err := p.Feed(
        podcasts.Author("Boris Filan"),
        podcasts.Block,
        podcasts.Explicit,
        podcasts.Complete,
        podcasts.NewFeedURL("http://www.rtvs.sk/radio/relacie/detail/palenica-borisa-filana"),
        podcasts.Subtitle("Zábavný program pre každého, komu to páli."),
        podcasts.Summary("Zábavný program pre každého, komu to páli."),
        podcasts.Owner("Rozhlas a televízia Slovenska", "vsv@rtvs.sk"),
        podcasts.Image("http://cdn.srv.rtvs.sk/a501/image/file/13/0006/wRe0.filan_boris_700.jpg"),
    )

    if err != nil {
        log.Fatal(err)
    }

    feed.Write(os.Stdout)
}

```
