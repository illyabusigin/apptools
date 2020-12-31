package xcassets

type AppIconAppStore struct {
	Source AppIconSource

	enabled bool
}

func (b *AppIconAppStore) Configure(f func(*AppIconAppStore)) {
	f(b)
}

func (b *AppIconAppStore) Validate(s AppIconSource) error {
	if !b.enabled {
		return nil
	}

	b.Source.minDimension = 1024
	if b.Source.Empty() {
		b.Source.Apply(s)
	}

	if err := b.Source.Validate(); err != nil {
		return err
	}

	return nil
}

func (b *AppIconAppStore) Build(name string, s AppIconSource) ([]AppIconImageInput, error) {
	if err := b.Validate(s); err != nil {
		return nil, err
	}

	builder := IconImageBuilder{}

	images := []AppIconImageInput{
		builder.buildInput(name, "ios-marketing", 1, 1024, b.Source),
	}

	return images, nil
}
