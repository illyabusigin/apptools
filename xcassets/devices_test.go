package xcassets

import (
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestDevices_validate(t *testing.T) {
	d := Devices{}
	assert.NotNil(t, d.Validate(), "Specifying no devices should return an error")

	d.Universal()
	assert.Nil(t, d.Validate(), "Validation should pass with one device specified")
}

func TestDevices_build(t *testing.T) {
	d := Devices{}
	d.Universal()
	d.IPhone()
	d.IPad()
	d.Catalyst()
	d.CarPlay()
	d.AppleWatch()
	d.AppleTV()
	d.Mac()

	idioms := d.build()
	expected := []string{"universal", "iphone", "ipad", "car", "watch", "tv", "mac"}

	assert.Equal(t, expected, idioms, "Idioms should equal specified idioms")
}

func TestDevices_subtypes(t *testing.T) {
	d := Devices{}
	d.Catalyst()

	assert.True(t, d.iPad)
	assert.True(t, d.catalyst)

	expected := []string{"mac-catalyst"}
	subtypes := d.subtypes()

	assert.Equal(t, expected, subtypes)
}

func TestDevices_intersects(t *testing.T) {
	type fields struct {
		first  func() *Devices
		second func() *Devices
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "No overlap should be empty",
			fields: fields{
				first: func() *Devices {
					d := &Devices{}
					d.Universal()
					return d
				},
				second: func() *Devices {
					d := &Devices{}
					d.IPad()
					return d
				},
			},
			want: []string{},
		},
		{
			name: "Overlap should return intersection",
			fields: fields{
				first: func() *Devices {
					d := &Devices{}
					d.IPad()
					return d
				},
				second: func() *Devices {
					d := &Devices{}
					d.IPad()
					return d
				},
			},
			want: []string{"ipad"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g1 := tt.fields.first()
			g2 := tt.fields.second()

			i := g1.intersects(g2)
			assert.Equal(t, tt.want, i)
		})
	}
}
