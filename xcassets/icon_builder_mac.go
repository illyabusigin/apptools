package xcassets

type AppIconMac struct {
	Source  AppIconSource
	Size16  AppIconSource
	Size32  AppIconSource
	Size128 AppIconSource
	Size256 AppIconSource
	Size512 AppIconSource
	enabled bool
}

func (b *AppIconMac) Configure(f func(*AppIconMac)) {
	f(b)
}

func (b *AppIconMac) Validate(s AppIconSource) error {
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

func (b *AppIconMac) Build(name string, s AppIconSource) ([]AppIconImageInput, error) {
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
