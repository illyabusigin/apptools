package plist

const (
	// PLIST
	keyUISupportedInterfaceOrientations     = "UISupportedInterfaceOrientations"
	keyUISupportedInterfaceOrientationsIPad = "UISupportedInterfaceOrientations~ipad"
	keyUIStatusBarHidden                    = "UIStatusBarHidden"
	keyUIStatusBarStyle                     = "UIStatusBarStyle"
	keyCFBundleDisplayName                  = "CFBundleDisplayName"
	keyCFBundleName                         = "CFBundleName"
	keyCFBundleIdentifier                   = "CFBundleIdentifier"
	keyCFBundleDevelopmentRegion            = "CFBundleDevelopmentRegion"
	keyCFBundleExecutable                   = "CFBundleExecutable"
	keyCFBundleInfoDictionaryVersion        = "CFBundleInfoDictionaryVersion"
	keyCFBundlePackageType                  = "CFBundlePackageType"
	keyCFBundleShortVersionString           = "CFBundleShortVersionString"
	keyCFBundleVersion                      = "CFBundleVersion"
	keyLSRequiresIPhoneOS                   = "LSRequiresIPhoneOS"
	keyNSAppTransportSecurity               = "NSAppTransportSecurity"

	// ATS
	atsNSAllowsArbitraryLoads             = "NSAllowsArbitraryLoads"
	atsNSAllowsArbitraryLoadsForMedia     = "NSAllowsArbitraryLoadsForMedia"
	atsNSAllowsArbitraryLoadsInWebContent = "NSAllowsArbitraryLoadsInWebContent"
	atsNSAllowsLocalNetworking            = "NSAllowsLocalNetworking"
	atsNSExceptionDomains                 = "NSExceptionDomains"
)

// Platform represents the targeted platform
type Platform string

const (
	// PlatformIOS is for validating the plist against the iOS plaftform
	PlatformIOS Platform = "ios"

	// PlatformMac is for validating the plist against the Mac/OSX plaftform
	PlatformMac Platform = "mac"
)

var privacyKeys = map[string]string{
	"appleEvents":             "NSAppleEventsUsageDescription",
	"fileProviderDomain":      "NSFileProviderDomainUsageDescription",
	"calendar":                "NSCalendarsUsageDescription",
	"camera":                  "NSCameraUsageDescription",
	"bluetoothAlways":         "NSBluetoothAlwaysUsageDescription",
	"bluetoothPeripheral":     "NSBluetoothPeripheralUsageDescription",
	"contacts":                "NSContactsUsageDescription",
	"callkit":                 "NSVoIPUsageDescription",
	"desktopFolder":           "NSDesktopFolderUsageDescription",
	"documentsFolder":         "NSDocumentsFolderUsageDescription",
	"downloadsFolder":         "NSDownloadsFolderUsageDescription",
	"driverExtension":         "OSBundleUsageDescription",
	"faceID":                  "NSFaceIDUsageDescription",
	"fileProviderPresent":     "NSFileProviderPresenceUsageDescription",
	"healthRecordsUsage":      "NSHealthClinicalHealthRecordsShareUsageDescription",
	"healthRecordsShareUsage": "NSHealthShareUsageDescription",
	"healthRecordUpdateUsage": "NSHealthUpdateUsageDescription",
	"homeKit":                 "NSHomeKitUsageDescription",
	"locationEverything":      "NSLocationAlwaysAndWhenInUseUsageDescription",
	"locationAlways":          "NSLocationAlwaysUsageDescription",
	"locationUsage":           "NSLocationUsageDescription",
	"locationWhenInUse":       "NSLocationWhenInUseUsageDescription",
	"appleMusic":              "NSAppleMusicUsageDescription",
	"microphone":              "NSMicrophoneUsageDescription",
	"motion":                  "NSMotionUsageDescription",
	"music":                   "kTCCServiceMediaLibrary",
	"networkVolumes":          "NSNetworkVolumesUsageDescription",
	"nfcScan":                 "NFCReaderUsageDescription",
	"photoLibraryAdd":         "NSPhotoLibraryAddUsageDescription",
	"photoLibrayUsage":        "NSPhotoLibraryUsageDescription",
	"reminders":               "NSRemindersUsageDescription",
	"removableStorage":        "NSRemovableVolumesUsageDescription",
	"siri":                    "NSSiriUsageDescription",
	"speechRecognition":       "NSSpeechRecognitionUsageDescription",
	"systemAdminstration":     "NSSystemAdministrationUsageDescription",
	"systemExtension":         "NSSystemExtensionUsageDescription",
	"tvProvider":              "NSVideoSubscriberAccountUsageDescription",
}
