package entitlements

import (
	"strings"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	e := New()

	assert.NotNil(t, e)
	assert.NotNil(t, e.APS)
	assert.NotNil(t, e.DataProtection)
}

func TestEntitlements_SkipValidation(t *testing.T) {
	e := New()
	e.SkipValidation()
	assert.True(t, e.skipValidation)
}

func TestEntitlements_Build(t *testing.T) {
	type fields struct {
		builder func() *Entitlements
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "Empty entitlements should return an error",
			fields: fields{
				builder: func() *Entitlements {
					return New()
				},
			},
			wantErr: true,
		},
		{
			name: "Build with APS entitlements should succeed",
			fields: fields{
				builder: func() *Entitlements {
					e := New()
					e.APS.Production()
					return e
				},
			},
			want: `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0"><dict><key>aps-environment</key><string>production</string></dict></plist>`,
			wantErr: false,
		},
		{
			name: "Custom entitlements should build",
			fields: fields{
				builder: func() *Entitlements {
					e := New()
					e.Set("foo", "bar")
					return e
				},
			},
			want: `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0"><dict><key>foo</key><string>bar</string></dict></plist>`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := tt.fields.builder()
			got, err := e.Build()
			if (err != nil) != tt.wantErr {
				t.Errorf("Entitlements.Build() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Entitlements.Build() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEntitlements_Write(t *testing.T) {
	e := New()
	e.APS.Production()

	buf := strings.Builder{}

	err := e.Write(&buf)
	assert.Nil(t, err)

	expected := `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0"><dict><key>aps-environment</key><string>production</string></dict></plist>`

	assert.Equal(t, expected, buf.String())
}

func TestEntitlements_WriteError(t *testing.T) {
	e := New()

	buf := strings.Builder{}

	err := e.Write(&buf)
	assert.NotNil(t, err)
}

func TestEntitlements_Set(t *testing.T) {
	e := New()

	e.Set("foo", "bar")
	assert.Equal(t, "bar", e.custom["foo"])
}
