package xcassets

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"path/filepath"

	"github.com/nfnt/resize"
)

// AssetOutput represents the `Contents.json` file used in an image set.
type AssetOutput struct {
	Images     []AssetImage     `json:"images"`
	Info       info             `json:"info"`
	Properties *AssetProperties `json:"properties,omitempty"`

	inputs []assetInput
}

// WriteImages will write the file in `Images` to the specified path.
func (o *AssetOutput) WriteImages(path string) error {

	loader := assetLoader{}
	cache := map[string]image.Image{}

	for _, input := range o.inputs {
		img, ok := cache[loader.Key(input.Source)]
		var err error = nil
		if !ok {
			img, err = loader.Load(input.Source)
			if err != nil {
				return fmt.Errorf("Failed to load image from source %v, error: %w", input.Source, err)
			}

			cache[loader.Key(input.Source)] = img
		}

		m := resize.Resize(uint(input.Width), uint(input.Height), img, resize.Lanczos3)
		fileName := input.Filename + ".png"
		dest := filepath.Join(path, fileName)

		out, err := os.Create(dest)
		if err != nil {
			return fmt.Errorf("Failed to create image file, error: %w", err)
		}
		defer out.Close()

		if err := png.Encode(out, m); err != nil {
			return fmt.Errorf("Failed to write image file, error: %w", err)
		}
	}

	return nil
}
