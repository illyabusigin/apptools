package xcassets

import (
	"fmt"
	"image"
	"image/png"
	"math"
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

	loader := assetLoader{}
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

		m := resize.Resize(uint(input.Size*float64(input.Scale)), uint(input.Size*float64(input.Scale)), img, resize.Lanczos3)
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
	Size     float64
	Idiom    string
	Filename string
	Scale    int
	Role     string
	Subtype  string
	Source   AssetSource
}

func (i *AppIconImageInput) Image() AppIconImage {
	if delta := math.Floor(i.Size) - i.Size; delta != 0 {
		return AppIconImage{
			Size:     fmt.Sprintf("%.1fx%.1f", i.Size, i.Size),
			Idiom:    i.Idiom,
			Filename: i.Filename,
			Scale:    fmt.Sprintf("%dx", i.Scale),
			Subtype:  i.Subtype,
		}
	} else {
		return AppIconImage{
			Size:     fmt.Sprintf("%.0fx%.0f", i.Size, i.Size),
			Idiom:    i.Idiom,
			Filename: i.Filename,
			Scale:    fmt.Sprintf("%dx", i.Scale),
			Subtype:  i.Subtype,
		}
	}
}

type AppIconImage struct {
	Size     string `json:"size"`
	Idiom    string `json:"idiom"`
	Filename string `json:"filename"`
	Scale    string `json:"scale"`
	Role     string `json:"role,omitempty"`
	Subtype  string `json:"subtype,omitempty"`
}

type AppIconVersion struct {
	Version int    `json:"version"`
	Author  string `json:"author"`
}

type IconImageBuilder struct {
}

func (b *IconImageBuilder) buildInput(name, idiom string, scale int, size float64, source AssetSource) AppIconImageInput {

	if delta := math.Floor(size) - size; delta != 0 {
		return AppIconImageInput{
			Size:     size,
			Idiom:    idiom,
			Filename: fmt.Sprintf("%v-%.1fx%.1f@%dx", name, size, size, scale),
			Scale:    scale,
			Source:   source,
		}
	} else {
		return AppIconImageInput{
			Size:     size,
			Idiom:    idiom,
			Filename: fmt.Sprintf("%v-%.0fx%.0f@%dx", name, size, size, scale),
			Scale:    scale,
			Source:   source,
		}
	}
}

func (b *IconImageBuilder) buildExtendedInput(name, idiom, role, subtype string, scale int, size float64, source AssetSource) AppIconImageInput {

	if delta := math.Floor(size) - size; delta != 0 {
		return AppIconImageInput{
			Size:     size,
			Idiom:    idiom,
			Filename: fmt.Sprintf("%v-%.1fx%.1f@%dx", name, size, size, scale),
			Scale:    scale,
			Source:   source,
			Role:     role,
			Subtype:  subtype,
		}
	} else {
		return AppIconImageInput{
			Size:     size,
			Idiom:    idiom,
			Filename: fmt.Sprintf("%v-%.0fx%.0f@%dx", name, size, size, scale),
			Scale:    scale,
			Source:   source,
			Role:     role,
			Subtype:  subtype,
		}
	}
}
