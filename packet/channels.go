package packet

type Versions struct {
	Schema int `json:"schema"`
	Data   int `json:"data"`
}

type ChannelsDef struct {
	Versions Versions   `json:"versions,omitempty"`
	Channels []*Channel `json:"channels,omitempty"`
}

type Channel struct {
	ID          string `json:"id,omitempty"`
	Type        string `json:"type,omitempty"`
	Units       string `json:"units,omitempty"`
	Description string `json:"description,omitempty"`
}
