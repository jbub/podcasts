package podcasts

import (
	"encoding/xml"
	"io"
	"time"
)

const (
	rssXmlns   = "http://www.itunes.com/dtds/podcast-1.0.dtd"
	rssVersion = "2.0"
)

type PubDate struct {
	time.Time
}

func (p PubDate) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	e.EncodeToken(xml.CharData([]byte(p.Format("Mon Jan 02 15:04:05 -0700 2006"))))
	e.EncodeToken(xml.EndElement{start.Name})
	return nil
}

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
	XMLName xml.Name `xml:"enclosure"`
	URL     string   `xml:"url,attr"`
	Length  string   `xml:"length,attr"`
	Type    string   `xml:"type,attr"`
}

type Item struct {
	XMLName         xml.Name      `xml:"item"`
	Title           string        `xml:"title"`
	GUID            string        `xml:"guid"`
	PubDate         *PubDate      `xml:"pubdate"`
	Author          string        `xml:"itunes:author,omitempty""`
	Subtitle        string        `xml:"itunes:subtitle,omitempty""`
	Summary         string        `xml:"itunes:summary,omitempty""`
	Explicit        string        `xml:"itunes:explicit,omitempty"`
	Block           string        `xml:"itunes:block,omitempty"`
	ClosedCaptioned string        `xml:"itunes:isClosedCaptioned,omitempty"`
	Order           int           `xml:"itunes:order,omitempty"`
	Duration        time.Duration `xml:"itunes:duration,omitempty""`
	Image           *Image
	Enclosure       *Enclosure
}

type Channel struct {
	XMLName     xml.Name `xml:"channel"`
	Title       string   `xml:"title"`
	Link        string   `xml:"link"`
	Language    string   `xml:"language"`
	Description string   `xml:"description"`
	Copyright   string   `xml:"copyright"`
	Subtitle    string   `xml:"itunes:subtitle,omitempty""`
	Author      string   `xml:"itunes:author,omitempty""`
	Summary     string   `xml:"itunes:summary,omitempty""`
	Block       string   `xml:"itunes:block,omitempty"`
	Complete    string   `xml:"itunes:complete,omitempty"`
	NewFeedURL  string   `xml:"itunes:new-feed-url,omitempty"`
	Owner       *Owner
	Image       *Image
	Categories  []*Category
	Items       []*Item
}

type Feed struct {
	XMLName xml.Name `xml:"rss"`
	Xmlns   string   `xml:"xmlns:itunes,attr"`
	Version string   `xml:"version,attr"`
	Channel *Channel
}

func (f *Feed) Xml() (string, error) {
	data, err := xml.MarshalIndent(f, "", "  ")
	if err != nil {
		return "", err
	}
	s := xml.Header + string(data)
	return s, nil
}

func (f *Feed) Write(w io.Writer) error {
	if _, err := w.Write([]byte(xml.Header)); err != nil {
		return err
	}
	enc := xml.NewEncoder(w)
	enc.Indent("", "  ")
	return enc.Encode(f)
}
