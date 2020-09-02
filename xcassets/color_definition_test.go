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
