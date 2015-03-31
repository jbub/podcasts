# podcasts [![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/jbub/podcasts) [![Build Status](http://img.shields.io/travis/jbub/podcasts.svg?style=flat-square)](https://travis-ci.org/jbub/podcasts) [![Coverage Status](http://img.shields.io/coveralls/jbub/podcasts.svg?style=flat-square)](https://coveralls.io/r/jbub/podcasts)

Podcast generator written in Go.

## Install

```bash
go get github.com/jbub/podcasts
```

## Docs

http://godoc.org/github.com/jbub/podcasts

## Example usage

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

    // handle error
    if err != nil {
        log.Fatal(err)
    }

    // finally write the xml to any io.Writer
    feed.Write(os.Stdout)
}
```

Which gives us this XML output:

```xml
<?xml version="1.0" encoding="UTF-8"?>
<rss xmlns:itunes="http://www.itunes.com/dtds/podcast-1.0.dtd" version="2.0">
  <channel>
    <title>My podcast</title>
    <link>http://www.example-podcast.com/my-podcast</link>
    <copyright>2015 My podcast copyright</copyright>
    <language>EN</language>
    <description>This is my very simple podcast.</description>
    <itunes:author>Author Name</itunes:author>
    <itunes:block>yes</itunes:block>
    <itunes:explicit>yes</itunes:explicit>
    <itunes:complete>yes</itunes:complete>
    <itunes:new-feed-url>http://www.example-podcast.com/new-feed-url</itunes:new-feed-url>
    <itunes:subtitle>This is my very simple podcast subtitle.</itunes:subtitle>
    <itunes:summary>This is my very simple podcast summary.</itunes:summary>
    <itunes:owner>
      <itunes:name>Podcast Owner</itunes:name>
      <itunes:email>owner@example-podcast.com</itunes:email>
    </itunes:owner>
    <itunes:image href="http://www.example-podcast.com/my-podcast.jpg"></itunes:image>
    <item>
      <title>Episode 1</title>
      <guid>http://www.example-podcast.com/my-podcast/1/episode-one</guid>
      <pubDate>Wed, 25 Feb 2015 10:43:17 +0100</pubDate>
      <enclosure url="http://www.example-podcast.com/my-podcast/1/episode.mp3" length="12312" type="MP3"></enclosure>
    </item>
    <item>
      <title>Episode 2</title>
      <guid>http://www.example-podcast.com/my-podcast/2/episode-two</guid>
      <pubDate>Wed, 25 Feb 2015 10:43:17 +0100</pubDate>
      <enclosure url="http://www.example-podcast.com/my-podcast/2/episode.mp3" length="46732" type="MP3"></enclosure>
    </item>
  </channel>
</rss>
```