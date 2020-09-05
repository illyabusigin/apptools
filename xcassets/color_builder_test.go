package xcassets

import (
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
		// {
		// 	name: "Testing",
		// 	fields: fields{
		// 		build: func() *ColorBuilder {
		// 			b := Color("TestColor", func(b *ColorBuilder) {

		// 			})

		// 			return b
		// 		},
		// 	},
		// 	want:    `{"colors":[],"info":{"xcode":"xcode","version":1},"properties":{"localizable":true}}`,
		// 	wantErr: false,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			colorBuilder := tt.fields.build()
			out, err := colorBuilder.Build()

			if tt.wantErr {
				assert.NotNil(t, err)
			}

			assert.Equal(t, tt.want, out)
		})
	}
}
