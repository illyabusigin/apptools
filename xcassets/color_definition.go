package xcassets

import (
	"fmt"

	"gopkg.in/go-playground/colors.v1"
)

// ColorDefinition defines the a specific color and its various properties such
// as colorspace, devices, color gamut.
type ColorDefinition struct {
	// Use the functions on ColorSpace to define the colorspace.
	ColorSpace ColorSpace
	Gamut      Gamut
	Appearance Appearance
	Devices    Devices

	hex string

	r, g, b       float64
	eighBit       bool
	floatingPoint bool

	alpha float64
	white float64
}

// Hex specifies a color in hexadecimal format (i.e. #262d44).
func (d *ColorDefinition) Hex(v string) error {
	hex, err := colors.ParseHEX(v)
	if err != nil {
		return err
	}

	d.hex = hex.String()
	return nil
}

// Alpha defines the alpha, or opacity of the color.
func (d *ColorDefinition) Alpha(a float64) error {
	if a < 0 || a > 1 {
		return fmt.Errorf("Alpha is invalid (%v), must be between 0 and 1", a)
	}

	d.alpha = a
	return nil
}

// White defines the white level of your color. Should only be used in
// conjunction with gray-scale color spaces.
func (d *ColorDefinition) White(v float64) error {
	if !d.ColorSpace.grayscale {
		return fmt.Errorf("White is not currently available - Can only be used in conjunction with grayscale color spaces")
	}

	if v < 0 || v > 1 {
		return fmt.Errorf("White is invalid (%v), must be between 0 and 1", v)
	}

	d.white = v
	return nil
}

// RGB allows you to specify an 8-bit RGB color. Each value must be
// between 0 and 255.
func (d *ColorDefinition) RGB(r, g, b int) error {
	d.eighBit = true
	d.floatingPoint = false

	if r < 0 || r > 255 {
		return fmt.Errorf("RGB.red is invalid (%v), must be between 0 and 255", r)
	}

	if g < 0 || g > 255 {
		return fmt.Errorf("RGB.green is invalid (%v), must be between 0 and 255", g)
	}

	if b < 0 || b > 255 {
		return fmt.Errorf("RGB.blue is invalid (%v), must be between 0 and 255", b)
	}

	d.r = float64(r)
	d.g = float64(g)
	d.b = float64(b)

	return nil
}

// RGBFloat specifies a floating poiny RGB color. Each value must be between
// 0 and 1.
func (d *ColorDefinition) RGBFloat(r, g, b float64) error {
	d.floatingPoint = true
	d.eighBit = false

	if r < 0 || r > 1 {
		return fmt.Errorf("RGBFloat.red is invalid (%v), must be between 0 and 1", r)
	}

	if g < 0 || g > 1 {
		return fmt.Errorf("RGBFloat.green is invalid (%v), must be between 0 and 1", g)
	}

	if b < 0 || b > 1 {
		return fmt.Errorf("RGBFloat.blue is invalid (%v), must be between 0 and 1", b)
	}

	d.r = r
	d.g = g
	d.b = b

	return nil
}
