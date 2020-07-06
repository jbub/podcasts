package podcasts

import (
	"encoding/xml"
	"testing"
	"time"
)

func TestPubDateMarshalling(t *testing.T) {
	tm := time.Date(2015, time.January, 1, 0, 0, 0, 0, time.UTC)
	pubDate := NewPubDate(tm)
	want := "<PubDate>Thu, 01 Jan 2015 00:00:00 +0000</PubDate>"
	out, err := xml.Marshal(pubDate)
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}
	if got := string(out); want != got {
		t.Errorf("expected %v got %v", want, got)
	}
}

func TestDurationMarshalling(t *testing.T) {
	cases := []struct {
		dur  time.Duration
		want string
	}{
		{
			dur:  0,
			want: "<Duration>0:00</Duration>",
		},
		{
			dur:  time.Second * 6,
			want: "<Duration>0:06</Duration>",
		},
		{
			dur:  time.Second * 64,
			want: "<Duration>1:04</Duration>",
		},
		{
			dur:  time.Second * 125,
			want: "<Duration>2:05</Duration>",
		},
		{
			dur:  time.Second * 3600,
			want: "<Duration>1:00:00</Duration>",
		},
		{
			dur:  time.Second * 37000,
			want: "<Duration>10:16:40</Duration>",
		},
	}

	for _, cs := range cases {
		t.Run(cs.want, func(t *testing.T) {
			dur := NewDuration(cs.dur)
			out, err := xml.Marshal(dur)
			if err != nil {
				t.Errorf("unexpected error %v", err)
			}
			if got := string(out); cs.want != got {
				t.Errorf("expected %v got %v", cs.want, got)
			}
		})
	}
}
