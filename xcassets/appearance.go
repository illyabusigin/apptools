package xcassets

// Appearance allows you to configure your color for different appearance types.
// See https://developer.apple.com/documentation/uikit/uiimage/providing_images_for_different_appearances
type Appearance struct {
	any          bool
	light        bool
	dark         bool
	highContrast bool
}

func (a *Appearance) build() appearances {
	app := appearances{
		[]appearance{},
	}

	if a.dark {
		app = append(app, []appearance{
			{
				Appearance: "luminosity",
				Value:      "dark",
			},
		})
	}

	if a.light {
		app = append(app, []appearance{
			{
				Appearance: "luminosity",
				Value:      "light",
			},
		})
	}

	if a.highContrast {
		app = append(app, []appearance{
			{
				Appearance: "contrast",
				Value:      "high",
			},
		})

		if a.dark {
			app = append(app, []appearance{
				{
					Appearance: "luminosity",
					Value:      "dark",
				},
				{
					Appearance: "contrast",
					Value:      "high",
				},
			})
		}

		if a.light {
			app = append(app, []appearance{
				{
					Appearance: "luminosity",
					Value:      "light",
				},
				{
					Appearance: "contrast",
					Value:      "high",
				},
			})
		}
	}

	return app
}

func (a *Appearance) intersects(a2 *Appearance) []string {
	overlapping := []string{}

	if a.any == a2.any && a.any == true {
		overlapping = append(overlapping, "Any")
	}

	if a.light == a2.light && a.light == true {
		overlapping = append(overlapping, "Light")
	}

	if a.dark == a2.dark && a.dark == true {
		overlapping = append(overlapping, "Dark")
	}

	if a.highContrast == a2.highContrast && a.highContrast == true {
		overlapping = append(overlapping, "HighContrast")
	}

	return overlapping
}

// Any specifies that this asset is available for any appearances.
func (a *Appearance) Any() {
	a.any = true
}

// Light specifies that this asset is available for light appearances.
func (a *Appearance) Light() {
	a.light = true
}

// Dark specifies that this asset is available for dark appearances.
func (a *Appearance) Dark() {
	a.any = true
	a.dark = true
}

// HighContrast specifies that this asset is availbale for high contrast
// appearances.
func (a *Appearance) HighContrast() {
	a.highContrast = true
}
