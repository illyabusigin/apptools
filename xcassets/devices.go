package xcassets

import "errors"

// Devices contains functions for specifying the idioms for an xcasset.
// See https://developer.apple.com/library/archive/documentation/Xcode/Reference/xcode_ref-Asset_Catalog_Format/ImageSetType.html#//apple_ref/doc/uid/TP40015170-CH25-SW2 for more information.
type Devices struct {
	universal    bool
	iPhone, iPad bool
	catalyst     bool
	carPlay      bool
	appleWatch   bool
	appleTV      bool
	mac          bool
}

// Validate will validate the devices configuration. At least one specified
// device is required (i.e. Universal()).
func (d *Devices) Validate() error {
	if len(d.idioms()) == 0 {
		return errors.New("No devices specified")
	}

	return nil
}

func (d *Devices) idioms() []string {
	idioms := []string{}

	if d.universal {
		idioms = append(idioms, "universal")
	}

	if d.iPhone {
		idioms = append(idioms, "iphone")
	}

	if d.iPad {
		idioms = append(idioms, "ipad")
	}

	if d.carPlay {
		idioms = append(idioms, "car")
	}

	if d.appleWatch {
		idioms = append(idioms, "watch")
	}

	if d.appleTV {
		idioms = append(idioms, "tv")
	}

	if d.mac {
		idioms = append(idioms, "mac")
	}

	return idioms
}

func (d *Devices) subtypes() []string {
	subtypes := []string{}

	if d.catalyst {
		subtypes = append(subtypes, "mac-catalyst")
	}

	return subtypes
}

// Universal specifies that this asset works on any device and platform.
func (d *Devices) Universal() *Devices {
	d.universal = true
	return d
}

// IPhone specifies that this asset is for iPhone devices.
func (d *Devices) IPhone() *Devices {
	d.iPhone = true
	return d
}

// IPad specifies that this asset is for iPhone devices.
func (d *Devices) IPad() *Devices {
	d.iPhone = true
	return d
}

// Catalyst specifies that this asset is for iPad and and Mac devices using
// Catalyst.
func (d *Devices) Catalyst() *Devices {
	d.iPad = true
	d.catalyst = true
	return d
}

// CarPlay specifies that this asset is for CarPlay devices.
func (d *Devices) CarPlay() *Devices {
	d.carPlay = true
	return d
}

// AppleWatch specifies that this asset is for Apple Watch devices.
func (d *Devices) AppleWatch() *Devices {
	d.appleWatch = true
	return d
}

// AppleTV specifies that this asset is for Apple TV.
func (d *Devices) AppleTV() *Devices {
	d.appleTV = true
	return d
}

// Mac specifies that this asset is for Mac computers.
func (d *Devices) Mac() *Devices {
	d.mac = true
	return d
}
