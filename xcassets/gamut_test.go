package xcassets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGamut_Any(t *testing.T) {
	g := Gamut{}
	g.Any()

	assert.Empty(t, g.values)
}

func TestGamut_SRGBAndDisplayP3(t *testing.T) {
	g := Gamut{}
	g.SRGBAndDisplayP3()

	assert.NotEmpty(t, g.values)
	assert.Equal(t, []string{"display-P3", "sRGB"}, g.values)
}
