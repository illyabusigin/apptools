package xcassets

import "io"

func _colorBuilder() {
	Color("SplashScreenColor", func(b *ColorBuilder) {
		b.Color(func(d *ColorDefinition) {
			d.Devices.Universal().IPhone()
			d.ColorSpace.SRGB()
			d.Gamut.Any()
			d.Gamut.SRGBAndDisplayP3()
			d.Appearance.Any()
			d.Appearance.Light()
			d.Appearance.Dark()
			d.Appearance.HighContrast()

			d.Hex("#262D44")
			d.White(1) //used for gray colors
			d.RGB(146, 144, 0)
			d.RGBFloat(0.682, 0.682, 0.682)

			d.Alpha(.4)

			// d.Color.System.DarkTextColor()
		})
		// Define colors, then assign them  idioms, gammut, appearance, high contrast
	}) //.To("path/to/folder")
}

// Color creates a named color type with the specified name, returning a
// `ColorBuilder` function  that you can use to customize your color.
// See https://developer.apple.com/library/archive/documentation/Xcode/Reference/xcode_ref-Asset_Catalog_Format/Named_Color.html#//apple_ref/doc/uid/TP40015170-CH59-SW1 for more information.
func Color(name string, f func(b *ColorBuilder)) *ColorBuilder {
	b := ColorBuilder{
		name: name,
	}

	return &b
}

type ColorBuilder struct {
	d    *ColorDefinition
	name string
}

func (b *ColorBuilder) Name(v string) {
	b.name = v
}

func (b *ColorBuilder) Color(f func(d *ColorDefinition)) {
	b.d = &ColorDefinition{}
	f(b.d)
}

func (b *ColorBuilder) Build() (string, error) {
	return "", nil
}

func (b *ColorBuilder) Write(w io.Writer) error {
	data, err := b.Build()
	if err != nil {
		return err
	}

	_, err = w.Write([]byte(data))

	return err
}

// d.Hex("#262D44")
// 			d.White(1) //used for gray colors
// 			d.RGB(146, 144, 0)
// 			d.RGBFloat(0.682, 0.682, 0.682)

// 			d.Alpha(.4)
