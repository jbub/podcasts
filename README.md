podcasts
========

```go
package main

import (
    "log"
    "os"

    "github.com/jbub/podcasts"
    "github.com/jbub/podcasts/itunes"
)

func main() {
    p := &podcasts.Podcast{
        Title:       "Palenica borisa filana",
        Description: "Zábavný program pre každého, komu to páli.",
        Language:    "SK",
        Link:        "http://www.rtvs.sk/radio/relacie/detail/palenica-borisa-filana",
        Copyright:   "2013 RTVS - Rozhlas a televízia Slovenska",
    }

    opts, err := itunes.Options(
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

    p.SetOptions(opts)

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

    p.Feed().Write(os.Stdout)
}

```
