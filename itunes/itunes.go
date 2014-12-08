package itunes

func NewSettings(options ...func(s *Settings) error) (*Settings, error) {
	s := Settings{}
	return &s, s.SetOption(options...)
}

type Category struct {
	Text       string
	Categories []*Category
}

type Settings struct {
	Author     string
	Block      string
	Explicit   string
	Complete   string
	NewFeedURL string
	Subtitle   string
	Summary    string
	Owner      string
	Email      string
	Image      string
	Categories []*Category
}

func (s *Settings) SetOption(options ...func(*Settings) error) error {
	for _, opt := range options {
		if err := opt(s); err != nil {
			return err
		}
	}
	return nil
}

func (s *Settings) AddCategory(c *Category) {
	s.Categories = append(s.Categories, c)
}
