package packet

type PacketsDef struct {
	Versions Versions    `json:"versions,omitempty"`
	ID       string      `json:"id,omitempty"`
	Header   Header      `json:"header,omitempty"`
	Packets  []PacketDef `json:"packets,omitempty"`
}

type Header struct {
	Channels []any `json:"channels,omitempty"`
}

type PacketDef struct {
	ID       string   `json:"id,omitempty"`
	Channels []string `json:"channels,omitempty"`
}
