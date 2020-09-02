package xcassets

import (
	"encoding/json"
	"io"
)

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

// ColorBuilder contains methods and properties for manipulating color properties.
type ColorBuilder struct {
	d    *ColorDefinition
	name string
}

// Color specifies the color definition.
func (b *ColorBuilder) Color(f func(d *ColorDefinition)) *ColorBuilder {
	b.d = &ColorDefinition{}
	f(b.d)

	return b
}

// Validate the color set configuration.
func (b *ColorBuilder) Validate() error {
	return nil
}

// Build will construct the Contents.json of the color.
func (b *ColorBuilder) Build() (string, error) {
	if err := b.Validate(); err != nil {
		return "", err
	}

	colorSet := colorSet{
		Info: info{
			Author:  "xcode",
			Version: 1,
		},
		Properties: properties{
			Localizable: true,
		},
		Colors: []colorContainer{},
	}

	data, err := json.Marshal(&colorSet)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

// Write will write the Contents.json to the specified `io.Writer`.
func (b *ColorBuilder) Write(w io.Writer) error {
	data, err := b.Build()
	if err != nil {
		return err
	}

	_, err = w.Write([]byte(data))

	return err
}

type info struct {
	Author  string `json:"xcode"`
	Version int    `json:"version"`
}

type properties struct {
	Localizable bool `json:"localizable"`
}

type appearance struct {
	Appearance string `json:"appearance"`
	Value      string `json:"value"`
}

type appearances [][]appearance

type colorContainer struct {
	Appearances appearances `json:"appearances"`
	Color       color       `json:"color"`
	Idiom       string      `json:"idiom"`
}

type color struct {
	ColorSpace string          `json:"color-space"`
	Components colorComponents `json:"components"`
}

type colorComponents struct {
	Alpha float64 `json:"alpha"`
	Red   float64 `json:"red,omitempty"`
	Green float64 `json:"green,omitempty"`
	Blue  float64 `json:"blue,omitempty"`
	White float64 `json:"white,omitempty"`
}

type colorSet struct {
	Colors     []colorContainer `json:"colors"`
	Info       info             `json:"info"`
	Properties properties       `json:"properties"`
}
