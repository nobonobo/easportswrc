package packet

type Config struct {
	Schema int  `json:"schema"`
	UDP    UDP  `json:"udp"`
	Lcd    Lcd  `json:"lcd"`
	DBox   DBox `json:"dBox"`
}

type Packets struct {
	Structure   string `json:"structure"`
	Packet      string `json:"packet"`
	IP          string `json:"ip"`
	Port        int    `json:"port"`
	FrequencyHz int    `json:"frequencyHz"`
	BEnabled    bool   `json:"bEnabled"`
}

type UDP struct {
	Packets []Packets `json:"packets"`
}

type Lcd struct {
	BDisplayGears bool `json:"bDisplayGears"`
}

type DBox struct {
	BEnabled bool `json:"bEnabled"`
}
