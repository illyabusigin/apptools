package entitlements

// DataProtection allows you to specify the level of data protection that
// encrypts sensitive user data when accessed on some devices.
// See https://developer.apple.com/documentation/bundleresources/entitlements/com_apple_developer_default-data-protection for more information.
type DataProtection struct {
	value string
}

// Apply will apply the data protection entitlements
func (p *DataProtection) Apply(e *Entitlements) {
	if p.value != "" {
		e.data["com.apple.developer.default-data-protection"] = p.value
	}
}

// Complete specifies the file is stored in an encrypted format on disk and
// cannot be read from or written to while the device is locked or booting.
// See https://developer.apple.com/documentation/foundation/nsfileprotectioncomplete for more information.
func (p *DataProtection) Complete() {
	p.value = "NSFileProtectionComplete"
}

// CompleteUnlessOpen specifies the file is stored in an encrypted format on
// disk after it is closed.
// See https://developer.apple.com/documentation/foundation/nsfileprotectioncompleteunlessopen for more information.
func (p *DataProtection) CompleteUnlessOpen() {
	p.value = "NSFileProtectionCompleteUnlessOpen"
}

// CompleteUntilFirstUserAuthentication specifies the file is stored in an
// encrypted format on disk and cannot be accessed until after the device has
// booted.
// See https://developer.apple.com/documentation/foundation/nsfileprotectioncompleteuntilfirstuserauthentication for more information.
func (p *DataProtection) CompleteUntilFirstUserAuthentication() {
	p.value = "NSFileProtectionCompleteUntilFirstUserAuthentication"
}

// None specifies the file has no special protections associated with it.
// See https://developer.apple.com/documentation/foundation/nsfileprotectionnone for more information.
func (p *DataProtection) None() {
	p.value = "NSFileProtectionNone"
}
