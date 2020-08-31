package xcassets

type ColorDefinition struct {
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

func (d *ColorDefinition) Hex(v string) {
	d.hex = v
}

func (d *ColorDefinition) Alpha(a float64) {
}

func (d *ColorDefinition) White(v float64) {
	d.white = v
}

func (d *ColorDefinition) RGB(r, g, b int) {
	d.r = float64(r)
	d.g = float64(g)
	d.b = float64(b)
}

func (d *ColorDefinition) RGBFloat(r, g, b float64) {
	d.r = r
	d.g = g
	d.b = b
}
