package entitlements

import (
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestDataProtection_Complete(t *testing.T) {
	dataProtection := DataProtection{}
	dataProtection.Complete()
	assert.Equal(t, "NSFileProtectionComplete", dataProtection.value)
}

func TestDataProtection_CompleteUnlessOpen(t *testing.T) {
	dataProtection := DataProtection{}
	dataProtection.CompleteUnlessOpen()
	assert.Equal(t, "NSFileProtectionCompleteUnlessOpen", dataProtection.value)
}

func TestDataProtection_CompleteUntilFirstUserAuthentication(t *testing.T) {
	dataProtection := DataProtection{}
	dataProtection.CompleteUntilFirstUserAuthentication()
	assert.Equal(t, "NSFileProtectionCompleteUntilFirstUserAuthentication", dataProtection.value)
}

func TestDataProtection_None(t *testing.T) {
	dataProtection := DataProtection{}
	dataProtection.None()
	assert.Equal(t, "NSFileProtectionNone", dataProtection.value)
}

func TestDataProtection_Apply(t *testing.T) {
	dataProtection := DataProtection{}
	dataProtection.None()

	e := Entitlements{
		data: map[string]interface{}{},
	}

	dataProtection.Apply(&e)
	assert.Equal(t, "NSFileProtectionNone", e.data["com.apple.developer.default-data-protection"])
}
