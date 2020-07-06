package podcasts

import (
	"bytes"
	"encoding/xml"
	"io"
	"strconv"
	"strings"
	"time"
)

const (
	rssXmlns   = "http://www.itunes.com/dtds/podcast-1.0.dtd"
	rssVersion = "2.0"
	rfc2822    = "Mon, 02 Jan 2006 15:04:05 -0700"
)

// NewPubDate returns a new PubDate.
func NewPubDate(d time.Time) *PubDate {
	return &PubDate{d}
}

// PubDate represents pubDate attribute of given podcast item.
type PubDate struct {
	time.Time
}

// MarshalXML marshalls pubdate using the rfc2822 time format.
func (p PubDate) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeToken(start); err != nil {
		return err
	}
	if err := e.EncodeToken(xml.CharData(p.Format(rfc2822))); err != nil {
		return err
	}
	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

// NewDuration returns a new Duration.
func NewDuration(d time.Duration) *Duration {
	return &Duration{d}
}

// Duration represents itunes:duration attribute of given podcast item.
type Duration struct {
	time.Duration
}

// MarshalXML marshalls duration using HH:MM:SS, H:MM:SS, MM:SS, M:SS formats.
func (d Duration) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeToken(start); err != nil {
		return err
	}
	if err := e.EncodeToken(xml.CharData(formatDuration(d.Duration))); err != nil {
		return err
	}
	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

// formatDuration formats duration in these formats: HH:MM:SS, H:MM:SS, MM:SS, M:SS.
func formatDuration(d time.Duration) string {
	total := int(d.Seconds())
	hours := total / 3600
	total %= 3600
	minutes := total / 60
	total %= 60

	var b strings.Builder
	if hours > 0 {
		b.WriteString(strconv.Itoa(hours) + ":")
	}
	if hours > 0 && minutes < 10 {
		b.WriteString("0")
	}
	b.WriteString(strconv.Itoa(minutes) + ":")
	if total < 10 {
		b.WriteString("0")
	}
	b.WriteString(strconv.Itoa(total))
	return b.String()
}

// ItunesOwner represents the itunes:owner of given channel.
type ItunesOwner struct {
	XMLName xml.Name `xml:"itunes:owner"`
	Name    string   `xml:"itunes:name"`
	Email   string   `xml:"itunes:email"`
}

// ItunesImage represents the itunes:image of given item or channel.
type ItunesImage struct {
	XMLName xml.Name `xml:"itunes:image"`
	Href    string   `xml:"href,attr"`
}

// ItunesCategory represents itunes:category of given channel.
type ItunesCategory struct {
	XMLName    xml.Name `xml:"itunes:category"`
	Text       string   `xml:"text,attr"`
	Categories []*ItunesCategory
}

// Enclosure represents audio or video file of given item.
type Enclosure struct {
	XMLName xml.Name `xml:"enclosure"`
	URL     string   `xml:"url,attr"`
	Length  string   `xml:"length,attr,omitempty"`
	Type    string   `xml:"type,attr"`
}

// ItunesSummary represents a summary of a Channel or Item that may contain
// embedded HTML such as <a href="...">...</a> links.
type ItunesSummary struct {
	Value string `xml:",cdata"`
}

// Item represents item of given channel.
type Item struct {
	XMLName         xml.Name       `xml:"item"`
	Title           string         `xml:"title"`
	GUID            string         `xml:"guid"`
	PubDate         *PubDate       `xml:"pubDate"`
	Author          string         `xml:"itunes:author,omitempty"`
	Block           string         `xml:"itunes:block,omitempty"`
	Duration        *Duration      `xml:"itunes:duration,omitempty"`
	Explicit        string         `xml:"itunes:explicit,omitempty"`
	ClosedCaptioned string         `xml:"itunes:isClosedCaptioned,omitempty"`
	Order           int            `xml:"itunes:order,omitempty"`
	Subtitle        string         `xml:"itunes:subtitle,omitempty"`
	Summary         *ItunesSummary `xml:"itunes:summary,omitempty"`
	Enclosure       *Enclosure
	Image           *ItunesImage
}

// Channel represents a RSS channel for given podcast.
type Channel struct {
	XMLName     xml.Name       `xml:"channel"`
	Title       string         `xml:"title"`
	Link        string         `xml:"link"`
	Copyright   string         `xml:"copyright"`
	Language    string         `xml:"language"`
	Description string         `xml:"description"`
	Author      string         `xml:"itunes:author,omitempty"`
	Block       string         `xml:"itunes:block,omitempty"`
	Explicit    string         `xml:"itunes:explicit,omitempty"`
	Complete    string         `xml:"itunes:complete,omitempty"`
	NewFeedURL  string         `xml:"itunes:new-feed-url,omitempty"`
	Subtitle    string         `xml:"itunes:subtitle,omitempty"`
	Summary     *ItunesSummary `xml:"itunes:summary,omitempty"`
	Owner       *ItunesOwner
	Image       *ItunesImage
	Items       []*Item
	Categories  []*ItunesCategory
}

// Feed wraps the given RSS channel.
type Feed struct {
	XMLName xml.Name `xml:"rss"`
	Xmlns   string   `xml:"xmlns:itunes,attr"`
	Version string   `xml:"version,attr"`
	Channel *Channel
}

// SetOptions sets options of given feed.
func (f *Feed) SetOptions(options ...func(f *Feed) error) error {
	for _, opt := range options {
		if err := opt(f); err != nil {
			return err
		}
	}
	return nil
}

// XML marshalls feed to XML string.
func (f *Feed) XML() (string, error) {
	var buf bytes.Buffer
	if err := f.Write(&buf); err != nil {
		return "", err
	}
	return buf.String(), nil
}

// Write writes marshalled XML to the given writer.
func (f *Feed) Write(w io.Writer) error {
	if _, err := w.Write([]byte(xml.Header)); err != nil {
		return err
	}
	enc := xml.NewEncoder(w)
	enc.Indent("", "  ")
	return enc.Encode(f)
}
