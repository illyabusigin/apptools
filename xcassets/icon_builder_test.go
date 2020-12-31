package xcassets

import (
	"os"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestAppIcon(t *testing.T) {
	builder := AppIcon("AppIcon", func(b *AppIconBuilder) {
		b.File("./testdata/Icon.png")
		b.Phone().Configure(func(b *AppIconPhone) {
			b.Settings.File("./testdata/Icon.png") // override individual icon configurations
		})
		b.AppStore()

	})

	assert.NotNil(t, builder)
	assert.Nil(t, builder.Validate())

	err := builder.SaveTo("./_test/", true)
	assert.Nil(t, err)

	assert.Nil(t, os.RemoveAll("./_test/AppIcon.appiconset"))
}
