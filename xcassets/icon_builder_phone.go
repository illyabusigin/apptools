package xcassets

type AppIconPhone struct {
	Source       AppIconSource
	Notification AppIconSource
	Spotlight    AppIconSource
	Settings     AppIconSource
	App          AppIconSource
	enabled      bool
}

func (b *AppIconPhone) Configure(f func(*AppIconPhone)) {
	f(b)
}

func (b *AppIconPhone) Validate(s AppIconSource) error {
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

func (b *AppIconPhone) Build(name string, s AppIconSource) ([]AppIconImageInput, error) {
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
