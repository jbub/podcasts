package podcasts

import (
	"testing"
)

func TestAuthor(t *testing.T) {
	feed := &Feed{
		Channel: &Channel{},
	}
	author := "john"
	if err := Author(author)(feed); err != nil {
		t.Errorf("unexpected error %v", err)
	}
	if author != feed.Channel.Author {
		t.Errorf("expected %v got %v", author, feed.Channel.Author)
	}
}

func TestBlock(t *testing.T) {
	feed := &Feed{
		Channel: &Channel{},
	}
	if err := Block(feed); err != nil {
		t.Errorf("unexpected error %v", err)
	}
	if ValueYes != feed.Channel.Block {
		t.Errorf("expected %v got %v", ValueYes, feed.Channel.Block)
	}
}

func TestExplicit(t *testing.T) {
	feed := &Feed{
		Channel: &Channel{},
	}
	if err := Explicit(feed); err != nil {
		t.Errorf("unexpected error %v", err)
	}
	if ValueYes != feed.Channel.Explicit {
		t.Errorf("expected %v got %v", ValueYes, feed.Channel.Explicit)
	}
}

func TestComplete(t *testing.T) {
	feed := &Feed{
		Channel: &Channel{},
	}
	if err := Complete(feed); err != nil {
		t.Errorf("unexpected error %v", err)
	}
	if ValueYes != feed.Channel.Complete {
		t.Errorf("expected %v got %v", ValueYes, feed.Channel.Complete)
	}
}

func TestNewFeedURL(t *testing.T) {
	feed := &Feed{
		Channel: &Channel{},
	}
	newURL := "http://example.com/test"
	if err := NewFeedURL(newURL)(feed); err != nil {
		t.Errorf("unexpected error %v", err)
	}
	if newURL != feed.Channel.NewFeedURL {
		t.Errorf("expected %v got %v", ValueYes, feed.Channel.NewFeedURL)
	}
}

func TestNewFeedURLInvalid(t *testing.T) {
	feed := &Feed{
		Channel: &Channel{},
	}
	newURL := "invalid url"
	if err := NewFeedURL(newURL)(feed); err == nil {
		t.Error("expected error")
	}
}

func TestSubtitle(t *testing.T) {
	feed := &Feed{
		Channel: &Channel{},
	}
	subtitle := "this is subtitle"
	if err := Subtitle(subtitle)(feed); err != nil {
		t.Errorf("unexpected error %v", err)
	}
	if subtitle != feed.Channel.Subtitle {
		t.Errorf("expected %v got %v", subtitle, feed.Channel.Subtitle)
	}
}

func TestSummary(t *testing.T) {
	feed := &Feed{
		Channel: &Channel{},
	}
	summary := `this is summary. <a href="http://example.com/more">more</a>`
	if err := Summary(summary)(feed); err != nil {
		t.Errorf("unexpected error %v", err)
	}
	if summary != feed.Channel.Summary.Value {
		t.Errorf("expected %v got %v", summary, feed.Channel.Summary.Value)
	}
}

func TestOwner(t *testing.T) {
	feed := &Feed{
		Channel: &Channel{},
	}
	name := "anabelle"
	email := "test@test.com"
	if err := Owner(name, email)(feed); err != nil {
		t.Errorf("unexpected error %v", err)
	}
	if name != feed.Channel.Owner.Name {
		t.Errorf("expected %v got %v", name, feed.Channel.Owner.Name)
	}
	if email != feed.Channel.Owner.Email {
		t.Errorf("expected %v got %v", email, feed.Channel.Owner.Email)
	}
}

func TestImage(t *testing.T) {
	feed := &Feed{
		Channel: &Channel{},
	}
	href := "http://example.com/test/image.jpg"
	if err := Image(href)(feed); err != nil {
		t.Errorf("unexpected error %v", err)
	}
	if href != feed.Channel.Image.Href {
		t.Errorf("expected %v got %v", href, feed.Channel.Image.Href)
	}
}

func TestImageInvalid(t *testing.T) {
	feed := &Feed{
		Channel: &Channel{},
	}
	href := "invalid img url"
	if err := Image(href)(feed); err == nil {
		t.Errorf("expected error")
	}
}
