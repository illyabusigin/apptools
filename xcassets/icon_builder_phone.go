package xcassets

// AppIconPhone contains confifuration and customization options for the
// iPhone icon.
type AppIconPhone struct {
	Source       AssetSource
	Notification AssetSource
	Spotlight    AssetSource
	Settings     AssetSource
	App          AssetSource
	enabled      bool
}

// Configure allows you to override the default source and configuration.
func (b *AppIconPhone) Configure(f func(*AppIconPhone)) {
	f(b)
}

// Validate will validate the icon with the provided parent asset source.
func (b *AppIconPhone) Validate(s AssetSource) error {
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

	if b.Notification.Empty() {
		b.Notification.Apply(b.Source)
	}

	if b.Spotlight.Empty() {
		b.Spotlight.Apply(b.Source)
	}

	if b.Settings.Empty() {
		b.Settings.Apply(b.Source)
	}

	if b.Settings.Empty() {
		b.Settings.Apply(b.Source)
	}

	if b.App.Empty() {
		b.App.Apply(b.Source)
	}

	return nil
}

// Build will validate and build the icon using the provided parent asset
// source. The parent asset source will only be used if no source has been
// specified for this icon.
func (b *AppIconPhone) Build(name string, s AssetSource) ([]AppIconImageInput, error) {
	if err := b.Validate(s); err != nil {
		return nil, err
	}

	builder := IconImageBuilder{}

	images := []AppIconImageInput{
		// iPhone Notifications iOS 7-14
		builder.buildInput(name, "iphone", 2, 20, b.Notification),
		builder.buildInput(name, "iphone", 3, 20, b.Notification),

		// iPhone iOS 7-14
		builder.buildInput(name, "iphone", 2, 29, b.Settings),
		builder.buildInput(name, "iphone", 3, 29, b.Settings),

		// iPhone Spotlight iOS 7-14
		builder.buildInput(name, "iphone", 2, 40, b.Spotlight),
		builder.buildInput(name, "iphone", 3, 40, b.Spotlight),

		// iPhone App Icon iOS 7-14
		builder.buildInput(name, "iphone", 2, 60, b.App),
		builder.buildInput(name, "iphone", 3, 60, b.App),
	}

	return images, nil
}
