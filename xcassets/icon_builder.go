package xcassets

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

const MinDimension = 1024

//
func AppIcon(name string, f func(b *AppIconBuilder)) *AppIconBuilder {
	b := &AppIconBuilder{
		Name: name,
		AssetSource: AssetSource{
			minDimension: MinDimension,
		},
	}

	if f != nil {
		f(b)
	}

	return b
}

type AppIconBuilder struct {
	Name string
	AssetSource
	iPhone   AppIconPhone
	iPad     AppIconTablet
	watch    AppIconWatch
	carPlay  AppIconCarPlay
	mac      AppIconMac
	appStore AppIconAppStore
}

func (b *AppIconBuilder) Validate() error {
	if !b.AssetSource.Empty() {
		if err := b.AssetSource.Validate(); err != nil {
			return fmt.Errorf("Source is invalid: %w", err)
		}
	}

	// Validate each idiom
	if err := b.iPhone.Validate(b.AssetSource); err != nil {
		return fmt.Errorf("Failed to validate iphone idiom: %w", err)
	}

	if err := b.iPad.Validate(b.AssetSource); err != nil {
		return fmt.Errorf("Failed to validate ipad idiom: %w", err)
	}

	if err := b.watch.Validate(b.AssetSource); err != nil {
		return fmt.Errorf("Failed to validate watch idiom: %w", err)
	}

	if err := b.mac.Validate(b.AssetSource); err != nil {
		return fmt.Errorf("Failed to validate mac idiom: %w", err)
	}

	if err := b.carPlay.Validate(b.AssetSource); err != nil {
		return fmt.Errorf("Failed to validate car idiom: %w", err)
	}

	if err := b.appStore.Validate(b.AssetSource); err != nil {
		return fmt.Errorf("Failed to validate ios-marketing idiom: %w", err)
	}

	return nil
}

// Build will validate and build the app icon.
func (b *AppIconBuilder) Build() (*AppIconOuput, error) {
	if err := b.Validate(); err != nil {
		return nil, err
	}
	output := AppIconOuput{
		Images: []AppIconImage{},
		Inputs: []AppIconImageInput{},
		Info: AppIconVersion{
			Version: 1,
			Author:  "xcode",
		},
	}

	// Build structs for json
	if b.iPhone.enabled {
		inputs, err := b.iPhone.Build(b.Name, b.AssetSource)
		if err != nil {
			return nil, fmt.Errorf("Failed to build images for iphone: %w", err)
		}
		output.Inputs = append(output.Inputs, inputs...)
	}

	if b.iPad.enabled {
		inputs, err := b.iPad.Build(b.Name, b.AssetSource)
		if err != nil {
			return nil, fmt.Errorf("Failed to build images for ipad: %w", err)
		}
		output.Inputs = append(output.Inputs, inputs...)
	}

	if b.watch.enabled {
		inputs, err := b.watch.Build(b.Name, b.AssetSource)
		if err != nil {
			return nil, fmt.Errorf("Failed to build images for watch: %w", err)
		}
		output.Inputs = append(output.Inputs, inputs...)
	}

	if b.carPlay.enabled {
		inputs, err := b.carPlay.Build(b.Name, b.AssetSource)
		if err != nil {
			return nil, fmt.Errorf("Failed to build images for car: %w", err)
		}
		output.Inputs = append(output.Inputs, inputs...)
	}

	if b.mac.enabled {
		inputs, err := b.mac.Build(b.Name, b.AssetSource)
		if err != nil {
			return nil, fmt.Errorf("Failed to build images for mac: %w", err)
		}
		output.Inputs = append(output.Inputs, inputs...)
	}

	if b.appStore.enabled {
		inputs, err := b.appStore.Build(b.Name, b.AssetSource)
		if err != nil {
			return nil, fmt.Errorf("Failed to build images for ios-marketing: %w", err)
		}
		output.Inputs = append(output.Inputs, inputs...)
	}

	output.Images = make([]AppIconImage, len(output.Inputs))

	for idx, input := range output.Inputs {
		output.Images[idx] = input.Image()
	}

	return &output, nil
}

func (b *AppIconBuilder) exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// SaveTo will save the application icon to the specified path.
func (b *AppIconBuilder) SaveTo(path string, overwrite bool) error {
	stat, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("Path does not exist: %w", err)
		} else {
			return fmt.Errorf("Failed to validate path: %w", err)
		}
	}

	if !stat.IsDir() {
		return fmt.Errorf("SaveTo path must be a directory")
	}

	// Create .appiconset folder
	folder := filepath.Join(path, fmt.Sprintf("%v.appiconset", b.Name))
	if exists, err := b.exists(folder); exists && err == nil && overwrite {
		if err := os.RemoveAll(folder); err != nil {
			return fmt.Errorf("Failed to remove .appiconset folder: %w", err)
		}
	}

	if err = os.Mkdir(folder, os.ModePerm); err != nil {
		return fmt.Errorf("Failed to create appiconset folder: %w", err)
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

// Phone enables phone icons.
func (b *AppIconBuilder) Phone() *AppIconPhone {
	b.iPhone.enabled = true
	return &b.iPhone
}

// Tablet enables iPad icons.
func (b *AppIconBuilder) Tablet() *AppIconTablet {
	b.iPad.enabled = true
	return &b.iPad
}

// Mac enables Mac icons.
func (b *AppIconBuilder) Mac() *AppIconMac {
	b.mac.enabled = true
	return &b.mac
}

// CarPlay enables CarPlay icons.
func (b *AppIconBuilder) CarPlay() *AppIconCarPlay {
	b.carPlay.enabled = true
	return &b.carPlay
}

// AppStore enables App Store icons.
func (b *AppIconBuilder) AppStore() *AppIconBuilder {
	b.appStore.enabled = true
	return b
}

// Watch enables Apple Watch icons.
func (b *AppIconBuilder) Watch() *AppIconWatch {
	b.watch.enabled = true
	return &b.watch
}
