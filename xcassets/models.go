package xcassets

type info struct {
	Author  string `json:"author"`
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
	Appearances  []appearance `json:"appearances,omitempty"`
	Color        color        `json:"color"`
	Idiom        string       `json:"idiom"`
	DisplayGamut string       `json:"display-gamut,omitempty"`
}

type color struct {
	ColorSpace string          `json:"color-space"`
	Components colorComponents `json:"components"`
}

type colorComponents struct {
	Alpha string `json:"alpha"`
	Red   string `json:"red,omitempty"`
	Green string `json:"green,omitempty"`
	Blue  string `json:"blue,omitempty"`
	White string `json:"white,omitempty"`
}

type colorSet struct {
	Colors     []colorContainer `json:"colors"`
	Info       info             `json:"info"`
	Properties properties       `json:"properties"`
}
