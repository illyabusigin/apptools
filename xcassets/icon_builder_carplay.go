package xcassets

// AppIconCarPlay contains confifuration and customization options for the
// CarPlay icon.
type AppIconCarPlay struct {
	Source AssetSource
	Icon   AssetSource

	enabled bool
}

// Configure allows you to override the default source and configuration.
func (b *AppIconCarPlay) Configure(f func(*AppIconCarPlay)) {
	f(b)
}

// Validate will validate the icon with the provided parent asset source.
func (b *AppIconCarPlay) Validate(s AssetSource) error {
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

	if b.Icon.Empty() {
		b.Icon.Apply(b.Source)
	}

	return nil
}

// Build will validate and build the icon using the provided parent asset
// source. The parent asset source will only be used if no source has been
// specified for this icon.
func (b *AppIconCarPlay) Build(name string, s AssetSource) ([]AppIconImageInput, error) {
	if err := b.Validate(s); err != nil {
		return nil, err
	}

	builder := IconImageBuilder{}

	images := []AppIconImageInput{
		builder.buildInput(name, "car", 2, 60, b.Icon),
		builder.buildInput(name, "car", 3, 60, b.Icon),
	}

	return images, nil
}
