package entitlements

import (
	"bytes"
	"fmt"
	"io"

	"howett.net/plist"
)

// Entitlements is a builder for assigning key-value pairs that grant
// executable permission to use a service or technology.
// See https://developer.apple.com/documentation/bundleresources/entitlements/ for more information.
type Entitlements struct {
	skipValidation bool

	APS            *APS
	DataProtection *DataProtection

	data   map[string]interface{}
	custom map[string]interface{}
}

// SkipValidation will skip all validation when building entitlements
func (e *Entitlements) SkipValidation() {
	e.skipValidation = true
}

// Set will set an arbitrary key-value pair in your entitlements. Keys set in this
// manner will override any keys set by any of the builder functions.
func (e *Entitlements) Set(key string, value interface{}) {
	e.custom[key] = value
}

// Build will build the Entitlements property list
func (e *Entitlements) Build() (string, error) {
	buf := bytes.Buffer{}

	e.data = map[string]interface{}{}

	e.APS.Apply(e)
	e.DataProtection.Apply(e)

	for key, val := range e.custom {
		e.data[key] = val
	}

	if len(e.data) == 0 {
		return "", fmt.Errorf("No entitlements found")
	}

	encoder := plist.NewEncoder(&buf)
	if err := encoder.Encode(e.data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

// Write the entitlements to the specified io.Writer.
func (e *Entitlements) Write(w io.Writer) error {
	data, err := e.Build()
	if err != nil {
		return err
	}

	_, err = w.Write([]byte(data))

	return err
}

// New returns a new `Entitlements` builder.
func New() *Entitlements {
	return &Entitlements{
		APS:            &APS{},
		DataProtection: &DataProtection{},
		custom:         map[string]interface{}{},
	}
}
