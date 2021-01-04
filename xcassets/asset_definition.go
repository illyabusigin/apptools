package xcassets

import (
	"fmt"
	"math"
	"strings"
)

// AssetImage is used to construct the JSON in Contents.json `images`.
type AssetImage struct {
	Filename string `json:"filename"`
	Idiom    string `json:"idiom"`
	Scale    string `json:"scale"`
	Subtype  string `json:"subtype,omitempty"`

	Appearances  []appearance `json:"appearances,omitempty"`
	DisplayGamut string       `json:"display-gamut,omitempty"`
}

type assetInput struct {
	Width    int
	Height   int
	Idiom    string
	Filename string
	Scale    float64
	Role     string
	Subtype  string

	Appearances  []appearance
	DisplayGamut string

	Source AssetSource
}

func (i *assetInput) image() AssetImage {
	if delta := math.Floor(i.Scale) - i.Scale; delta != 0 {
		return AssetImage{
			Filename:    i.Filename + ".png",
			Idiom:       i.Idiom,
			Scale:       fmt.Sprintf("%.1fx", i.Scale),
			Subtype:     i.Subtype,
			Appearances: i.Appearances,
		}
	}

	return AssetImage{
		Filename:    i.Filename + ".png",
		Idiom:       i.Idiom,
		Scale:       fmt.Sprintf("%0.fx", i.Scale),
		Subtype:     i.Subtype,
		Appearances: i.Appearances,
	}
}

// AssetDefinition defines the a specific color and its various properties such
// as colorspace, devices, color gamut.
type AssetDefinition struct {
	Appearance Appearance
	Devices    Devices

	Source AssetSource
}

// Validate will ensure that you have a valid `ColorDefinition`, returning
// any errors.
func (d *AssetDefinition) Validate() error {
	if err := d.Devices.Validate(); err != nil {
		return err
	}

	if d.Source.Empty() {
		return fmt.Errorf("No asset present - please specify an asset source")
	}

	if err := d.Source.Validate(); err != nil {
		return err
	}

	return nil
}

func (d *AssetDefinition) detectOverlap(d2 *AssetDefinition) error {
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

type assetImageInput struct {
	Size     float64
	Idiom    string
	Filename string
	Scale    int
	Role     string
	Subtype  string
	Source   AssetSource
}

func (d *AssetDefinition) build(name string, gamuts []string) []assetInput {
	// devices(1) * appearance(3) * gamut(2)
	// if high contrast * 2
	// * scale factor
	inputs := []assetInput{}

	devices := d.Devices.build()
	apps := d.Appearance.build()

	for _, device := range devices {
		for _, appearance := range apps {
			inputs = append(inputs, d.buildInputs(name, device, appearance, d.Source, gamuts)...)
		}
	}

	return inputs
}

func (d *AssetDefinition) buildInputs(name, device string, appearance []appearance, source AssetSource, gamuts []string) []assetInput {
	containers := []assetInput{}

	c := assetInput{
		Appearances: appearance,
		Idiom:       device,
		Source:      d.Source,
	}

	if len(gamuts) > 0 {
		c.DisplayGamut = gamuts[0]
	}

	containers = append(containers, c)

	if len(gamuts) > 1 {
		c2 := c
		c2.DisplayGamut = gamuts[1]

		containers = append(containers, c2)
	}

	final := []assetInput{}

	switch device {
	case "universal":
		for _, c := range containers {
			variations := []assetInput{
				d.buildInput(c, name, device, 1),
				d.buildInput(c, name, device, 2),
				d.buildInput(c, name, device, 3),
			}

			final = append(final, variations...)
		}
		// TODO: other device types
	}

	return final
}

func (d *AssetDefinition) buildInput(c assetInput, name, device string, scale float64) assetInput {
	c.Scale = scale
	c.Width = int(float64(d.Source.desiredWidth) * scale)
	c.Height = int(float64(d.Source.desiredHeight) * scale)

	if delta := math.Floor(scale) - scale; delta != 0 {
		c.Filename = fmt.Sprintf("%v-%v-%.1fx%.1f@%1.fx", name, device, float64(d.Source.desiredWidth)*scale, float64(d.Source.desiredHeight)*scale, scale)
	} else {
		c.Filename = fmt.Sprintf("%v-%v-%.0fx%.0f@%.0fx", name, device, float64(d.Source.desiredWidth)*scale, float64(d.Source.desiredHeight)*scale, scale)
	}

	return c
}
