package xcassets

// ColorSpace allows you to specify the color space for the xcasset. If no
// colorspace is specifies the sRGB color space is used.
type ColorSpace struct {
	grayscale bool
	space     string
}

func (c *ColorSpace) build() string {
	if c.space == "" {
		c.SRGB()
	}

	return c.space
}

// SRGB specifies the asset uses the standard sRGB color space.
func (c *ColorSpace) SRGB() {
	c.grayscale = false
	c.space = "srgb"
}

// DisplayP3 specifies the asset uses a wide gamut color space.
func (c *ColorSpace) DisplayP3() {
	c.grayscale = false
	c.space = "display-p3"
}

// ExtendedRangeSRGB specifies the asset uses the extended range sRGB color space.
func (c *ColorSpace) ExtendedRangeSRGB() {
	c.grayscale = false
	c.space = "extended-srgb"
}

// ExtendedRangeLinearSRGB specifies the asset uses the extended range linear
// sRGB color space.
func (c *ColorSpace) ExtendedRangeLinearSRGB() {
	c.grayscale = false
	c.space = "extended-linear-srgb"
}

// GrayGamma22 specifies the asset uses the gray gamma 2.2 color space.
func (c *ColorSpace) GrayGamma22() {
	c.grayscale = true
	c.space = "gray-gamma-22"
}

// ExtendedRangeGray specifies the asset uses the extended range gray color space.
func (c *ColorSpace) ExtendedRangeGray() {
	c.grayscale = true
	c.space = "extended-gray"
}
