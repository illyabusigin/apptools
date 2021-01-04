package xcassets

import (
	"encoding/json"
)

// AssetProperties contains all properties for the image set.
type AssetProperties struct {
	compressionType    *string
	preserveVectorData *bool
	renderAs           *string

	RenderAs    AssetRendering   `json:"-"`
	Vector      AssetVectorData  `json:"-"`
	Compression AssetCompression `json:"-"`
}

// MarshalJSON marshals the asset properties to JSON.
func (p *AssetProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		CompressionType    *string `json:"compression-type,omitempty"`
		PreserveVectorData *bool   `json:"preserves-vector-representation,omitempty"`
		RenderAs           *string `json:"template-rendering-intent,omitempty"`
	}{
		p.compressionType,
		p.preserveVectorData,
		p.renderAs,
	})
}

func (p *AssetProperties) build() {
	p.compressionType = p.Compression.Value
	p.preserveVectorData = p.Vector.preserveVectorData
	p.renderAs = p.RenderAs.Value
}

// Empty returns a boolean value indicating whether or not the properties are empty.
func (p *AssetProperties) Empty() bool {
	if p.RenderAs.Value != nil {
		return false
	}

	if p.Vector.preserveVectorData != nil {
		return false
	}

	if p.Compression.Value != nil {
		return false
	}

	return true
}

// AssetVectorData contains vector data information
type AssetVectorData struct {
	preserveVectorData *bool
}

// PreserveVectorData set to `true` will preserve the vector information for
// a PDF file.
func (r *AssetVectorData) PreserveVectorData(v bool) {
	r.preserveVectorData = &v
}

// AssetRendering contains data specifiying if the image is a template for use
// with visual effects such as replacing colors
type AssetRendering struct {
	Value *string
}

// Default behavior. If the name of the image ends in "Template", use the image
// as a template, otherwise render it as the original image
func (r *AssetRendering) Default() {
	r.Value = nil
}

// Original renders the image as the original image.
func (r *AssetRendering) Original() {
	v := "original"
	r.Value = &v
}

// Template  the image as a template for visual effects such as replacing
// colors.
func (r *AssetRendering) Template() {
	v := "template"
	r.Value = &v
}

// AssetCompression specifies the compression used for the asset.
type AssetCompression struct {
	Value *string
}

// Inherited wil inherient the cmompression type from the parent. If no parent
// is set the type is set to losses compression.
func (c *AssetCompression) Inherited() {
	c.Value = nil
}

// Automatic specifies the image uses an automatic lossy compression.
func (c *AssetCompression) Automatic() {
	v := "automatic"
	c.Value = &v
}

// Lossless uses lossless compression. This is the default if no compression is
// specified.
func (c *AssetCompression) Lossless() {
	v := "lossless"
	c.Value = &v
}

// Lossy specifies the image uses basic lossy commpression.
func (c *AssetCompression) Lossy() {
	v := "lossy"
	c.Value = &v
}

// GPUBestQuality specifies the image uses a lossy GPU compression format
// optimized for quality.
func (c *AssetCompression) GPUBestQuality() {
	v := "gpu-optimized-best"
	c.Value = &v
}

// GPUSmallestSize specifies the image uses a lossy GPU compression format
// optimized for memory size
func (c *AssetCompression) GPUSmallestSize() {
	v := "gpu-optimized-smallest"
	c.Value = &v
}
