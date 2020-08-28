package plist

import (
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestATSExceptionDomain_init(t *testing.T) {
	e := ATSExceptionDomain{}
	e.init()
	assert.True(t, e.requiresForwardSecrecy)
}

func TestATSExceptionDomain_IncludesSubdomains(t *testing.T) {
	e := ATSExceptionDomain{}
	e.IncludesSubdomains(true)
	assert.True(t, e.includeSubdomains)
}

func TestATSExceptionDomain_AllowsInsecureHTTPLoads(t *testing.T) {
	e := ATSExceptionDomain{}
	e.AllowsInsecureHTTPLoads(true)
	assert.True(t, e.allowsInsecureHTTPLoads)
}

func TestATSExceptionDomain_MinimumTLSVersion(t *testing.T) {
	e := ATSExceptionDomain{}
	e.MinimumTLSVersion("TLSv1.3")
	assert.Equal(t, "TLSv1.3", e.minimumTLSVersion)
}

func TestATSExceptionDomain_RequiresForwardSecrecy(t *testing.T) {
	e := ATSExceptionDomain{}
	e.RequiresForwardSecrecy(true)
	assert.True(t, e.requiresForwardSecrecy)
}

func TestATSExceptionDomain_RequiresCertificateTransparency(t *testing.T) {
	e := ATSExceptionDomain{}
	e.RequiresCertificateTransparency(true)
	assert.True(t, e.requiresCertificateTransparency)
}

func TestATSExceptionDomain_build(t *testing.T) {
	type fields struct {
		builder func() *ATSExceptionDomain
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]interface{}
	}{
		{
			name: "Building with all configuration values should succeed",
			fields: fields{
				builder: func() *ATSExceptionDomain {
					e := ATSExceptionDomain{}
					e.RequiresForwardSecrecy(true)
					e.MinimumTLSVersion("TLSv1.3")
					return &e
				},
			},
			want: map[string]interface{}{
				"NSExceptionAllowsInsecureHTTPLoads": false,
				"NSExceptionRequiresForwardSecrecy":  true,
				"NSIncludesSubdomains":               false,
				"NSRequiresCertificateTransparency":  false,
				"NSExceptionMinimumTLSVersion":       "TLSv1.3",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := tt.fields.builder()
			assert.Equal(t, tt.want, e.build())
		})
	}
}
