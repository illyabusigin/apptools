package xcassets

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestColorBuilder_Build(t *testing.T) {
	type fields struct {
		build func() *ColorBuilder
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "Valid color definition should return no errors",
			fields: fields{
				build: func() *ColorBuilder {
					b := Color("TestColor", func(b *ColorBuilder) {
						b.Color(func(d *ColorDefinition) {
							d.RGBFloat(1, 0, 0)
							d.ColorSpace.SRGB()
							d.Devices.Universal()
						})
					})

					return b
				},
			},
			want:    `fixture1.json`,
			wantErr: false,
		},
		{
			name: "Valid color definition should return no errors",
			fields: fields{
				build: func() *ColorBuilder {
					b := Color("TestColor", func(b *ColorBuilder) {
						b.Gamut.SRGBAndDisplayP3()
						b.Color(func(d *ColorDefinition) {
							d.RGBFloat(1, 0, 0)
							d.ColorSpace.SRGB()
							d.Devices.Universal()
						})
					})

					return b
				},
			},
			want:    `fixture2.json`,
			wantErr: false,
		},
		{
			name: "Invalid color definition should return an error",
			fields: fields{
				build: func() *ColorBuilder {
					b := Color("TestColor", func(b *ColorBuilder) {
						b.Gamut.SRGBAndDisplayP3()
						b.Color(func(d *ColorDefinition) {
							d.RGBFloat(1, 0, 0)
							d.ColorSpace.SRGB()
						})
					})

					return b
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			colorBuilder := tt.fields.build()
			out, err := colorBuilder.Build()

			if tt.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)

				expected, err := ioutil.ReadFile("testdata/" + tt.want)
				assert.Nil(t, err)

				if !assert.JSONEq(t, string(expected), out) {
					// fmt.Println("out:", out) // useful for debugging JSON
				}
			}

		})
	}
}

func TestColorBuilder_Validate(t *testing.T) {
	type fields struct {
		build func() *ColorBuilder
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Empty ColorBuilder should NOT return an error",
			fields: fields{
				build: func() *ColorBuilder {
					b := Color("TestColor", func(b *ColorBuilder) {
						b.Color(func(d *ColorDefinition) {
							d.Devices.Universal()
							d.Hex("#ff0000")
						})
					})
					return b
				},
			},
			wantErr: false,
		},
		{
			name: "Empty ColorBuilder should return an error",
			fields: fields{
				build: func() *ColorBuilder {
					b := Color("test", func(b *ColorBuilder) {

					})
					return b
				},
			},
			wantErr: true,
		},
		{
			name: "Invalid color definition should throw an error",
			fields: fields{
				build: func() *ColorBuilder {
					b := Color("test", func(b *ColorBuilder) {
						b.Color(func(d *ColorDefinition) {
							d.Hex("foo")
						})
					})
					return b
				},
			},
			wantErr: true,
		},
		{
			name: "Overlapping color definitions should throw an error",
			fields: fields{
				build: func() *ColorBuilder {
					b := Color("test", func(b *ColorBuilder) {
						b.Color(func(d *ColorDefinition) {
							d.Hex("#ff0000")
							d.Devices.Universal()
						})
						b.Color(func(d *ColorDefinition) {
							d.Hex("#ff0000")
							d.Devices.Universal()
						})
					})
					return b
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := tt.fields.build()
			err := b.Validate()
			if tt.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
