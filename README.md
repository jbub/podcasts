go-podcasts
===========


```go
itunes := &podcasts.Itunes{
    Author: "jbub",
    Block: "yes",
    Explicit: "clean",
    Complete: "yes",
    NewFeedUrl: "http://example.com/feed/url",
    Subtitle: "random subtitle",
    Summary: "random summary",
    Owner: "jbub",
    Email: "email@example.com",
    Image: "http://example.com/hello.jpg",
}

podcast := &podcasts.Podcast{
	Title:       "Palenica borisa filana",
	Description: "Zábavný program pre každého, komu to páli.",
	Language:    "SK",
	Link:        "http://www.rtvs.sk/radio/relacie/detail/palenica-borisa-filana",
	Copyright:   "2013 RTVS - Rozhlas a televízia Slovenska",
	Itunes:      itunes,
}

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
```
