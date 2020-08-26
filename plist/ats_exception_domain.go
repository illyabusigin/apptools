package plist

// ATSExceptionDomain specifies the custom configurations for App Transport
// Security named domains.
type ATSExceptionDomain struct {
	domain                          string
	includeSubdomains               bool
	allowsInsecureHTTPLoads         bool
	minimumTLSVersion               string
	requiresForwardSecrecy          bool
	requiresCertificateTransparency bool
}

func (e *ATSExceptionDomain) init() {
	e.requiresForwardSecrecy = true
}

func (e *ATSExceptionDomain) build() map[string]interface{} {
	data := map[string]interface{}{}

	data["NSIncludesSubdomains"] = e.includeSubdomains
	data["NSExceptionAllowsInsecureHTTPLoads"] = e.allowsInsecureHTTPLoads

	if (e.minimumTLSVersion) != "" {
		data["NSExceptionMinimumTLSVersion"] = e.minimumTLSVersion
	}

	data["NSExceptionRequiresForwardSecrecy"] = e.requiresForwardSecrecy
	data["NSRequiresCertificateTransparency"] = e.requiresCertificateTransparency

	return data
}

// IncludesSubdomains allows you to apply the ATS exceptions for the given
// domain to all subdomains of the domain.
// This key is optional. The default value is NO.
func (e *ATSExceptionDomain) IncludesSubdomains(v bool) {
	e.includeSubdomains = v
}

// AllowsInsecureHTTPLoads specified as `true` allows insecure HTTP loads for
// the given domain, or to be able to loosen the server trust evaluation
// requirements for HTTPS connections to the domain, as described in Performing
//  Manual Server Trust Authentication.
// This key is optional. The default value is NO.
// **NOTE**: You must supply a justification during App Store review if you set
// the key’s value to `true`, as described in https://developer.apple.com/documentation/security/preventing_insecure_network_connections#3138036.
func (e *ATSExceptionDomain) AllowsInsecureHTTPLoads(v bool) {
	e.allowsInsecureHTTPLoads = v
}

// MinimumTLSVersion specifies the minimum Transport Layer Security (TLS)
// version for network connections. This key is optional. The value is a
// string, with a default value of TLSv1.2. Possible values are:
//  TLSv1.0
//  TLSv1.1
//  TLSv1.2
//  TLSv1.3
// **NOTE**: You must supply a justification during App Store review if you use
// this key to set a protocol version lower than 1.2, as described in https://developer.apple.com/documentation/security/preventing_insecure_network_connections#3138036.
func (e *ATSExceptionDomain) MinimumTLSVersion(v string) {
	e.minimumTLSVersion = v
}

// RequiresForwardSecrecy allows you set the value for this key to `false` to
// override the requirement that a server support perfect forward secrecy
// (PFS) for the given domain.
// This key is optional. The default value is `true`, which limits the accepted
// ciphers to those that support PFS through Elliptic Curve Diffie-Hellman
// Ephemeral (ECDHE) key exchange.
func (e *ATSExceptionDomain) RequiresForwardSecrecy(v bool) {
	e.requiresForwardSecrecy = v
}

// RequiresCertificateTransparency allows you to specify `true` so that ATS can
// use the  Certificate Transparency (CT)  protocol to identify mistakenly or
// maliciously issued X.509 certificates.
// This key is optional. The default value is NO.
func (e *ATSExceptionDomain) RequiresCertificateTransparency(v bool) {
	e.requiresCertificateTransparency = v
}
