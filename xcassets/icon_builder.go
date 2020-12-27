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
		AppIconSource: AppIconSource{
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
	AppIconSource
	iPhone   AppIconPhone
	iPad     AppIconTablet
	watch    AppIconWatch
	carPlay  AppIconWatch
	mac      AppIconMac
	appStore AppIconAppStore
}

func (b *AppIconBuilder) Validate() error {
	if !b.AppIconSource.Empty() {
		if err := b.AppIconSource.Validate(); err != nil {
			return fmt.Errorf("Source is invalid: %w", err)
		}
	}

	// Validate each idiom
	if err := b.iPhone.Validate(b.AppIconSource); err != nil {
		return fmt.Errorf("Failed to validate iPhone idiom: %w", err)
	}
	// Validate the source works
	// Validate image dimensions for each idom

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

	// Build struct for json
	if b.iPhone.enabled {
		inputs, err := b.iPhone.Build(b.Name, b.AppIconSource)
		if err != nil {
			return nil, fmt.Errorf("Failed to build images for iPhone: %w", err)
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
func (b *AppIconBuilder) Phone() *AppIconBuilder {
	b.iPhone.enabled = true
	return b
}

// Tablet enables iPad icons.
func (b *AppIconBuilder) Tablet() *AppIconBuilder {
	b.iPad.enabled = true
	return b
}

// Mac enables Mac icons.
func (b *AppIconBuilder) Mac() *AppIconBuilder {
	b.mac.enabled = true
	return b
}

// CarPlay enables CarPlay icons.
func (b *AppIconBuilder) CarPlay() *AppIconBuilder {
	b.carPlay.enabled = true
	return b
}

// AppStore enables App Store icons.
func (b *AppIconBuilder) AppStore() *AppIconBuilder {
	b.appStore.enabled = true
	return b
}

// Watch enables Apple Watch icons.
func (b *AppIconBuilder) Watch() *AppIconBuilder {
	b.watch.enabled = true
	return b
}

type AppIconPhone struct {
	Source       AppIconSource
	Notification AppIconSource
	Spotlight    AppIconSource
	Settings     AppIconSource
	App          AppIconSource
	enabled      bool
}

func (b *AppIconPhone) Validate(s AppIconSource) error {
	if !b.enabled {
		return nil
	}

	b.Source.minDimension = 1024
	if b.Source.Empty() {
		b.Source.Apply(s)
	}

	if err := b.Source.Validate(); err != nil {
		return err
	}

	if b.Notification.Empty() {
		b.Notification.Apply(b.Source)
	}

	if b.Spotlight.Empty() {
		b.Spotlight.Apply(b.Source)
	}

	if b.Settings.Empty() {
		b.Settings.Apply(b.Source)
	}

	if b.Settings.Empty() {
		b.Settings.Apply(b.Source)
	}

	if b.App.Empty() {
		b.App.Apply(b.Source)
	}

	return nil
}

func (b *AppIconPhone) Build(name string, s AppIconSource) ([]AppIconImageInput, error) {
	if err := b.Validate(s); err != nil {
		return nil, err
	}

	builder := IconImageBuilder{}

	images := []AppIconImageInput{
		// iPhone Notifications iOS 7-14
		builder.buildInput(name, "iphone", 2, 20, b.Notification),
		builder.buildInput(name, "iphone", 3, 20, b.Notification),

		// iPhone iOS 7-14
		builder.buildInput(name, "iphone", 2, 29, b.Settings),
		builder.buildInput(name, "iphone", 3, 29, b.Settings),

		// iPhone Spotlight iOS 7-14
		builder.buildInput(name, "iphone", 2, 40, b.Spotlight),
		builder.buildInput(name, "iphone", 3, 40, b.Spotlight),

		// iPhone App Icon iOS 7-14
		builder.buildInput(name, "iphone", 2, 60, b.App),
		builder.buildInput(name, "iphone", 3, 60, b.App),

		// Application Icon
		builder.buildInput(name, "ios-marketing", 1, 1024, b.App),
	}

	return images, nil
}

type AppIconTablet struct {
	Source       AppIconSource
	Notification AppIconSource
	Settings     AppIconSource
	Spotlight    AppIconSource
	App          AppIconSource
	Pro          AppIconSource
	enabled      bool
}

type AppIconMac struct {
	Source  AppIconSource
	Size16  AppIconSource
	Size32  AppIconSource
	Size128 AppIconSource
	Size256 AppIconSource
	Size512 AppIconSource
	enabled bool
}

type AppIconWatch struct {
	Source       AppIconSource
	Notification AppIconSource
	Settings     AppIconSource
	HomeScreen   AppIconSource
	ShortLook    AppIconSource
	enabled      bool
}

type AppIconAppStore struct {
	Source  AppIconSource
	enabled bool
}

type AppIconSource struct {
	url          string
	file         string
	validated    bool
	minDimension int
	dimension    int
}

func (s *AppIconSource) Validate() error {
	if s.validated {
		return nil
	}

	if s.Empty() {
		return fmt.Errorf("No URL or file location specified for icon")
	}

	if s.minDimension <= 0 {
		return fmt.Errorf("Minimum icon dimension invalid or not specified (%v)", s.minDimension)
	}

	loader := iconImageLoader{
		source: *s,
	}

	if err := loader.Validate(); err != nil {
		return err
	}

	s.validated = true

	return nil
}

func (s *AppIconSource) validateFile() error {
	if s.file == "" {
		return nil
	}

	return nil
}

func (s *AppIconSource) URL(url string) {
	s.url = url
	s.validated = false
}

func (s *AppIconSource) File(path string) {
	s.file = path
	s.validated = false
}

func (s AppIconSource) Empty() bool {
	return s.url == "" && s.file == ""
}

func (s *AppIconSource) Apply(from AppIconSource) {
	s.file = from.file
	s.url = from.url
	s.validated = from.validated
}

// MinDimension specifies the minimum dimension for the icon image.
func (s *AppIconSource) MinDimension(d int) {
	s.minDimension = d
	s.validated = false
}

func __icon() {
	// Generate all icons from a file or URL
	// Specify icons individually
	AppIcon("AppIcon", func(b *AppIconBuilder) {
		b.File("path/to/icon.png")
		b.URL("https://www.google.com/icon.png")

		b.Tablet().Phone().Mac().CarPlay().Watch()

		// b.IPhone(func(b *AppIconIPhone) {
		// 	b.Source.URL("https://www.google.com/icon.png")
		// 	b.Notification.File("path/to/icon.png")
		// })
	}).SaveTo("path/to/folder", true)
}
