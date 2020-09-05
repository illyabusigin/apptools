package xcassets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGamut_SRGBAndDisplayP3(t *testing.T) {
	g := Gamut{}
	g.SRGBAndDisplayP3()

	assert.NotEmpty(t, g.values)
	assert.Equal(t, []string{"display-P3", "sRGB"}, g.values)
}

func TestGamut_build(t *testing.T) {
	type fields struct {
		gamut func() *Gamut
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "Empty gamut should be empty",
			fields: fields{
				gamut: func() *Gamut {
					g := &Gamut{}
					return g
				},
			},
			want: []string{},
		},
		{
			name: "Any gamut should be empty",
			fields: fields{
				gamut: func() *Gamut {
					g := &Gamut{}
					g.Any()
					return g
				},
			},
			want: []string{},
		},
		{
			name: "SRGBAndDisplayP3 gamut should not be empty",
			fields: fields{
				gamut: func() *Gamut {
					g := &Gamut{}
					g.SRGBAndDisplayP3()
					return g
				},
			},
			want: []string{"display-P3", "sRGB"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := tt.fields.gamut()
			fields := g.build()
			assert.Equal(t, tt.want, fields)
		})
	}
}
