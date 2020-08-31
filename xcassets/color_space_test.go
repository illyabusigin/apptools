package xcassets

import (
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestColorSpace_build(t *testing.T) {
	c := ColorSpace{}
	assert.Equal(t, "srgb", c.build())

	c.DisplayP3()
	assert.Equal(t, "display-p3", c.build())
}

func TestColorSpace_SRGB(t *testing.T) {
	type fields struct {
		builder func() *ColorSpace
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Empty color space should default to sRGB",
			fields: fields{
				builder: func() *ColorSpace {
					c := ColorSpace{}
					return &c
				},
			},
			want: "srgb",
		},
		{
			name: "sRGB color space should equal srgb",
			fields: fields{
				builder: func() *ColorSpace {
					c := ColorSpace{}
					c.SRGB()

					return &c
				},
			},
			want: "srgb",
		},
		{
			name: "DisplayP3 color space should equal display-p3",
			fields: fields{
				builder: func() *ColorSpace {
					c := ColorSpace{}
					c.DisplayP3()

					return &c
				},
			},
			want: "display-p3",
		},
		{
			name: "ExtendedRangeSRGB color space should equal extended-srgb",
			fields: fields{
				builder: func() *ColorSpace {
					c := ColorSpace{}
					c.ExtendedRangeSRGB()

					return &c
				},
			},
			want: "extended-srgb",
		},
		{
			name: "ExtendedRangeLinearSRGB color space should equal extended-linear-srgb",
			fields: fields{
				builder: func() *ColorSpace {
					c := ColorSpace{}
					c.ExtendedRangeLinearSRGB()

					return &c
				},
			},
			want: "extended-linear-srgb",
		},
		{
			name: "GrayGamma22 color space should equal gray-gamma-22",
			fields: fields{
				builder: func() *ColorSpace {
					c := ColorSpace{}
					c.GrayGamma22()

					return &c
				},
			},
			want: "gray-gamma-22",
		},
		{
			name: "ExtendedRangeGray color space should equal extended-gray",
			fields: fields{
				builder: func() *ColorSpace {
					c := ColorSpace{}
					c.ExtendedRangeGray()

					return &c
				},
			},
			want: "extended-gray",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.fields.builder()
			assert.Equal(t, tt.want, c.build())
		})
	}
}
