package xcassets

type AppIconWatch struct {
	Source       AssetSource
	Notification AssetSource
	Settings     AssetSource
	HomeScreen   AssetSource
	ShortLook    AssetSource
	enabled      bool
}

func (b *AppIconWatch) Configure(f func(*AppIconWatch)) {
	f(b)
}

func (b *AppIconWatch) Validate(s AssetSource) error {
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

	if b.Settings.Empty() {
		b.Settings.Apply(b.Source)
	}

	if b.Settings.Empty() {
		b.Settings.Apply(b.Source)
	}

	if b.HomeScreen.Empty() {
		b.HomeScreen.Apply(b.Source)
	}

	if b.ShortLook.Empty() {
		b.ShortLook.Apply(b.Source)
	}

	return nil
}

func (b *AppIconWatch) Build(name string, s AssetSource) ([]AppIconImageInput, error) {
	if err := b.Validate(s); err != nil {
		return nil, err
	}

	builder := IconImageBuilder{}

	images := []AppIconImageInput{
		// Watch Notifications
		builder.buildExtendedInput(name, "watch", "notificationCenter", "38mm", 2, 24, b.Notification),
		builder.buildExtendedInput(name, "watch", "notificationCenter", "42mm", 2, 27.5, b.Notification),

		// Watch Companion Settings
		builder.buildExtendedInput(name, "watch", "companionSettings", "", 2, 29, b.Settings),
		builder.buildExtendedInput(name, "watch", "companionSettings", "", 3, 29, b.Settings),

		// Watch Home Screen
		builder.buildExtendedInput(name, "watch", "appLauncher", "38mm", 2, 40, b.HomeScreen),
		builder.buildExtendedInput(name, "watch", "appLauncher", "40mm", 2, 44, b.HomeScreen),
		builder.buildExtendedInput(name, "watch", "appLauncher", "44mm", 2, 50, b.HomeScreen),

		// Quick Look
		builder.buildExtendedInput(name, "watch", "quickLook", "38mm", 2, 86, b.ShortLook),
		builder.buildExtendedInput(name, "watch", "quickLook", "42mm", 2, 98, b.ShortLook),
		builder.buildExtendedInput(name, "watch", "quickLook", "44mm", 2, 108, b.ShortLook),
	}

	return images, nil
}
