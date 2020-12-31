package xcassets

type AppIconTablet struct {
	Source       AppIconSource
	Notification AppIconSource
	Settings     AppIconSource
	Spotlight    AppIconSource
	App          AppIconSource
	Pro          AppIconSource
	enabled      bool
}

func (b *AppIconTablet) Configure(f func(*AppIconTablet)) {
	f(b)
}

func (b *AppIconTablet) Validate(s AppIconSource) error {
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

func (b *AppIconTablet) Build(name string, s AppIconSource) ([]AppIconImageInput, error) {
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
