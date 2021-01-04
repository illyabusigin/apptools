package xcassets

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Asset creates a named image set with the specified name, returning an
// `AssetBuilder` that you can use to customize your asset.
// See https://developer.apple.com/library/archive/documentation/Xcode/Reference/xcode_ref-Asset_Catalog_Format/ImageSetType.html#//apple_ref/doc/uid/TP40015170-CH25-SW1 for more information.
func Asset(name string, f func(b *AssetBuilder)) *AssetBuilder {
	b := AssetBuilder{
		name: name,
		defs: []*AssetDefinition{},
	}

	b.Gamut.Any()

	f(&b)

	return &b
}

// AssetBuilder contains methods and properties for manipulating asset properties.
type AssetBuilder struct {
	defs []*AssetDefinition
	name string

	Gamut      Gamut
	Properties AssetProperties
}

// Asset specifies the asset definition. Certain properties are set by default
// and can be overridden, specifically:
//  d.Appearance.Any()
func (b *AssetBuilder) Asset(f func(d *AssetDefinition)) *AssetBuilder {
	d := &AssetDefinition{}
	d.Appearance.Any()

	b.defs = append(b.defs, d)
	f(d)

	return b
}

// Validate the asset set configuration.
func (b *AssetBuilder) Validate() error {
	if len(b.defs) == 0 {
		return fmt.Errorf("No assets defined for %v", b.name)
	}

	for _, d := range b.defs {
		if err := d.Validate(); err != nil {
			return fmt.Errorf("Invalid asset definition: %v", err)
		}
	}

	// Validate against each other
	for _, d1 := range b.defs {
		for _, d2 := range b.defs {
			if err := d1.detectOverlap(d2); err != nil {
				// TODO: Need a way to identify the invalid def, and bubble up
				return fmt.Errorf("Overlapping asset definitions: %v", err)
			}
		}
	}

	return nil
}

// Build will construct the Contents.json of the asset and validate the
// configuration.
func (b *AssetBuilder) Build() (*AssetOutput, error) {
	if err := b.Validate(); err != nil {
		return nil, err
	}

	output := AssetOutput{
		Info: info{
			Author:  "xcode",
			Version: 1,
		},
		Properties: func() *AssetProperties {
			if b.Properties.Empty() {
				return nil
			}

			b.Properties.build()
			return &b.Properties
		}(),
	}

	inputs, err := b.buildInputs()
	if err != nil {
		return nil, err
	}

	output.inputs = inputs
	images := make([]AssetImage, len(inputs))
	for idx, input := range inputs {
		images[idx] = input.image()
	}
	output.Images = images

	return &output, nil
}

func (b *AssetBuilder) buildInputs() ([]assetInput, error) {
	gamuts := b.Gamut.build()
	inputs := []assetInput{}

	for _, d := range b.defs {
		inputs = append(inputs, d.build(b.name, gamuts)...)
	}

	return inputs, nil
}

func (b *AssetBuilder) exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// SaveTo will save the asset to the specified path.
func (b *AssetBuilder) SaveTo(path string, overwrite bool) error {
	stat, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("Path does not exist: %w", err)
		}

		return fmt.Errorf("Failed to validate path: %w", err)
	}

	if !stat.IsDir() {
		return fmt.Errorf("SaveTo path must be a directory")
	}

	// Create .appiconset folder
	folder := filepath.Join(path, fmt.Sprintf("%v.imageset", b.name))
	if exists, err := b.exists(folder); exists && err == nil && overwrite {
		if err := os.RemoveAll(folder); err != nil {
			return fmt.Errorf("Failed to remove .imageset folder: %w", err)
		}
	}

	if err = os.Mkdir(folder, os.ModePerm); err != nil {
		return fmt.Errorf("Failed to create imageset folder: %w", err)
	}

	output, err := b.Build()
	if err != nil {
		return err
	}

	data, err := json.Marshal(output)
	if err != nil {
		return fmt.Errorf("Failed to marshal Contents.json: %w", err)
	}

	err = ioutil.WriteFile(filepath.Join(folder, "Contents.json"), data, os.ModePerm)
	if err != nil {
		return fmt.Errorf("Failed to write Contents.json to file: %w", err)
	}

	err = output.WriteImages(folder)
	if err != nil {
		return fmt.Errorf("Failed to write images to file: %w", err)
	}

	return nil
}
