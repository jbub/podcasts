package itunes

func New(options ...func(i *Options) error) (*Options, error) {
	opts := &Options{}
	for _, option := range options {
		err := option(opts)
		if err != nil {
			return nil, err
		}
	}
	return opts, nil
}

type Category struct {
	Text       string
	Categories []*Category
}

type Options struct {
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

func (i *Options) AddCategory(c *Category) {
	i.Categories = append(i.Categories, c)
}

func Author(author string) func(i *Options) error {
	return func(i *Options) error {
		i.Author = author
		return nil
	}
}

func Explicit(explicit string) func(i *Options) error {
	return func(i *Options) error {
		i.Explicit = explicit
		return nil
	}
}

func NewFeedURL(url string) func(i *Options) error {
	return func(i *Options) error {
		i.NewFeedURL = url
		return nil
	}
}

func Subtitle(subtitle string) func(i *Options) error {
	return func(i *Options) error {
		i.Subtitle = subtitle
		return nil
	}
}

func Summary(summary string) func(i *Options) error {
	return func(i *Options) error {
		i.Summary = summary
		return nil
	}
}

func Complete(complete string) func(i *Options) error {
	return func(i *Options) error {
		i.Complete = complete
		return nil
	}
}

func Owner(name string, email string) func(i *Options) error {
	return func(i *Options) error {
		i.Owner = name
		i.Email = email
		return nil
	}
}

func Image(href string) func(i *Options) error {
	return func(i *Options) error {
		i.Image = href
		return nil
	}
}
