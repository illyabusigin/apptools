package plist

// Orientations contains device orientation related data
type Orientations struct {
	orientations []string
}

// Validate will verify that you have specified valid device orientations
func (d *Orientations) Validate() error {
	if len(d.orientations) == 0 {
		return errMissingProperty("Device orientations are not defined")
	}

	return nil
}

// Portrait specifies the portrait device orientation.
func (d *Orientations) Portrait() {
	d.orientations = append(d.orientations, "UIInterfaceOrientationPortrait")
}

// LandscapeLeft specifies the landscape left device orientation.
func (d *Orientations) LandscapeLeft() {
	d.orientations = append(d.orientations, "UIInterfaceOrientationLandscapeLeft")
}

// LandscapeRight specifies the landscape right device orientation.
func (d *Orientations) LandscapeRight() {
	d.orientations = append(d.orientations, "UIInterfaceOrientationLandscapeRight")
}

// UpsideDown specifies the upside down device orientation.
func (d *Orientations) UpsideDown() {
	d.orientations = append(d.orientations, "UIInterfaceOrientationUpsideDown")
}
