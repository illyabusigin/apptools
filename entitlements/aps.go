package entitlements

// APS allows you to specify the environment for push notifications.
// See https://developer.apple.com/documentation/bundleresources/entitlements/aps-environment for more information.
type APS struct {
	environment string
}

// Apply will apply the APS entitlements
func (a *APS) Apply(e *Entitlements) {
	if a.environment != "" {
		e.data["aps-environment"] = a.environment
	}
}

// Development specifies the APNs development environment.
func (a *APS) Development() {
	a.environment = "development"
}

// Production specifies the APNs production environment.
func (a *APS) Production() {
	a.environment = "production"
}
