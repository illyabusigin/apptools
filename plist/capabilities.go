package plist

// DeviceCapabilities allows you to specify required device capabilities.
// See https://developer.apple.com/documentation/bundleresources/information_property_list/uirequireddevicecapabilities for more information.
type DeviceCapabilities struct {
	accelerometer     bool
	arKit             bool
	armv7             bool
	arm64             bool
	autoFocusCamera   bool
	bluetoothLE       bool
	cameraFlash       bool
	frontFacingCamera bool
	gameKit           bool
	gps               bool
	gyroscope         bool
	healthKit         bool
	minPerfA12        bool
	locationServices  bool
	magnetometer      bool
	metal             bool
	microphone        bool
	nfc               bool
	openGLES1         bool
	openGLES2         bool
	openGLES3         bool
	p2pBT             bool
	sms               bool
	stillCamera       bool
	telephony         bool
	videoCamera       bool
	wifi              bool
}

// Apply will apply the device capabilities to the specified property list
func (c *DeviceCapabilities) Apply(p *PropertyList) {

}

// Accelerometer ensure the presence of accelerometers. Available in iOS 3.0
// and later.
func (c *DeviceCapabilities) Accelerometer() *DeviceCapabilities {
	c.accelerometer = true
	return c
}

// ARKit ensures the support for ARKit. Available in iOS 11.0 and later.
func (c *DeviceCapabilities) ARKit() *DeviceCapabilities {
	c.arKit = true
	return c
}

// ARMv7 ensures compilation for the armv7 instruction set, or as a 32/64-bit
// universal app. Available in iOS 3.1 and later.
func (c *DeviceCapabilities) ARMv7() *DeviceCapabilities {
	c.armv7 = true
	return c
}

// ARM64 ensures compilation for the arm64 instruction set. Include this key
// for all 64-bit apps and embedded bundles, like extensions and frameworks.
// Available in iOS 8.0 and later.
func (c *DeviceCapabilities) ARM64() *DeviceCapabilities {
	c.arm64 = true
	return c
}

// AutoFocusCamera ensures autofocus capabilities in the device’s still camera.
// You might need to include this value if your app supports macro photography
//  or requires sharper images to perform certain image-processing tasks.
// Available in iOS 3.0 and later.
func (c *DeviceCapabilities) AutoFocusCamera() *DeviceCapabilities {
	c.autoFocusCamera = true
	return c
}

// Bluetooth ensures the presence of bluetooth low-energy hardware. Available
// in iOS 5.0 and later.
func (c *DeviceCapabilities) Bluetooth() *DeviceCapabilities {
	c.bluetoothLE = true
	return c
}

// CameraFlash ensures the presence of a camera flash. Available
// in iOS 3.0 and later.
func (c *DeviceCapabilities) CameraFlash() *DeviceCapabilities {
	c.cameraFlash = true
	return c
}

// FrontFacingCamera ensures the presence of a front-facing camera. Available
// in iOS 3.0 and later.
func (c *DeviceCapabilities) FrontFacingCamera() *DeviceCapabilities {
	c.frontFacingCamera = true
	return c
}

// GameKit ensures the device has access to the Game Center service. Available
// in iOS 4.1 and later.
func (c *DeviceCapabilities) GameKit() *DeviceCapabilities {
	c.gameKit = true
	return c
}

// GPS ensures the device has GPS (or AGPS) hardware for tracking locations. If
// you include this value, you should also include the location-services value.
// Available in iOS 3.0 and later.
func (c *DeviceCapabilities) GPS() *DeviceCapabilities {
	c.gps = true
	return c
}

// Gyroscope ensures hardware access to a gyroscope. Available in iOS 3.0 and
// later.
func (c *DeviceCapabilities) Gyroscope() *DeviceCapabilities {
	c.gyroscope = true
	return c
}

// HealthKit ensures the device has support for HealthKit. Available in
// iOS 8.0 and later.
func (c *DeviceCapabilities) HealthKit() *DeviceCapabilities {
	c.healthKit = true
	return c
}

// MinimumPerformanceA12 ensures the performance and capabilities of the A12
// Bionic and later chips. Available in iOS 12.0 and later.
func (c *DeviceCapabilities) MinimumPerformanceA12() *DeviceCapabilities {
	c.minPerfA12 = true
	return c
}

// LocationServices ensures access to the device’s current location using the
// Core Location framework. This value refers to the general location services
// feature. If you specifically need GPS-level accuracy, also include the GPS
// feature. Available in iOS 3.0 and later.
func (c *DeviceCapabilities) LocationServices() *DeviceCapabilities {
	c.locationServices = true
	return c
}

// Magnetometer ensures the presence of magnetometer hardware. Apps use this
// hardware to receive heading-related events through the Core Location
// framework. Available in iOS 3.0 and later.
func (c *DeviceCapabilities) Magnetometer() *DeviceCapabilities {
	c.magnetometer = true
	return c
}

// Metal ensures the device has support for graphics processing with Metal.
// Available in iOS 8.0 and later.
func (c *DeviceCapabilities) Metal() *DeviceCapabilities {
	c.metal = true
	return c
}

// Microphone ensures the device has access to the built-in microphone or
// accessories that provide a microphone. Available in iOS 3.0 and later.
func (c *DeviceCapabilities) Microphone() *DeviceCapabilities {
	c.microphone = true
	return c
}

// NFC ensures the device has support Near Field Communication (NFC) tag
// detection and access to messages that contain NFC Data Exchange Format data.
//  Use the Core NFC framework to detect and read NFC tags. Available in iOS
// 11.0 and later.
func (c *DeviceCapabilities) NFC() *DeviceCapabilities {
	c.nfc = true
	return c
}

// OpenGLES1 ensures the device  has support for OpenGL ES 1.1. Available in
// iOS 3.0 and later.
func (c *DeviceCapabilities) OpenGLES1() *DeviceCapabilities {
	c.openGLES1 = true
	return c
}

// OpenGLES2 ensures the device  has support for OpenGL ES 2.0. Available in
// iOS 3.0 and later.
func (c *DeviceCapabilities) OpenGLES2() *DeviceCapabilities {
	c.openGLES2 = true
	return c
}

// OpenGLES3 ensures the device  has support for OpenGL ES 3.0. Available in
// iOS 7.0 and later.
func (c *DeviceCapabilities) OpenGLES3() *DeviceCapabilities {
	c.openGLES3 = true
	return c
}

// PeerToPeerConnectivity ensures support for peer-to-peer connectivity over a
// Bluetooth network. Available in iOS 3.1 and later.
func (c *DeviceCapabilities) PeerToPeerConnectivity() *DeviceCapabilities {
	c.p2pBT = true
	return c
}

// SMS ensures the device has the Messages app. You might require this feature
//  if your app opens URLs with the sms scheme. Available in iOS 3.0 and later.
func (c *DeviceCapabilities) SMS() *DeviceCapabilities {
	c.sms = true
	return c
}

// StillCamera ensures the device has a camera. Available in iOS 3.0 and later.
func (c *DeviceCapabilities) StillCamera() *DeviceCapabilities {
	c.stillCamera = true
	return c
}

// Telephony ensures the device has the Phone app. Available in iOS 3.0 and
// later.
func (c *DeviceCapabilities) Telephony() *DeviceCapabilities {
	c.telephony = true
	return c
}

// VideoCamera ensures the device a camera with video capabilities on the
// device. Available in iOS 3.0 and later.
func (c *DeviceCapabilities) VideoCamera() *DeviceCapabilities {
	c.videoCamera = true
	return c
}

// WiFi ensures the device has access to networking features related to Wi-Fi
//  access. Available in iOS 3.0 and later.
func (c *DeviceCapabilities) WiFi() *DeviceCapabilities {
	c.wifi = true
	return c
}
