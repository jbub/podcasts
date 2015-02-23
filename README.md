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
    // initialize the podcast
    p := &podcasts.Podcast{
        Title:       "My podcast",
        Description: "This is my very simple podcast.",
        Language:    "EN",
        Link:        "http://www.example-podcast.com/my-podcast",
        Copyright:   "2015 My podcast copyright",
    }

    // add first podcast item
    p.AddItem(&podcasts.Item{
        Title:   "Episode 1",
        GUID:    "http://www.example-podcast.com/my-podcast/1/episode-one",
        PubDate: &podcasts.PubDate{time.Now()},
        Enclosure: &podcasts.Enclosure{
            URL:    "http://www.example-podcast.com/my-podcast/1/episode.mp3",
            Length: "12312",
            Type:   "MP3",
        },
    })

    // add second podcast item
    p.AddItem(&podcasts.Item{
        Title:   "Episode 2",
        GUID:    "http://www.example-podcast.com/my-podcast/2/episode-two",
        PubDate: &podcasts.PubDate{time.Now()},
        Enclosure: &podcasts.Enclosure{
            URL:    "http://www.example-podcast.com/my-podcast/2/episode.mp3",
            Length: "46732",
            Type:   "MP3",
        },
    })

    // get podcast feed, you can pass options to customize it
    feed, err := p.Feed(
        podcasts.Author("Author Name"),
        podcasts.Block,
        podcasts.Explicit,
        podcasts.Complete,
        podcasts.NewFeedURL("http://www.example-podcast.com/new-feed-url"),
        podcasts.Subtitle("This is my very simple podcast subtitle."),
        podcasts.Summary("This is my very simple podcast summary."),
        podcasts.Owner("Podcast Owner", "owner@example-podcast.com"),
        podcasts.Image("http://www.example-podcast.com/my-podcast.jpg"),
    )

    if err != nil {
        log.Fatal(err)
    }

    // finally write the xml
    feed.Write(os.Stdout)
}
```
