package packet

type Definitions struct {
	Versions Versions     `json:"versions,omitempty"`
	ID       string       `json:"id,omitempty"`
	Header   Header       `json:"header,omitempty"`
	Packets  []Definition `json:"packets,omitempty"`
}

type Header struct {
	Channels []any `json:"channels,omitempty"`
}

type Definition struct {
	ID       string   `json:"id,omitempty"`
	Channels []string `json:"channels,omitempty"`
}
