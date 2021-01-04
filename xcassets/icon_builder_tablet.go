package xcassets

// AppIconTablet contains confifuration and customization options for the
// iPad icon.
type AppIconTablet struct {
	Source       AssetSource
	Notification AssetSource
	Settings     AssetSource
	Spotlight    AssetSource
	App          AssetSource
	Pro          AssetSource
	enabled      bool
}

// Configure allows you to override the default source and configuration.
func (b *AppIconTablet) Configure(f func(*AppIconTablet)) {
	f(b)
}

// Validate will validate the icon with the provided parent asset source.
func (b *AppIconTablet) Validate(s AssetSource) error {
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

	if b.Pro.Empty() {
		b.Pro.Apply(b.Source)
	}

	return nil
}

// Build will validate and build the icon using the provided parent asset
// source. The parent asset source will only be used if no source has been
// specified for this icon.
func (b *AppIconTablet) Build(name string, s AssetSource) ([]AppIconImageInput, error) {
	if err := b.Validate(s); err != nil {
		return nil, err
	}

	builder := IconImageBuilder{}

	images := []AppIconImageInput{
		// iPad Notifications iOS 7-14
		builder.buildInput(name, "ipad", 1, 20, b.Notification),
		builder.buildInput(name, "ipad", 2, 20, b.Notification),

		// iPad Settings iOS 7-14
		builder.buildInput(name, "ipad", 1, 29, b.Settings),
		builder.buildInput(name, "ipad", 2, 29, b.Settings),

		// iPad Spotlight iOS 7-14
		builder.buildInput(name, "ipad", 1, 40, b.Spotlight),
		builder.buildInput(name, "ipad", 2, 40, b.Spotlight),

		// iPad App Icon iOS 7-14
		builder.buildInput(name, "ipad", 1, 76, b.App),
		builder.buildInput(name, "ipad", 2, 76, b.App),

		// iPad Pro App Icon iOS 9-14
		builder.buildInput(name, "ipad", 2, 83.5, b.Pro),
	}

	return images, nil
}
