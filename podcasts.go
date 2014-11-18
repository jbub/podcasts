package podcasts

import (
	"encoding/xml"
	"fmt"
	"os"
	"time"
)

type EnclosureType string

const (
	MP3  EnclosureType = "audio/mpeg"
	M4A  EnclosureType = "audio/mpeg"
	MP4  EnclosureType = "video/mp4"
	M4V  EnclosureType = "video/x-m4v"
	MOV  EnclosureType = "video/quicktime"
	PDF  EnclosureType = "application/pdf"
	EPUB EnclosureType = "document/x-epub"
)

type Owner struct {
	XMLName xml.Name `xml:"itunes:owner"`
	Name    string   `xml:"itunes:name"`
	Email   string   `xml:"itunes:email"`
}

type Image struct {
	XMLName xml.Name `xml:"itunes:image"`
	Href    string   `xml:"href,attr"`
}

type Category struct {
	XMLName    xml.Name `xml:"itunes:category"`
	Href       string   `xml:"text,attr"`
	Categories []*Category
}

type Enclosure struct {
	XMLName xml.Name      `xml:"enclosure"`
	URL     string        `xml:"url,attr"`
	Length  string        `xml:"length,attr"`
	Type    EnclosureType `xml:"type,attr"`
}

type Item struct {
	XMLName   xml.Name `xml:"item"`
	Title     string   `xml:"title"`
	Author    string   `xml:"itunes:author"`
	Subtitle  string   `xml:"itunes:subtitle"`
	Summary   string   `xml:"itunes:summary"`
	Image     *Image
	Enclosure *Enclosure
	Guid      string        `xml:"guid"`
	PubDate   time.Time     `xml:"pubdate"`
	Duration  time.Duration `xml:"itunes:duration"`
}

type Channel struct {
	XMLName     xml.Name `xml:"channel"`
	Title       string   `xml:"title"`
	Link        string   `xml:"link"`
	Copyright   string   `xml:"copyright"`
	Subtitle    string   `xml:"itunes:subtitle"`
	Author      string   `xml:"itunes:author"`
	Summary     string   `xml:"itunes:summary"`
	Description string   `xml:"description"`
	Owner       *Owner
	Image       *Image
	Categories  []*Category
	Items       []*Item
}

type Feed struct {
	XMLName xml.Name `xml:"rss"`
	Xmlns   string   `xml:"xmlns:itunes,attr"`
	Version string   `xml:"version,attr"`
	Channel *Channel `xml:"channel"`
}

func main() {
	feed := &Feed{
		Xmlns:   "http://www.itunes.com/dtds/podcast-1.0.dtd",
		Version: "2.0",
		Channel: &Channel{
			Title:     "lopata",
			Link:      "lopata",
			Copyright: "dasdasd",
			Owner: &Owner{
				Name:  "john",
				Email: "dsa@das.sk",
			},
			Image: &Image{Href: "dsas"},
			Categories: []*Category{
				&Category{
					Href: "dsadsa",
					Categories: []*Category{
						&Category{Href: "dsadsa"},
					},
				},
				&Category{
					Href: "hgfhgf",
				},
			},
			Items: []*Item{
				&Item{
					Title: "dsadsa das dsa as",
				},
			},
		},
	}

	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("  ", "    ")
	if err := enc.Encode(feed); err != nil {
		fmt.Printf("error: %v\n", err)
	}
}
