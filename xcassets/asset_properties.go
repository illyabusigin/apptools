package xcassets

type AssetProperties struct {
	compressionType    *string `json:"compression-type,omitempty"`
	preserveVectorData *bool   `json:"preserves-vector-representation,omitempty"`
	renderAs           *string `json:"template-rendering-intent,omitempty"`

	RenderAs    AssetRendering   `json:"-"`
	Resizing    AssetResizing    `json:"-"`
	Compression AssetCompression `json:"-"`
}

func (p *AssetProperties) build() {
	p.compressionType = p.Compression.Value
	p.preserveVectorData = p.Resizing.preserveVectorData
	p.renderAs = p.RenderAs.Value
}

func (p *AssetProperties) Empty() bool {
	if p.RenderAs.Value != nil {
		return false
	}

	if p.Resizing.preserveVectorData != nil {
		return false
	}

	if p.Compression.Value != nil {
		return false
	}

	return true
}

type AssetResizing struct {
	preserveVectorData *bool
}

func (r *AssetResizing) PreserveVectorData(v bool) {
	r.preserveVectorData = &v
}

type AssetRendering struct {
	Value *string
}

func (r *AssetRendering) Default() {
	r.Value = nil
}

func (r *AssetRendering) Original() {
	v := "original"
	r.Value = &v
}

func (r *AssetRendering) Template() {
	v := "template"
	r.Value = &v
}

type AssetCompression struct {
	Value *string
}

func (c *AssetCompression) Inherited() {
	c.Value = nil
}

func (c *AssetCompression) Automatic() {
	v := "automatic"
	c.Value = &v
}

func (c *AssetCompression) Lossless() {
	v := "lossless"
	c.Value = &v
}

func (c *AssetCompression) Lossy() {
	v := "lossy"
	c.Value = &v
}

func (c *AssetCompression) GPUBestQuality() {
	v := "gpu-optimized-best"
	c.Value = &v
}

func (c *AssetCompression) GPUSmallestSize() {
	v := "gpu-optimized-smallest"
	c.Value = &v
}
