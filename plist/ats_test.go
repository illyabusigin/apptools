package plist

import (
	"testing"

	assert "github.com/stretchr/testify/require"
)

func _TestATSCanary(t *testing.T) {
	ats := &AppTransportSecurity{}
	ats.AllowArbitraryLoads(true)
	ats.AllowArbitraryLoadForMedia(true)
	ats.AllowArbitraryLoadForWebContent(true)
	ats.AllowArbitraryLoadForLocalNetworking(true)
	ats.ExceptionDomain("google.com", func(d *ATSExceptionDomain) {
		d.IncludesSubdomains(true)
		d.AllowsInsecureHTTPLoads(true)
		d.MinimumTLSVersion("1.2")
		d.RequiresForwardSecrecy(true)
		d.RequiresCertificateTransparency(true)
	})

	data :=	ats.build()
	assert.NotNil(t, data)
}