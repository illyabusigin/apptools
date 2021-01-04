package xcassets

import (
	"os"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestAsset(t *testing.T) {
	builder := Asset("Logo", func(b *AssetBuilder) {
		b.Gamut.Any()
		b.Asset(func(d *AssetDefinition) {
			d.Devices.Universal()
			d.Appearance.Dark()

			d.Source.File("./testdata/Icon.png")
			d.Source.Size(256, 256)
		})
	})

	// How do we know the asset size? Assume the maximum size is passed

	assert.NotNil(t, builder)
	assert.Nil(t, builder.Validate())

	err := builder.SaveTo("./_test/", true)
	assert.Nil(t, err)
	assert.Nil(t, os.RemoveAll("./_test/Logo.imageset"))
}
