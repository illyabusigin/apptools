package entitlements

import (
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestAPS_Development(t *testing.T) {
	aps := APS{}
	aps.Development()
	assert.Equal(t, "development", aps.environment)
}

func TestAPS_Production(t *testing.T) {
	aps := APS{}
	aps.Production()
	assert.Equal(t, "production", aps.environment)
}

func TestAPS_Apply(t *testing.T) {
	aps := APS{}
	aps.Development()

	e := Entitlements{
		data: map[string]interface{}{},
	}

	aps.Apply(&e)
	assert.Equal(t, "development", e.data["aps-environment"])
}
