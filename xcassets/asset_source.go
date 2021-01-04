package xcassets

import (
	"fmt"
	"image"
	"math"
)

type AssetSource struct {
	url          string
	file         string
	validated    bool
	minDimension int

	desiredWidth, desiredHeight int
}

func (s *AssetSource) hasDimensions() bool {
	if s.minDimension != 0 {
		return false
	}

	if s.desiredWidth != 0 && s.desiredHeight != 0 {
		return true
	}

	return false
}

func (s *AssetSource) Validate() error {
	if s.validated {
		return nil
	}

	if s.Empty() {
		return fmt.Errorf("No URL or file location specified for asset source")
	}

	if s.minDimension <= 0 && !s.hasDimensions() {
		return fmt.Errorf("Minimum asset dimension invalid or not specified (%v)", s.minDimension)
	}

	loader := assetLoader{
		source: *s,
	}

	if err := loader.Validate(); err != nil {
		return err
	}

	s.validated = true

	return nil
}

func (s *AssetSource) validateImage(image image.Config, fileName string) error {
	if s.hasDimensions() {
		desiredAspectRatio := float64(s.desiredWidth) / float64(s.desiredHeight)
		aspectRatio := float64(image.Width) / float64(image.Height)
		delta := aspectRatio - desiredAspectRatio

		if math.Abs(delta) > 0.05 {
			return fmt.Errorf("%v aspect ratio (%vx%v) does not match the provided image aspect ratio (%vx%v)",
				fileName, s.desiredWidth, s.desiredHeight,
				image.Width, image.Height)
		}

		scaleFactor := 3 // TODO: Expose this at some point to make it configurable
		if image.Width < s.desiredWidth*scaleFactor || image.Height < s.desiredHeight*scaleFactor {
			return fmt.Errorf("%v dimensions (%vx%v) have dimensions less than the minimum  required for scaling (%vx%v)",
				fileName, image.Width, image.Height, s.desiredWidth, s.desiredHeight)
		}
	} else {
		if image.Width < s.minDimension || image.Height < s.minDimension {
			return fmt.Errorf("%v dimensions (%vx%v) have dimensions less than the minimum (%v)",
				fileName, image.Width, image.Height, s.minDimension)
		}
	}

	return nil
}

func (s *AssetSource) URL(url string) {
	s.url = url
	s.validated = false
}

func (s *AssetSource) File(path string) {
	s.file = path
	s.validated = false
}

func (s AssetSource) Empty() bool {
	return s.url == "" && s.file == ""
}

func (s *AssetSource) Apply(from AssetSource) {
	s.file = from.file
	s.url = from.url
	s.validated = from.validated
}

// MinDimension specifies the minimum dimension for the icon image.
func (s *AssetSource) MinDimension(d int) {
	s.minDimension = d
	s.validated = false
}

// Size of the asset at 1x scale factor. These dimensions are used to
// ensure that the source image is the correct size and aspect ratio.
func (s *AssetSource) Size(width, height uint) {
	s.desiredWidth = int(width)
	s.desiredHeight = int(height)
	s.validated = false
}
