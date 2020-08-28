package plist

import (
	"sync"
)

// Privacy allows you to specify privacy request text for a variety of device
// permissions.
// See https://iosdevcenters.blogspot.com/2016/09/infoplist-privacy-settings-in-ios-10.html#comment-3531316086 for details.
type Privacy struct {
	values map[string]interface{}
	once   sync.Once
}

func (p *Privacy) init() {
	p.once.Do(func() {
		if p.values == nil {
			p.values = map[string]interface{}{}
		}
	})
}

// Set a custom privacy related key. This can be used if no built-in method
// is provided for your key type.
func (p *Privacy) Set(key string, value interface{}) {
	p.init()
	p.values[key] = value
}

// Apply will apply the privacy configuration to the provider property list.
func (p *Privacy) Apply(pl *PropertyList) {
	for key, value := range p.values {
		pl.data[key] = value
	}
}

type devicePermission struct {
	key, value string
}

// FileProviderDomain specifies a message that tells the user why the app
// needs access to files managed by a file provider.
//
// See https://developer.apple.com/documentation/bundleresources/information_property_list/nsappleeventsusagedescription for more information.
func (p *Privacy) FileProviderDomain(desc string) *Privacy {
	p.init()
	key := privacyKeys["fileProviderDomain"]
	p.values[key] = desc
	return p
}

// AppleEvents specifies a message that tells the user why the app is requesting
// the ability to send Apple events.
//
// See https://developer.apple.com/documentation/bundleresources/information_property_list/nsappleeventsusagedescription for more information.
func (p *Privacy) AppleEvents(desc string) *Privacy {
	p.init()
	key := privacyKeys["appleEvents"]
	p.values[key] = desc
	return p
}

// Calendar specifies a message that tells the user why the app is requesting access
// to the user’s calendar data.
//
// See https://developer.apple.com/documentation/bundleresources/information_property_list/NSCalendarsUsageDescription for more information.
func (p *Privacy) Calendar(desc string) *Privacy {
	p.init()
	key := privacyKeys["calendar"]
	p.values[key] = desc
	return p
}

// Camera specifies a message that tells the user why the app is requesting access
// to the device’s camera.
//
// See https://developer.apple.com/documentation/bundleresources/information_property_list/NSCameraUsageDescription for more information..
func (p *Privacy) Camera(desc string) *Privacy {
	p.init()
	key := privacyKeys["camera"]
	p.values[key] = desc
	return p
}

// BluetoothAlways specifies a message that tells the user why the app needs
// access to Bluetooth.
//
// See https://developer.apple.com/documentation/bundleresources/information_property_list/NSBluetoothAlwaysUsageDescription for more information.
func (p *Privacy) BluetoothAlways(desc string) *Privacy {
	p.init()
	key := privacyKeys["bluetoothAlways"]
	p.values[key] = desc
	return p
}

// BluetoothPeripheral specifies a message that tells the user why the app is
// requesting the ability to connect to Bluetooth peripherals.
//
// DEPRECATED: For apps with a deployment target of iOS 13 and later, use `BluetoothAlways` instead.
//
// See https://developer.apple.com/documentation/bundleresources/information_property_list/NSBluetoothPeripheralUsageDescription for more information.
func (p *Privacy) BluetoothPeripheral(desc string) *Privacy {
	p.init()
	key := privacyKeys["bluetoothPeripheral"]
	p.values[key] = desc
	return p
}

// Contacts specifies a mmessage that tells the user why the app is requesting
// access to the user’s contacts.
//
// See https://developer.apple.com/documentation/bundleresources/information_property_list/NSContactsUsageDescription for more information.
func (p *Privacy) Contacts(desc string) *Privacy {
	p.init()
	key := privacyKeys["contacts"]
	p.values[key] = desc
	return p
}

// CallKit specifies a message that tells the user why the app is requesting
// the ability use VoIP and/or CallKit.
//
// No official Apple documentation was found for this preference.
func (p *Privacy) CallKit(desc string) *Privacy {
	p.init()
	key := privacyKeys["callkit"]
	p.values[key] = desc
	return p
}

// DesktopFolder specifies a message that tells the user why the app needs
// access to the user’s Desktop folder.
//
// See https://developer.apple.com/documentation/bundleresources/information_property_list/NSDesktopFolderUsageDescription for more information.
func (p *Privacy) DesktopFolder(desc string) *Privacy {
	p.init()
	key := privacyKeys["desktopFolder"]
	p.values[key] = desc
	return p
}

// DocumentsFolder specifies a  message that tells the user why the app needs
// access to the user’s Documents folder.
//
// See https://developer.apple.com/documentation/bundleresources/information_property_list/NSDocumentsFolderUsageDescription for more information.
func (p *Privacy) DocumentsFolder(desc string) *Privacy {
	p.init()
	key := privacyKeys["documentsFolder"]
	p.values[key] = desc
	return p
}

// DownloadsFolder specifies a message that tells the user why the app needs
// access to the user’s Downloads folder.
//
// See https://developer.apple.com/documentation/bundleresources/information_property_list/NSDownloadsFolderUsageDescription for more information.
func (p *Privacy) DownloadsFolder(desc string) *Privacy {
	p.init()
	key := privacyKeys["downloadsFolder"]
	p.values[key] = desc
	return p
}

// DriverExtension specifies a message that tells the user why the app needs
// access to implement a driver extension.
//
// See https://developer.apple.com/documentation/driverkit for more information.
func (p *Privacy) DriverExtension(desc string) *Privacy {
	p.init()
	key := privacyKeys["driverExtension"]
	p.values[key] = desc
	return p
}

// FaceID specifies a message that tells the user why the app is requesting the
// ability to authenticate with Face ID.
//
// See url for more information.
func (p *Privacy) FaceID(desc string) *Privacy {
	p.init()
	key := privacyKeys["faceID"]
	p.values[key] = desc
	return p
}

// FileProviderPresence specifies a message that tells the user why the app
// needs to be informed when other apps access files that it manages.
//
// See https://developer.apple.com/documentation/bundleresources/information_property_list/NSFileProviderPresenceUsageDescription for more information.
func (p *Privacy) FileProviderPresence(desc string) *Privacy {
	p.init()
	key := privacyKeys["fileProviderPresent"]
	p.values[key] = desc
	return p
}

// ReadClinicalHealthRecords specifies a message to the user that explains why
// the app requested permission to read clinical records.
//
// See https://developer.apple.com/documentation/bundleresources/information_property_list/NSHealthClinicalHealthRecordsShareUsageDescription for more information.
func (p *Privacy) ReadClinicalHealthRecords(desc string) *Privacy {
	p.init()
	key := privacyKeys["healthRecordsUsage"]
	p.values[key] = desc
	return p
}

// HealthKitRead specifies a message to the user that explains why the app
// requested permission to read samples from the HealthKit store.
//
// See https://developer.apple.com/documentation/bundleresources/information_property_list/NSHealthShareUsageDescription for more information.
func (p *Privacy) HealthKitRead(desc string) *Privacy {
	p.init()
	key := privacyKeys["healthRecordsShareUsage"]
	p.values[key] = desc
	return p
}

// HealthKitUpdate specifies a message to the user that explains why the app
// requested permission to save samples to the HealthKit store.
//
// See https://developer.apple.com/documentation/bundleresources/information_property_list/NSHealthUpdateUsageDescription for more information.
func (p *Privacy) HealthKitUpdate(desc string) *Privacy {
	p.init()
	key := privacyKeys["healthRecordUpdateUsage"]
	p.values[key] = desc
	return p
}

// HomeKit specifies a message that tells the user why the app is requesting
//  access to the user’s HomeKit configuration data.
//
// See https://developer.apple.com/documentation/bundleresources/information_property_list/NSHomeKitUsageDescription for more information.
func (p *Privacy) HomeKit(desc string) *Privacy {
	p.init()
	key := privacyKeys["homeKit"]
	p.values[key] = desc
	return p
}

// Location specifies a message that tells the user why the app is requesting
// access to the user’s location information at all times.
//
// See https://developer.apple.com/documentation/bundleresources/information_property_list/NSLocationAlwaysAndWhenInUseUsageDescription for more information.
func (p *Privacy) Location(desc string) *Privacy {
	p.init()
	key := privacyKeys["locationEverything"]
	p.values[key] = desc
	return p
}

// LocationAlways specifies a message that tells the user why the app is
// requesting access to the user's location at all times.
//
// **NOTE**: Deprecated. This key is required if your iOS app uses APIs that
// access the user’s location at all times and deploys to targets earlier
// than iOS 11.
//
// See https://developer.apple.com/documentation/bundleresources/information_property_list/NSLocationAlwaysUsageDescription for more information.
func (p *Privacy) LocationAlways(desc string) *Privacy {
	p.init()
	key := privacyKeys["locationAlways"]
	p.values[key] = desc
	return p
}

// LocationUsage specifies a message that tells the user why the app is
// requesting access to the user’s location information.
//
// **NOTE**: Deprecated since iOS 8. This key is required if your macOS app uses
// APIs that access the user’s location information.
//
// See https://developer.apple.com/documentation/bundleresources/information_property_list/NSLocationUsageDescription for more information.
func (p *Privacy) LocationUsage(desc string) *Privacy {
	p.init()
	key := privacyKeys["locationUsage"]
	p.values[key] = desc
	return p
}

// AppleMusic specifies a message that tells the user why the app is requesting
// access to the user’s media library.
//
// See https://developer.apple.com/documentation/bundleresources/information_property_list/NSAppleMusicUsageDescription for more information.
func (p *Privacy) AppleMusic(desc string) *Privacy {
	p.init()
	key := privacyKeys["appleMusic"]
	p.values[key] = desc
	return p
}

// Microphone specifies a message that tells the user why the app is requesting
// access to the device’s microphone.
//
// See https://developer.apple.com/documentation/bundleresources/information_property_list/NSMicrophoneUsageDescription for more information.
func (p *Privacy) Microphone(desc string) *Privacy {
	p.init()
	key := privacyKeys["microphone"]
	p.values[key] = desc
	return p
}

// Motion specifies a message that tells the user why your app is requesting
// access to the device’s motion data.
//
// See https://developer.apple.com/documentation/bundleresources/information_property_list/NSMotionUsageDescription for more information.
func (p *Privacy) Motion(desc string) *Privacy {
	p.init()
	key := privacyKeys["motion"]
	p.values[key] = desc
	return p
}

// MediaLibray specifies a message for accessing the device media library.
//
// No official documentation is known to exist for this key: `kTCCServiceMediaLibrary`.
func (p *Privacy) MediaLibray(desc string) *Privacy {
	p.init()
	key := privacyKeys["music"]
	p.values[key] = desc
	return p
}

// NetworkVolumes specifies a message that tells the user why the app needs
// access to files on a network volume.
//
// See https://developer.apple.com/documentation/bundleresources/information_property_list/NSNetworkVolumesUsageDescription for more information.
func (p *Privacy) NetworkVolumes(desc string) *Privacy {
	p.init()
	key := privacyKeys["networkVolumes"]
	p.values[key] = desc
	return p
}

// NFCRead specifies a message that tells the user why the app is requesting
// access to the device’s NFC hardware.
//
// See https://developer.apple.com/documentation/bundleresources/information_property_list/NFCReaderUsageDescription for more information.
func (p *Privacy) NFCRead(desc string) *Privacy {
	p.init()
	key := privacyKeys["nfcScan"]
	p.values[key] = desc
	return p
}

// PhotoLibraryWrite specifies a message that tells the user why the app is
// requesting write-only access to the user’s photo library.
//
// See https://developer.apple.com/documentation/bundleresources/information_property_list/NSPhotoLibraryAddUsageDescription for more information.
func (p *Privacy) PhotoLibraryWrite(desc string) *Privacy {
	p.init()
	key := privacyKeys["photoLibraryAdd"]
	p.values[key] = desc
	return p
}

// PhotoLibrary specifies a message that tells the user why the app is
// requesting access to the user’s photo library.
//
// See https://developer.apple.com/documentation/bundleresources/information_property_list/NSPhotoLibraryUsageDescription for more information.
func (p *Privacy) PhotoLibrary(desc string) *Privacy {
	p.init()
	key := privacyKeys["photoLibrayUsage"]
	p.values[key] = desc
	return p
}

// Reminders specifies a message that tells the user why the app is requesting
// access to the user’s reminders.
//
// See https://developer.apple.com/documentation/bundleresources/information_property_list/NSRemindersUsageDescription for more information.
func (p *Privacy) Reminders(desc string) *Privacy {
	p.init()
	key := privacyKeys["reminders"]
	p.values[key] = desc
	return p
}

// RemovableStorage specifies a message that tells the user why the app needs access to files on a removable volume.
//
// See https://developer.apple.com/documentation/bundleresources/information_property_list/NSRemovableVolumesUsageDescription for more information.
func (p *Privacy) RemovableStorage(desc string) *Privacy {
	p.init()
	key := privacyKeys["removableStorage"]
	p.values[key] = desc
	return p
}

// Siri specifies a message that tells the user why the app is requesting to
// send user data to Siri.
//
// See https://developer.apple.com/documentation/bundleresources/information_property_list/NSSiriUsageDescription for more information.
func (p *Privacy) Siri(desc string) *Privacy {
	p.init()
	key := privacyKeys["siri"]
	p.values[key] = desc
	return p
}

// SpeechRecognition specifies a message that tells the user why the app is
// requesting to send user data to Apple’s speech recognition servers.
//
// See https://developer.apple.com/documentation/bundleresources/information_property_list/NSSpeechRecognitionUsageDescription for more information.
func (p *Privacy) SpeechRecognition(desc string) *Privacy {
	p.init()
	key := privacyKeys["speechRecognition"]
	p.values[key] = desc
	return p
}

// SystemAdministration specifies a message in macOS that tells the user why
// the app is requesting to manipulate the system configuration.
//
// See https://developer.apple.com/documentation/bundleresources/information_property_list/NSSystemAdministrationUsageDescription for more information.
func (p *Privacy) SystemAdministration(desc string) *Privacy {
	p.init()
	key := privacyKeys["systemAdminstration"]
	p.values[key] = desc
	return p
}

// SystemExtension specifies a message that tells the user why the app is
// trying to install a system extension bundle.
//
// See https://developer.apple.com/documentation/systemextensions/nssystemextensionusagedescriptionkey for more information.
func (p *Privacy) SystemExtension(desc string) *Privacy {
	p.init()
	key := privacyKeys["systemExtension"]
	p.values[key] = desc
	return p
}

// TVProviderAccess specifies a message that tells the user why the app is
// requesting access to the user’s TV provider account.
//
// See https://developer.apple.com/documentation/bundleresources/information_property_list/NSVideoSubscriberAccountUsageDescription for more information.
func (p *Privacy) TVProviderAccess(desc string) *Privacy {
	p.init()
	key := privacyKeys["tvProvider"]
	p.values[key] = desc
	return p
}
