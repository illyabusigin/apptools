package xcassets

// AppIconAppStore contains confifuration and customization options for the
// app store marketing icon.
type AppIconAppStore struct {
	Source AssetSource

	enabled bool
}

// Configure allows you to override the default source and configuration.
func (b *AppIconAppStore) Configure(f func(*AppIconAppStore)) {
	f(b)
}

// Validate will validate the icon with the provided parent asset source.
func (b *AppIconAppStore) Validate(s AssetSource) error {
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

// Build will validate and build the icon using the provided parent asset
// source. The parent asset source will only be used if no source has been
// specified for this icon.
func (b *AppIconAppStore) Build(name string, s AssetSource) ([]AppIconImageInput, error) {
	if err := b.Validate(s); err != nil {
		return nil, err
	}

	builder := IconImageBuilder{}

	images := []AppIconImageInput{
		builder.buildInput(name, "ios-marketing", 1, 1024, b.Source),
	}

	return images, nil
}
