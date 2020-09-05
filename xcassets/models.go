package xcassets

type info struct {
	Author  string `json:"xcode"`
	Version int    `json:"version"`
}

type properties struct {
	Localizable bool `json:"localizable"`
}

type appearance struct {
	Appearance string `json:"appearance"`
	Value      string `json:"value"`
}

type appearances [][]appearance

type colorContainer struct {
	Appearances appearances `json:"appearances"`
	Color       color       `json:"color"`
	Idiom       string      `json:"idiom"`
}

type color struct {
	ColorSpace string          `json:"color-space"`
	Components colorComponents `json:"components"`
}

type colorComponents struct {
	Alpha float64 `json:"alpha"`
	Red   float64 `json:"red,omitempty"`
	Green float64 `json:"green,omitempty"`
	Blue  float64 `json:"blue,omitempty"`
	White float64 `json:"white,omitempty"`
}

type colorSet struct {
	Colors     []colorContainer `json:"colors"`
	Info       info             `json:"info"`
	Properties properties       `json:"properties"`
}
