package xcassets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestColorDefinition_Hex(t *testing.T) {
	d := ColorDefinition{}
	err := d.Hex("#262d44")
	assert.Nil(t, err)
	assert.Equal(t, "#262d44", d.hex)

	err = d.Hex("#foo")
	assert.NotNil(t, err)
}

func TestColorDefinition_Alpha(t *testing.T) {
	d := ColorDefinition{}
	err := d.Alpha(0.5)
	assert.Nil(t, err)
	assert.Equal(t, 0.5, d.alpha)

	err = d.Alpha(1.5)
	assert.NotNil(t, err)
}

func TestColorDefinition_White(t *testing.T) {
	d := ColorDefinition{}
	d.ColorSpace.GrayGamma22()

	err := d.White(0.5)
	assert.Nil(t, err)
	assert.Equal(t, 0.5, d.white)

	err = d.White(1.5)
	assert.NotNil(t, err)

	d.ColorSpace.SRGB()
	err = d.White(0.5)
	assert.NotNil(t, err)
}

func TestColorDefinition_RGB(t *testing.T) {
	type fields struct {
		build func() error
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Valid RGB should succeed",
			fields: fields{
				build: func() error {
					d := ColorDefinition{}
					return d.RGB(1, 1, 1)
				},
			},
			wantErr: false,
		},
		{
			name: "Invalid red should fail",
			fields: fields{
				build: func() error {
					d := ColorDefinition{}
					return d.RGB(-2, 1, 1)
				},
			},
			wantErr: true,
		},
		{
			name: "Invalid green should fail",
			fields: fields{
				build: func() error {
					d := ColorDefinition{}
					return d.RGB(1, -2, 1)
				},
			},
			wantErr: true,
		},
		{
			name: "Invalid blue should fail",
			fields: fields{
				build: func() error {
					d := ColorDefinition{}
					return d.RGB(1, 1, -2)
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.fields.build()
			if tt.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestColorDefinition_RGBFloat(t *testing.T) {
	type fields struct {
		build func() error
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Valid RGB should succeed",
			fields: fields{
				build: func() error {
					d := ColorDefinition{}
					return d.RGBFloat(1, 1, 1)
				},
			},
			wantErr: false,
		},
		{
			name: "Invalid red should fail",
			fields: fields{
				build: func() error {
					d := ColorDefinition{}
					return d.RGBFloat(-2, 1, 1)
				},
			},
			wantErr: true,
		},
		{
			name: "Invalid green should fail",
			fields: fields{
				build: func() error {
					d := ColorDefinition{}
					return d.RGBFloat(1, -2, 1)
				},
			},
			wantErr: true,
		},
		{
			name: "Invalid blue should fail",
			fields: fields{
				build: func() error {
					d := ColorDefinition{}
					return d.RGBFloat(1, 1, -2)
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.fields.build()
			if tt.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestColorDefinition_colorPresent(t *testing.T) {
	type fields struct {
		build func() *ColorDefinition
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Empty color should return false",
			fields: fields{
				build: func() *ColorDefinition {
					d := &ColorDefinition{}
					return d
				},
			},
			want: false,
		},
		{
			name: "Hex color should return true",
			fields: fields{
				build: func() *ColorDefinition {
					d := &ColorDefinition{}
					d.Hex("#262d44")
					return d
				},
			},
			want: true,
		},
		{
			name: "RGB color should return true",
			fields: fields{
				build: func() *ColorDefinition {
					d := &ColorDefinition{}
					d.RGB(100, 0, 0)
					return d
				},
			},
			want: true,
		},
		{
			name: "RGB color should return true",
			fields: fields{
				build: func() *ColorDefinition {
					d := &ColorDefinition{}
					d.ColorSpace.GrayGamma22()
					d.White(0.5)
					return d
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.fields.build()
			if tt.want {
				assert.True(t, d.colorPresent())
			} else {
				assert.False(t, d.colorPresent())

			}
		})
	}
}

func TestColorDefinition_detectOverlap(t *testing.T) {
	same := &ColorDefinition{}

	type fields struct {
		first  func() *ColorDefinition
		second func() *ColorDefinition
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Identical definitions should return no errors",
			fields: fields{
				first: func() *ColorDefinition {
					return same
				},
				second: func() *ColorDefinition {
					return same
				},
			},
			wantErr: false,
		},
		{
			name: "Unique definitions should return no errors",
			fields: fields{
				first: func() *ColorDefinition {
					d := &ColorDefinition{}
					d.Devices.Universal()
					return d
				},
				second: func() *ColorDefinition {
					d := &ColorDefinition{}
					d.Devices.IPhone()
					return d
				},
			},
			wantErr: false,
		},
		{
			name: "Intersecting devices should return an error",
			fields: fields{
				first: func() *ColorDefinition {
					d := &ColorDefinition{}
					d.Devices.Universal()
					return d
				},
				second: func() *ColorDefinition {
					d := &ColorDefinition{}
					d.Devices.Universal()
					return d
				},
			},
			wantErr: true,
		},
		{
			name: "Intersecting appearances should return an error",
			fields: fields{
				first: func() *ColorDefinition {
					d := &ColorDefinition{}
					d.Appearance.Any()
					return d
				},
				second: func() *ColorDefinition {
					d := &ColorDefinition{}
					d.Appearance.Any()
					return d
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d1 := tt.fields.first()
			d2 := tt.fields.second()
			err := d1.detectOverlap(d2)
			if tt.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestColorDefinition_Validate(t *testing.T) {
	type fields struct {
		build func() *ColorDefinition
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Valid definition should return no errors",
			fields: fields{
				build: func() *ColorDefinition {
					d := &ColorDefinition{}
					d.Hex("#262d44")
					d.Devices.Universal()
					d.Appearance.Any()
					return d
				},
			},
			wantErr: false,
		},
		{
			name: "Definition with no color should return an error",
			fields: fields{
				build: func() *ColorDefinition {
					d := &ColorDefinition{}
					d.Devices.Universal()
					d.Appearance.Any()
					return d
				},
			},
			wantErr: true,
		},
		{
			name: "Definition with no devices should return an error",
			fields: fields{
				build: func() *ColorDefinition {
					d := &ColorDefinition{}
					d.Hex("#262d44")
					d.Appearance.Any()
					return d
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.fields.build()
			err := d.Validate()
			if tt.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
