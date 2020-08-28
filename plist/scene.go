package plist

import (
	"fmt"
)

// SceneManifest contains configuration information about the app's scene-based
// life-cycle support.
type SceneManifest struct {
	multipleWindows bool
	application     *SceneConfiguration
	externalDisplay *SceneConfiguration
}

// Validate will validate the SceneManifest configuration and return any
// errors found.
func (m *SceneManifest) Validate() error {
	if m.application == nil {
		return errMissingProperty(fmt.Sprintf("SceneManifest.Application (UISceneConfigurations)"))
	}

	if err := m.application.Validate(); err != nil {
		return err
	}

	return nil
}

// Apply will apply the scene manifest to the specified PropertyList.
func (m *SceneManifest) Apply(p *PropertyList) {
	p.data["UIApplicationSceneManifest"] = m.build()
}

func (m *SceneManifest) build() map[string]interface{} {
	data := map[string]interface{}{
		"UIApplicationSupportsMultipleScenes": m.multipleWindows,
	}

	configurations := map[string]interface{}{}

	if m.application != nil {
		application := m.application.build()
		configurations["UIWindowSceneSessionRoleApplication"] = application
	}

	if m.externalDisplay != nil {
		externalDisplay := m.externalDisplay.build()
		configurations["UIWindowSceneSessionRoleExternalDisplay"] = externalDisplay
	}

	data["UISceneConfigurations"] = configurations

	return data
}

// Application specifies the scenes that you use to display content on the
// device's main screen and respond to user interactions.
// See https://developer.apple.com/documentation/bundleresources/information_property_list/uiapplicationscenemanifest/uisceneconfigurations/uiwindowscenesessionroleapplication for more information.
func (m *SceneManifest) Application(f func(c *SceneConfiguration)) {
	m.application = &SceneConfiguration{}
	f(m.application)
}

// ExternalDisplay specifies the scenes that you use to display content on an
// externally connected display.
// See https://developer.apple.com/documentation/bundleresources/information_property_list/uiapplicationscenemanifest/uisceneconfigurations/uiwindowscenesessionroleexternaldisplay for more information.
func (m *SceneManifest) ExternalDisplay(f func(c *SceneConfiguration)) {
	m.externalDisplay = &SceneConfiguration{}
	f(m.externalDisplay)
}

// MultipleWindows specifies a boolean value indicating whether the app
// supports two or more scenes simultaneously.
// See https://developer.apple.com/documentation/bundleresources/information_property_list/uiapplicationscenemanifest/uiapplicationsupportsmultiplescenes for more information.
func (m *SceneManifest) MultipleWindows(v bool) {
	m.multipleWindows = v
}

// SceneConfiguration describes a UISceneConfiguration which contains
// information about the objects and storyboard for UKit to use when
// creating a particular scene.
// See https://developer.apple.com/documentation/uikit/uisceneconfiguration for more information.
type SceneConfiguration struct {
	name              string
	className         *string
	delegateClassName *string
	storyboardName    *string
}

// Validate will validate the SceneConfiguration, returning any errors.
func (c *SceneConfiguration) Validate() error {
	if c.name == "" {
		return errMissingProperty(fmt.Sprintf("SceneConfiguration.Name (UISceneConfigurationName)"))
	}

	return nil
}

func (c *SceneConfiguration) build() map[string]interface{} {
	data := map[string]interface{}{}

	data["UISceneConfigurationName"] = c.name

	if className := c.className; className != nil && *className != "" {
		data["UISceneClassName"] = className
	}

	if delegateClassName := c.delegateClassName; delegateClassName != nil && *delegateClassName != "" {
		data["UISceneDelegateClassName"] = delegateClassName
	}

	if storyboardName := c.storyboardName; storyboardName != nil && *storyboardName != "" {
		data["UISceneStoryboardFile"] = storyboardName
	}

	return data
}

// Name specifies the  app-specific name you use to identify the scene.
// See https://developer.apple.com/documentation/bundleresources/information_property_list/uiapplicationscenemanifest/uisceneconfigurations/uiwindowscenesessionroleapplication/uisceneconfigurationname for more information.
func (c *SceneConfiguration) Name(v string) {
	c.name = v
}

// ClassName specifies the name of the scene class you want UIKit to instantiate.
// See https://developer.apple.com/documentation/bundleresources/information_property_list/uiapplicationscenemanifest/uisceneconfigurations/uiwindowscenesessionroleapplication/uisceneclassname for more information.
func (c *SceneConfiguration) ClassName(v string) {
	c.className = &v
}

// DelegateClassName specifies the name of the app-specific class that you want
// UIKit to instantiate and use as the scene delegate object.
// See https://developer.apple.com/documentation/bundleresources/information_property_list/uiapplicationscenemanifest/uisceneconfigurations/uiwindowscenesessionroleapplication/uiscenedelegateclassname for more information.
func (c *SceneConfiguration) DelegateClassName(v string) {
	c.delegateClassName = &v
}

// Storyboard specifies the name of the storyboard file containing the scene's
// initial user interface.
// See https://developer.apple.com/documentation/bundleresources/information_property_list/uiapplicationscenemanifest/uisceneconfigurations/uiwindowscenesessionroleapplication/uiscenestoryboardfile for more information.
func (c *SceneConfiguration) Storyboard(v string) {
	c.storyboardName = &v
}

func _sceneManifest() {
	s := &SceneManifest{}
	s.MultipleWindows(false)
	s.Application(func(c *SceneConfiguration) {
		c.Name("Default")
		c.ClassName("Foo")
		c.DelegateClassName("Bar")
		c.Storyboard("Bar")
	})

	s.ExternalDisplay(func(c *SceneConfiguration) {
		c.Name("Default")
		c.ClassName("Foo")
		c.DelegateClassName("Bar")
		c.Storyboard("Bar")
	})
}
