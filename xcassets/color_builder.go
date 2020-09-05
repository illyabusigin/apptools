package xcassets

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

// Color creates a named color type with the specified name, returning a
// `ColorBuilder` function  that you can use to customize your color.
// See https://developer.apple.com/library/archive/documentation/Xcode/Reference/xcode_ref-Asset_Catalog_Format/Named_Color.html#//apple_ref/doc/uid/TP40015170-CH59-SW1 for more information.
func Color(name string, f func(b *ColorBuilder)) *ColorBuilder {
	b := ColorBuilder{
		name: name,
		defs: []*ColorDefinition{},
	}

	b.Gamut.Any()

	f(&b)

	return &b
}

// ColorBuilder contains methods and properties for manipulating color properties.
type ColorBuilder struct {
	defs []*ColorDefinition
	name string

	Gamut Gamut
}

// Color specifies the color definition. Certain properties are set by default and can be overridden, specifically:
//  d.Appearance.Any()
//	d.ColorSpace.SRGB()
//  d.Alpha(1)
func (b *ColorBuilder) Color(f func(d *ColorDefinition)) *ColorBuilder {
	d := &ColorDefinition{}
	d.Appearance.Any()
	d.ColorSpace.SRGB()
	d.Alpha(1)

	b.defs = append(b.defs, d)
	f(d)

	return b
}

// Validate the color set configuration.
func (b *ColorBuilder) Validate() error {
	if len(b.defs) == 0 {
		return fmt.Errorf("No colors defined for %v", b.name)
	}

	for _, d := range b.defs {
		if err := d.Validate(); err != nil {
			return fmt.Errorf("Invalid color definition: %v", err)
		}
	}

	// Validate against each other
	for _, d1 := range b.defs {
		for _, d2 := range b.defs {
			if err := d1.detectOverlap(d2); err != nil {
				// TODO: Need a way to identify the invalid def, and bubble up
				return fmt.Errorf("Overlapping color definitions: %v", err)
			}
		}
	}

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
		Colors: b.buildColors(),
	}

	data, err := json.Marshal(&colorSet)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func (b *ColorBuilder) buildColors() []colorContainer {
	colors := []colorContainer{}
	gamuts := b.Gamut.build()

	// devices(1) * appearance(3) * gamut(2)
	// if high contrast * 2

	for _, d := range b.defs {
		devices := d.Devices.build()
		apps := d.Appearance.build()

		for _, device := range devices {
			for _, appearance := range apps {
				c := colorContainer{
					Appearances: appearance,
					Color:       d.color(gamuts),
					Idiom:       device,
				}

				if len(gamuts) > 0 {
					c.DisplayGamut = gamuts[0]
					c.Color.ColorSpace = strings.ToLower(gamuts[0])
				}

				colors = append(colors, c)

				if len(gamuts) > 1 {
					c2 := c
					c2.DisplayGamut = gamuts[1]
					c2.Color.ColorSpace = strings.ToLower(gamuts[1])

					colors = append(colors, c2)
				}
			}
		}

	}

	return colors
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
