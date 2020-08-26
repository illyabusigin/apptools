package plist

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"sync"

	"howett.net/plist"
)

var (
	// ErrMissingRequiredProperty is the error returned for missing properties
	ErrMissingRequiredProperty = errors.New("Missing property")
	errMissingProperty         = func(p string) error {
		return fmt.Errorf("%w: %v", ErrMissingRequiredProperty, p)
	}
)

// PropertyList is a functional builder for your Info.plist. With a
// `PropertyList` struct you can easily declare your properties, device
// capabilities, privacy settings and more and quickly generate a string or
// write your property list to a file.
//
// This struct also contains help methods to populate default properties to save
// you time. PropertyList implements `io.Writer`.
type PropertyList struct {
	platform Platform

	developmentRegion       string
	bundleIdentifier        string
	executableFile          string
	version                 string
	bundleName              string
	packageType             string
	applicationVersionShort string
	bundleVersion           string
	requiresIphoneEnv       bool
	launchStoryboardName    string
	mainStoryboardName      string

	viewControllerBasedStatusBarAppearance bool

	statusBarStyle     string
	statusBarHidden    bool
	clientID           string
	endpointURL        string
	displayName        string
	ats                *AppTransportSecurity
	orientation        *Orientations
	tabletOrientations *Orientations
	privacy            *Privacy
	capabilities       *DeviceCapabilities
	scene              *SceneManifest

	skipValidation bool

	custom map[string]interface{}
	once   sync.Once
}

func (p *PropertyList) init() {
	p.once.Do(func() {
		if p.custom == nil {
			p.custom = map[string]interface{}{}
		}
	})
}

// New returns a new plist.PropertyList builder.
func New(platform Platform) *PropertyList {
	p := PropertyList{
		platform: platform,
	}

	p.init()

	return &p
}

// Defaults specifies a set of property list defaults for mobile devices.
// A list of the specified defaults can be found below:
//  plist.DevelopmentRegion("$(DEVELOPMENT_LANGUAGE)")
//  plist.BundleID("$(PRODUCT_BUNDLE_IDENTIFIER)")
//  plist.ExecutableFile("$(EXECUTABLE_NAME)")
//  plist.InfoDictionaryVersion("6.0")
//  plist.PackageType("APPLE")
//  plist.VersionShort("$(MARKETING_VERSION)")
//  plist.Version("1")
//  plist.RequiresIOS()
//  plist.MainStoryboard("Main")
//  plist.ViewControllerBasedStatusBarAppearance(true)
//  plist.StatusBarStyleDefault()
//  plist.AppTransportSecurity(func(s *AppTransportSecurity) {
// 	 s.AllowArbitraryLoads(true)
//  })
func (p *PropertyList) Defaults() {
	p.DevelopmentRegion("$(DEVELOPMENT_LANGUAGE)")
	p.BundleID("$(PRODUCT_BUNDLE_IDENTIFIER)")
	p.ExecutableFile("$(EXECUTABLE_NAME)")
	p.InfoDictionaryVersion("6.0")
	p.PackageType("APPLE")
	p.VersionShort("$(MARKETING_VERSION)")
	p.Version("1")
	p.RequiresIOS()
	p.MainStoryboard("Main")
	p.ViewControllerBasedStatusBarAppearance(true)
	p.StatusBarStyleDefault()
	p.StatusBarHidden(false)
	p.Capabilities(func(c *DeviceCapabilities) {
		c.ARMv7()
	})
	p.AppTransportSecurity(func(s *AppTransportSecurity) {
		s.AllowArbitraryLoads(true)
	})
	p.Orientations(func(o *Orientations) {
		o.Portrait()
	})
	// TODO: Scene manifest struct?
}

// SkipValidation will skip validation of required fields.
func (p *PropertyList) SkipValidation() *PropertyList {
	p.skipValidation = true
	return p
}

// Validate will validate the specified propery list configuration and return
// relevant and detailed errors for any issues discovered.
func (p *PropertyList) Validate() error {
	if p.bundleIdentifier == "" {
		return errMissingProperty(fmt.Sprintf("BundleID (CFBundleIdentifier)"))
	}

	if p.bundleName == "" {
		return errMissingProperty(fmt.Sprintf("BundleName (CFBundleName)"))
	}

	if p.statusBarStyle == "" {
		return errMissingProperty(fmt.Sprintf("StatusBarStyle (UIStatusBarStyle)"))
	}

	if p.displayName == "" {
		return errMissingProperty(fmt.Sprintf("DisplayName (CFBundleDisplayName)"))
	}

	if p.developmentRegion == "" {
		return errMissingProperty(fmt.Sprintf("DevelopmentRegion (CFBundleDevelopmentRegion)"))
	}

	if p.executableFile == "" {
		return errMissingProperty(fmt.Sprintf("ExecutableFile (CFBundleExecutable)"))
	}

	if p.version == "" {
		return errMissingProperty(fmt.Sprintf("InfoDictionaryVersion (CFBundleInfoDictionaryVersion)"))
	}

	if p.packageType == "" {
		return errMissingProperty(fmt.Sprintf("PackageType (CFBundlePackageType)"))
	}

	if p.applicationVersionShort == "" {
		return errMissingProperty(fmt.Sprintf("VersionShort (CFBundleShortVersionString)"))
	}

	if p.bundleVersion == "" {
		return errMissingProperty(fmt.Sprintf("Version (CFBundleVersion)"))
	}

	if p.ats == nil {
		return errMissingProperty(fmt.Sprintf("AppTransportSecurity (NSAppTransportSecurity)"))
	}

	if p.statusBarStyle == "" {
		return errMissingProperty(fmt.Sprintf("StatusBarStyle (UIStatusBarStyle)"))
	}

	if p.scene == nil {
		return errMissingProperty(fmt.Sprintf("SceneManifest (UIApplicationSceneManifest)"))
	}

	return nil
}

// DisplayName specifies the user-visible name of the bundle; used by Siri and
// visible on the Home screen in iOS.
// See https://developer.apple.com/library/archive/documentation/General/Reference/InfoPlistKeyReference/Articles/CoreFoundationKeys.html#//apple_ref/doc/uid/20001431-110725 for details.
func (p *PropertyList) DisplayName(n string) *PropertyList {
	p.displayName = n
	return p
}

// BundleID specifies an identifier string that specifies the app type of the
// bundle. The string should be in reverse DNS format using only the Roman
// alphabet in upper and lower case (A–Z, a–z), the dot (“.”), and the hyphen
// (“-”).
// See https://developer.apple.com/library/archive/documentation/General/Reference/InfoPlistKeyReference/Articles/CoreFoundationKeys.html#//apple_ref/doc/uid/20001431-102070 for details.
func (p *PropertyList) BundleID(id string) *PropertyList {
	p.bundleIdentifier = id
	return p
}

// DevelopmentRegion specifies the default language and region for the bundle,
// as a language ID.
// See https://developer.apple.com/documentation/bundleresources/information_property_list/cfbundledevelopmentregion for more information.
func (p *PropertyList) DevelopmentRegion(v string) *PropertyList {
	p.developmentRegion = v
	return p
}

// ExecutableFile specifies the name of the bundle’s executable file.
// See https://developer.apple.com/documentation/bundleresources/information_property_list/CFBundleExecutable for more information.
func (p *PropertyList) ExecutableFile(v string) *PropertyList {
	p.executableFile = v
	return p
}

// InfoDictionaryVersion sets the current version of the Information Property List structure.
// See https://developer.apple.com/documentation/bundleresources/information_property_list/CFBundleInfoDictionaryVersion for more information.
func (p *PropertyList) InfoDictionaryVersion(v string) *PropertyList {
	p.version = v
	return p
}

// BundleName specifies a user-visible short name for the bundle.
// See https://developer.apple.com/documentation/bundleresources/information_property_list/CFBundleName for more information.
func (p *PropertyList) BundleName(v string) *PropertyList {
	p.bundleName = v
	return p
}

// PackageType specifies the type of bundle.
// See https://developer.apple.com/documentation/bundleresources/information_property_list/CFBundlePackageType for more information.
func (p *PropertyList) PackageType(v string) *PropertyList {
	p.packageType = v
	return p
}

// VersionShort specifies the release or version number of the bundle.
// See https://developer.apple.com/documentation/bundleresources/information_property_list/CFBundleShortVersionString for more information.
func (p *PropertyList) VersionShort(v string) *PropertyList {
	p.applicationVersionShort = v
	return p
}

// Version specifies the version of the build that identifies an iteration of
// the bundle. Should follow semantic versioning.
// See https://developer.apple.com/documentation/bundleresources/information_property_list/CFBundleVersion for more information.
func (p *PropertyList) Version(v string) *PropertyList {
	p.bundleVersion = v
	return p
}

// RequiresIOS specifices a true boolean value indicating whether the app must run in iOS.
// See https://developer.apple.com/documentation/bundleresources/information_property_list/LSRequiresIPhoneOS for more information.
func (p *PropertyList) RequiresIOS() *PropertyList {
	p.requiresIphoneEnv = true
	return p
}

// StatusBarStyleDefault uses a dark status bar, intended for use on light
// backgrounds.
func (p *PropertyList) StatusBarStyleDefault() *PropertyList {
	p.statusBarStyle = "UIStatusBarStyleDefault"
	return p
}

// StatusBarStyleLightContent uses a light status bar, intended for use on dark
// backgrounds.
func (p *PropertyList) StatusBarStyleLightContent() *PropertyList {
	p.statusBarStyle = "UIStatusBarStyleLightContent"
	return p
}

// StatusBarStyleDarkContent uses a light status bar, intended for use on light
// backgrounds.
func (p *PropertyList) StatusBarStyleDarkContent() *PropertyList {
	p.statusBarStyle = "UIStatusBarStyleDarkContent"
	return p
}

// StatusBarHidden specifies a boolean value indicating whether the status bar
// is initially hidden when the app launches.
func (p *PropertyList) StatusBarHidden(v bool) *PropertyList {
	p.statusBarHidden = v
	return p
}

// ViewControllerBasedStatusBarAppearance specifies a boolean value indicating
// whether the status bar appearance is based on the style preferred for the
// current view controller.
// See https://developer.apple.com/documentation/bundleresources/information_property_list/uiviewcontrollerbasedstatusbarappearance for more information.
func (p *PropertyList) ViewControllerBasedStatusBarAppearance(v bool) *PropertyList {
	p.viewControllerBasedStatusBarAppearance = true
	return p
}

// LaunchScreenStoryboard specifies the  filename of the storyboard from which
// to generate the app’s launch image.
func (p *PropertyList) LaunchScreenStoryboard(storyboard string) *PropertyList {
	p.launchStoryboardName = storyboard
	return p
}

// MainStoryboard specifies the name of the app's main storyboard file.
// See https://developer.apple.com/documentation/bundleresources/information_property_list/uimainstoryboardfile for more information.
func (p *PropertyList) MainStoryboard(storyboard string) *PropertyList {
	p.mainStoryboardName = storyboard
	return p
}

// SceneManifest specifies the scene manifest for the application.
func (p *PropertyList) SceneManifest(f func(m *SceneManifest)) *PropertyList {
	p.scene = &SceneManifest{}
	f(p.scene)
	return p
}

// AppTransportSecurity allows you to specify App Transport Security (ATS).
// See https://developer.apple.com/documentation/bundleresources/information_property_list/NSAppTransportSecurity for more information.
func (p *PropertyList) AppTransportSecurity(f func(s *AppTransportSecurity)) *PropertyList {
	p.ats = &AppTransportSecurity{}
	f(p.ats)
	return p
}

// Orientations allows you to specify the initial orientation of the app’s user
// interface.
func (p *PropertyList) Orientations(f func(o *Orientations)) *PropertyList {
	p.orientation = &Orientations{}
	f(p.orientation)
	return p
}

// TabletOrientations allows you to specify the initial orientation of the app’s
// user interface when running on an iPad.
func (p *PropertyList) TabletOrientations(f func(o *Orientations)) *PropertyList {
	p.tabletOrientations = &Orientations{}
	f(p.tabletOrientations)
	return p
}

// Privacy allows you to specify privacy request text for a variety of device
// permissions.
func (p *PropertyList) Privacy(f func(p *Privacy)) *PropertyList {
	p.privacy = &Privacy{}
	f(p.privacy)
	return p
}

// Capabilities specifies the device-related features that your app requires
// to run. This is a required field.
// See https://developer.apple.com/documentation/bundleresources/information_property_list/uirequireddevicecapabilities for more information.
func (p *PropertyList) Capabilities(f func(c *DeviceCapabilities)) *PropertyList {
	p.capabilities = &DeviceCapabilities{}
	f(p.capabilities)
	return p
}

// Set will set an arbitrary key-value pair in the Info.plist. Key set in this
// manner will override any keys set by any of the builder functions.
func (p *PropertyList) Set(key string, value interface{}) *PropertyList {
	p.init()

	p.custom[key] = value

	return p
}

// Build will build the Info.plist
func (p *PropertyList) Build() (string, error) {
	if !p.skipValidation {
		if err := p.Validate(); err != nil {
			return "", err
		}
	}

	buf := bytes.Buffer{}

	data := map[string]interface{}{
		keyCFBundleIdentifier:            p.bundleIdentifier,
		keyCFBundleDisplayName:           p.displayName,
		keyCFBundleDevelopmentRegion:     p.developmentRegion,
		keyCFBundleExecutable:            p.executableFile,
		keyCFBundleInfoDictionaryVersion: p.version,
		keyCFBundleName:                  p.bundleName,
		keyCFBundlePackageType:           p.packageType,
		keyCFBundleShortVersionString:    p.applicationVersionShort,
		keyCFBundleVersion:               p.bundleVersion,
		keyLSRequiresIPhoneOS:            p.requiresIphoneEnv,
		keyUIStatusBarHidden:             p.statusBarHidden,
		keyUIStatusBarStyle:              p.statusBarStyle,
	}

	if scene := p.scene; scene != nil {
		sceneManifest := scene.build()
		data["UIApplicationSceneManifest"] = sceneManifest
	}

	// Custom keys are always applied last, overriding any builder keys
	for key, value := range p.custom{
		data[key] = value
	}

	encoder := plist.NewEncoder(&buf)
	if err := encoder.Encode(data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

// Write the property list to the specified io.Writer.
func (p *PropertyList) Write(w io.Writer) error {
	data, err := p.Build()
	if err != nil {
		return err
	}

	_, err = w.Write([]byte(data))

	return err
}
