package xcassets

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"path/filepath"

	"github.com/nfnt/resize"
)

type AppIconOuput struct {
	Inputs []AppIconImageInput `json:"-"`

	Images []AppIconImage `json:"images"`
	Info   AppIconVersion `json:"info"`
}

func (o *AppIconOuput) WriteImages(path string) error {

	loader := iconImageLoader{}
	cache := map[string]image.Image{}

	for _, input := range o.Inputs {
		img, ok := cache[loader.Key(input.Source)]
		var err error = nil
		if !ok {
			img, err = loader.Load(input.Source)
			if err != nil {
				return fmt.Errorf("Failed to load image from source %v, error: %w", input.Source, err)
			}

			cache[loader.Key(input.Source)] = img
		}

		m := resize.Resize(uint(input.Size*input.Scale), uint(input.Size*input.Scale), img, resize.Lanczos3)
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

type AppIconImageInput struct {
	Size     int
	Idiom    string
	Filename string
	Scale    int
	Source   AppIconSource
}

func (i *AppIconImageInput) Image() AppIconImage {
	return AppIconImage{
		Size:     fmt.Sprintf("%dx%d", i.Size, i.Size),
		Idiom:    i.Idiom,
		Filename: i.Filename,
		Scale:    fmt.Sprintf("%dx", i.Scale),
	}
}

type AppIconImage struct {
	Size     string `json:"size"`
	Idiom    string `json:"idiom"`
	Filename string `json:"filename"`
	Scale    string `json:"scale"`
}

type AppIconVersion struct {
	Version int    `json:"version"`
	Author  string `json:"author"`
}

type IconImageBuilder struct {
}

func (b *IconImageBuilder) buildInput(name, idiom string, scale int, size int, source AppIconSource) AppIconImageInput {
	image := AppIconImageInput{
		Size:     size,
		Idiom:    idiom,
		Filename: fmt.Sprintf("%v-%dx%d@%dx", name, size, size, scale),
		Scale:    scale,
		Source:   source,
	}

	return image
}
