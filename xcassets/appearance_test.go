package xcassets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppearance_Options(t *testing.T) {
	type fields struct {
		build func() *Appearance
	}
	tests := []struct {
		name   string
		fields fields
		want   *Appearance
	}{
		{
			name: "Any should mark any",
			fields: fields{
				build: func() *Appearance {
					a := Appearance{}
					a.Any()
					return &a
				},
			},
			want: &Appearance{
				any: true,
			},
		},
		{
			name: "Light should mark light",
			fields: fields{
				build: func() *Appearance {
					a := Appearance{}
					a.Light()
					return &a
				},
			},
			want: &Appearance{
				light: true,
			},
		},
		{
			name: "Dark should mark any and light",
			fields: fields{
				build: func() *Appearance {
					a := Appearance{}
					a.Dark()
					return &a
				},
			},
			want: &Appearance{
				dark: true,
				any:  true,
			},
		},
		{
			name: "HighContrast should mark high contrast",
			fields: fields{
				build: func() *Appearance {
					a := Appearance{}
					a.HighContrast()
					return &a
				},
			},
			want: &Appearance{
				highContrast: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := tt.fields.build()
			assert.Equal(t, tt.want, a)
		})
	}
}

func TestAppearance_build(t *testing.T) {
	type fields struct {
		build func() *Appearance
	}
	tests := []struct {
		name   string
		fields fields
		want   appearances
	}{
		{
			name: "Empty appearance shoulde return empty appearances",
			fields: fields{
				build: func() *Appearance {
					a := Appearance{}
					return &a
				},
			},
			want: appearances{
				[]appearance{},
			},
		},
		{
			name: "Dark appearance shoulde should return correct appearances",
			fields: fields{
				build: func() *Appearance {
					a := Appearance{}
					a.Dark()
					return &a
				},
			},
			want: appearances{
				[]appearance{},
				[]appearance{
					{
						Appearance: "luminosity",
						Value:      "dark",
					},
				},
			},
		},
		{
			name: "Dark and light appearance shoulde should return correct appearances",
			fields: fields{
				build: func() *Appearance {
					a := Appearance{}
					a.Dark()
					a.Light()
					return &a
				},
			},
			want: appearances{
				[]appearance{},
				[]appearance{
					{
						Appearance: "luminosity",
						Value:      "dark",
					},
				},
				[]appearance{
					{
						Appearance: "luminosity",
						Value:      "light",
					},
				},
			},
		},
		{
			name: "High contrast appearance shoulde should return correct appearances",
			fields: fields{
				build: func() *Appearance {
					a := Appearance{}
					a.Dark()
					a.Light()
					a.HighContrast()
					return &a
				},
			},
			want: appearances{
				[]appearance{},
				[]appearance{
					{
						Appearance: "luminosity",
						Value:      "dark",
					},
				},
				[]appearance{
					{
						Appearance: "luminosity",
						Value:      "light",
					},
				},
				[]appearance{
					{
						Appearance: "contrast",
						Value:      "high",
					},
				},
				[]appearance{
					{
						Appearance: "luminosity",
						Value:      "dark",
					},
					{
						Appearance: "contrast",
						Value:      "high",
					},
				},
				[]appearance{
					{
						Appearance: "luminosity",
						Value:      "light",
					},
					{
						Appearance: "contrast",
						Value:      "high",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := tt.fields.build()
			appearances := a.build()

			assert.Equal(t, tt.want, appearances)
		})
	}
}
