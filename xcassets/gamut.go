package xcassets

// Gamut specifes the color gamut of the device display.
// See https://developer.apple.com/library/archive/documentation/Xcode/Reference/xcode_ref-Asset_Catalog_Format/ImageSetType.html#//apple_ref/doc/uid/TP40015170-CH25-SW35 for more information.
type Gamut struct {
	values []string
}

func (g *Gamut) build() []string {
	if len(g.values) == 0 {
		g.Any()
	}

	return g.values
}

// Any specifes the xcasset uses the standard RGB gamut color space.
func (g *Gamut) Any() {
	g.values = []string{}
}

// SRGBAndDisplayP3 specifes the xcasses uses the standard RGB and wide gamut
// color space.
func (g *Gamut) SRGBAndDisplayP3() {
	g.values = []string{"display-P3", "sRGB"}
}
