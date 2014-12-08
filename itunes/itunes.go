package itunes

func Options(options ...func(i *ChannelOpts) error) (*ChannelOpts, error) {
	opts := ChannelOpts{}
	return &opts, opts.SetOption(options...)
}

type Category struct {
	Text       string
	Categories []*Category
}

type ChannelOpts struct {
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

func (i *ChannelOpts) SetOption(options ...func(*ChannelOpts) error) error {
	for _, opt := range options {
		if err := opt(i); err != nil {
			return err
		}
	}
	return nil
}

func (i *ChannelOpts) AddCategory(c *Category) {
	i.Categories = append(i.Categories, c)
}
