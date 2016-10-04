package podcasts

import (
	"encoding/xml"
	"testing"
	"time"
)

func TestPubDateMarshalling(t *testing.T) {
	tm := time.Date(2015, time.January, 1, 0, 0, 0, 0, time.UTC)
	pubDate := PubDate{tm}
	want := "<PubDate>Thu, 01 Jan 2015 00:00:00 +0000</PubDate>"
	out, err := xml.Marshal(pubDate)
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}
	if got := string(out); want != got {
		t.Errorf("expected %v got %v", want, got)
	}
}
