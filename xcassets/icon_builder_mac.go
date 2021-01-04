package xcassets

// AppIconMac contains confifuration and customization options for the
// Mac icon.
type AppIconMac struct {
	Source  AssetSource
	Size16  AssetSource
	Size32  AssetSource
	Size128 AssetSource
	Size256 AssetSource
	Size512 AssetSource
	enabled bool
}

// Configure allows you to override the default source and configuration.
func (b *AppIconMac) Configure(f func(*AppIconMac)) {
	f(b)
}

// Validate will validate the icon with the provided parent asset source.
func (b *AppIconMac) Validate(s AssetSource) error {
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

	if b.Size16.Empty() {
		b.Size16.Apply(b.Source)
	}

	if b.Size32.Empty() {
		b.Size32.Apply(b.Source)
	}

	if b.Size128.Empty() {
		b.Size128.Apply(b.Source)
	}

	if b.Size256.Empty() {
		b.Size256.Apply(b.Source)
	}

	if b.Size512.Empty() {
		b.Size512.Apply(b.Source)
	}

	return nil
}

// Build will validate and build the icon using the provided parent asset
// source. The parent asset source will only be used if no source has been
// specified for this icon.
func (b *AppIconMac) Build(name string, s AssetSource) ([]AppIconImageInput, error) {
	if err := b.Validate(s); err != nil {
		return nil, err
	}

	builder := IconImageBuilder{}

	images := []AppIconImageInput{
		builder.buildInput(name, "mac", 1, 16, b.Size16),
		builder.buildInput(name, "mac", 2, 16, b.Size16),

		builder.buildInput(name, "mac", 1, 32, b.Size32),
		builder.buildInput(name, "mac", 2, 32, b.Size32),

		builder.buildInput(name, "mac", 1, 128, b.Size128),
		builder.buildInput(name, "mac", 2, 128, b.Size128),

		builder.buildInput(name, "mac", 1, 256, b.Size256),
		builder.buildInput(name, "mac", 2, 256, b.Size256),

		builder.buildInput(name, "mac", 1, 512, b.Size512),
		builder.buildInput(name, "mac", 2, 512, b.Size512),
	}

	return images, nil
}
