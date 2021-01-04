package xcassets

type AppIconCarPlay struct {
	Source AssetSource
	Icon   AssetSource

	enabled bool
}

func (b *AppIconCarPlay) Configure(f func(*AppIconCarPlay)) {
	f(b)
}

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
