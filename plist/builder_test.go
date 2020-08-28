package plist

import (
	"errors"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func _TestPlistCanary(t *testing.T) {
	plist := New(PlatformIOS)
	plist.SkipValidation()
	plist.DisplayName("BestApp")
	plist.BundleID("com.best.app")
	plist.StatusBarStyleDarkContent()
	plist.StatusBarStyleDefault()
	plist.StatusBarStyleLightContent()
	plist.ViewControllerBasedStatusBarAppearance(true)
	plist.StatusBarHidden(true)
	plist.LaunchScreenStoryboard("Launch")
	plist.MainStoryboard("Main")
	plist.AppTransportSecurity(func(s *AppTransportSecurity) {
		s.AllowArbitraryLoads(true)
	})
	plist.Orientations(func(o *Orientations) {
		o.Portrait()
		o.LandscapeLeft()
		o.LandscapeRight()
		o.UpsideDown()
	})
	plist.TabletOrientations(func(o *Orientations) {
		o.Portrait()
		o.LandscapeLeft()
		o.LandscapeRight()
		o.UpsideDown()
	})
	plist.Privacy(func(p *Privacy) {
		p.Calendar("Let me use your calendar")
	})

	_, err := plist.Build()
	assert.Nil(t, err)
}

func TestPropertyList_Validate(t *testing.T) {
	type fields struct {
		builder func() *PropertyList
	}
	tests := []struct {
		name   string
		fields fields
		want   error
	}{
		{
			name: "Missing bundle ID should fail validation",
			fields: fields{
				builder: func() *PropertyList {
					plist := &PropertyList{}
					return plist
				},
			},
			want: errMissingProperty("BundleID"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			plist := tt.fields.builder()
			err := plist.Validate()

			if tt.want != nil {
				assert.True(t, errors.Is(err, ErrMissingRequiredProperty))
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestPropertyList_Properties(t *testing.T) {
	type fields struct {
		builder func() *PropertyList
	}
	tests := []struct {
		name   string
		fields fields
		want   *PropertyList
	}{
		{
			name: "Setting bundle ID should save",
			fields: fields{
				builder: func() *PropertyList {
					plist := &PropertyList{}
					plist.BundleID("com.best.app")
					return plist
				},
			},
			want: &PropertyList{
				bundleIdentifier: "com.best.app",
			},
		},
		{
			name: "Skip validation should save",
			fields: fields{
				builder: func() *PropertyList {
					plist := &PropertyList{}
					plist.SkipValidation()
					return plist
				},
			},
			want: &PropertyList{
				skipValidation: true,
			},
		},
		{
			name: "Setting display name should save",
			fields: fields{
				builder: func() *PropertyList {
					plist := &PropertyList{}
					plist.DisplayName("foo")
					return plist
				},
			},
			want: &PropertyList{
				displayName: "foo",
			},
		},
		{
			name: "Setting development region should save",
			fields: fields{
				builder: func() *PropertyList {
					plist := &PropertyList{}
					plist.DevelopmentRegion("foo")
					return plist
				},
			},
			want: &PropertyList{
				developmentRegion: "foo",
			},
		},
		{
			name: "Setting status bar style to default should save",
			fields: fields{
				builder: func() *PropertyList {
					plist := &PropertyList{}
					plist.StatusBarStyleDefault()
					return plist
				},
			},
			want: &PropertyList{
				statusBarStyle: "UIStatusBarStyleDefault",
			},
		},
		{
			name: "Setting status bar style to light content should save",
			fields: fields{
				builder: func() *PropertyList {
					plist := &PropertyList{}
					plist.StatusBarStyleLightContent()
					return plist
				},
			},
			want: &PropertyList{
				statusBarStyle: "UIStatusBarStyleLightContent",
			},
		},
		{
			name: "Setting status bar style to dark content should save",
			fields: fields{
				builder: func() *PropertyList {
					plist := &PropertyList{}
					plist.StatusBarStyleDarkContent()
					return plist
				},
			},
			want: &PropertyList{
				statusBarStyle: "UIStatusBarStyleDarkContent",
			},
		},
		{
			name: "Setting executable file should save",
			fields: fields{
				builder: func() *PropertyList {
					plist := &PropertyList{}
					plist.ExecutableFile("foo")
					return plist
				},
			},
			want: &PropertyList{
				executableFile: "foo",
			},
		},
		{
			name: "Setting bundle name should save",
			fields: fields{
				builder: func() *PropertyList {
					plist := &PropertyList{}
					plist.BundleName("foo")
					return plist
				},
			},
			want: &PropertyList{
				bundleName: "foo",
			},
		},
		{
			name: "Setting package type should save",
			fields: fields{
				builder: func() *PropertyList {
					plist := &PropertyList{}
					plist.PackageType("foo")
					return plist
				},
			},
			want: &PropertyList{
				packageType: "foo",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			plist := tt.fields.builder()
			assert.Equal(t, tt.want, plist)
		})
	}
}
