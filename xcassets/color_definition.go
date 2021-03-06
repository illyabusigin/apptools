package xcassets

import (
	"fmt"
	"strings"

	"gopkg.in/go-playground/colors.v1"
)

// ColorDefinition defines the a specific color and its various properties such
// as colorspace, devices, color gamut.
type ColorDefinition struct {
	// Use the functions on ColorSpace to define the colorspace.
	ColorSpace ColorSpace
	Appearance Appearance
	Devices    Devices

	hex string

	r, g, b       float64
	eighBit       bool
	floatingPoint bool

	alpha float64
	white float64
}

// Validate will ensure that you have a valid `ColorDefinition`, returning
// any errors.
func (d *ColorDefinition) Validate() error {
	if !d.colorPresent() {
		return fmt.Errorf("No color present - please specify a color")
	}

	if err := d.validateColor(); err != nil {
		return err
	}

	if err := d.Devices.Validate(); err != nil {
		return err
	}

	return nil
}

func (d *ColorDefinition) color(gamuts []string) color {
	c := color{
		ColorSpace: d.ColorSpace.build(),
		Components: colorComponents{
			Alpha: fmt.Sprintf("%.3f", d.alpha),
		},
	}

	if d.hex != "" {
		c.Components.Red = fmt.Sprintf("0x%v", d.hex[1:3])
		c.Components.Green = fmt.Sprintf("0x%v", d.hex[3:5])
		c.Components.Blue = fmt.Sprintf("0x%v", d.hex[5:])
	}
	if d.eighBit {
		c.Components.Red = fmt.Sprintf("%d", int(d.r))
		c.Components.Green = fmt.Sprintf("%d", int(d.g))
		c.Components.Blue = fmt.Sprintf("%d", int(d.b))
	}

	if d.floatingPoint {
		c.Components.Red = fmt.Sprintf("%.3f", d.r)
		c.Components.Green = fmt.Sprintf("%.3f", d.g)
		c.Components.Blue = fmt.Sprintf("%.3f", d.b)
	}

	if d.ColorSpace.grayscale {
		c.Components.White = fmt.Sprintf("%.3f", d.white)
	}

	return c
}

func (d *ColorDefinition) detectOverlap(d2 *ColorDefinition) error {
	if d == d2 {
		return nil
	}

	if intersection := d.Devices.intersects(&d2.Devices); len(intersection) > 0 {
		return fmt.Errorf("Devices between color definitions overlap (%v) - they must be unique", strings.Join(intersection, ","))
	}

	if intersection := d.Appearance.intersects(&d2.Appearance); len(intersection) > 0 {
		return fmt.Errorf("Appearances overlap (%v) - they must be unique", strings.Join(intersection, ","))
	}

	return nil
}

func (d *ColorDefinition) validateColor() error {
	if d.hex != "" {
		if err := d.Hex(d.hex); err != nil {
			return err
		}
	}

	if d.eighBit {
		if err := d.RGB(int(d.r), int(d.g), int(d.b)); err != nil {
			return err
		}
	}

	if d.floatingPoint {
		if err := d.RGBFloat(d.r, d.g, d.b); err != nil {
			return err
		}
	}

	if err := d.Alpha(d.alpha); err != nil {
		return err
	}

	if d.ColorSpace.grayscale {
		if err := d.White(d.white); err != nil {
			return err
		}
	}

	return nil
}

func (d *ColorDefinition) colorPresent() bool {
	if d.hex != "" {
		return true
	}

	if d.eighBit || d.floatingPoint {
		return true
	}

	if d.ColorSpace.grayscale {
		return true
	}

	return false
}

// Hex specifies a color in hexadecimal format (i.e. #262d44).
func (d *ColorDefinition) Hex(v string) error {
	d.eighBit = false
	d.floatingPoint = false

	d.hex = v
	hex, err := colors.ParseHEX(v)
	if err != nil {
		return err
	}

	d.hex = hex.String()
	return nil
}

// Alpha defines the alpha, or opacity of the color.
func (d *ColorDefinition) Alpha(a float64) error {
	d.alpha = a

	if a < 0 || a > 1 {
		return fmt.Errorf("Alpha is invalid (%v), must be between 0 and 1", a)
	}

	return nil
}

// White defines the white level of your color. Should only be used in
// conjunction with gray-scale color spaces. Value must be between 0 and 1.
func (d *ColorDefinition) White(v float64) error {
	d.hex = ""

	if !d.ColorSpace.grayscale {
		return fmt.Errorf("White is not currently available - Can only be used in conjunction with grayscale color spaces")
	}

	d.white = v

	if v < 0 || v > 1 {
		return fmt.Errorf("White is invalid (%v), must be between 0 and 1", v)
	}

	return nil
}

// RGB allows you to specify an 8-bit RGB color. Each value must be
// between 0 and 255.
func (d *ColorDefinition) RGB(r, g, b int) error {
	d.hex = ""
	d.eighBit = true
	d.floatingPoint = false

	d.r = float64(r)
	d.g = float64(g)
	d.b = float64(b)

	if r < 0 || r > 255 {
		return fmt.Errorf("RGB.red is invalid (%v), must be between 0 and 255", r)
	}

	if g < 0 || g > 255 {
		return fmt.Errorf("RGB.green is invalid (%v), must be between 0 and 255", g)
	}

	if b < 0 || b > 255 {
		return fmt.Errorf("RGB.blue is invalid (%v), must be between 0 and 255", b)
	}

	return nil
}

// RGBFloat specifies a floating poiny RGB color. Each value must be between
// 0 and 1.
func (d *ColorDefinition) RGBFloat(r, g, b float64) error {
	d.hex = ""
	d.floatingPoint = true
	d.eighBit = false

	d.r = r
	d.g = g
	d.b = b

	if r < 0 || r > 1 {
		return fmt.Errorf("RGBFloat.red is invalid (%v), must be between 0 and 1", r)
	}

	if g < 0 || g > 1 {
		return fmt.Errorf("RGBFloat.green is invalid (%v), must be between 0 and 1", g)
	}

	if b < 0 || b > 1 {
		return fmt.Errorf("RGBFloat.blue is invalid (%v), must be between 0 and 1", b)
	}

	return nil
}
