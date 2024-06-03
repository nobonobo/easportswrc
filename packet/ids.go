package packet

type IDs struct {
	Versions             IDsVersions            `json:"versions,omitempty"`
	Vehicles             []Vehicles             `json:"vehicles,omitempty"`
	VehicleClasses       []VehicleClasses       `json:"vehicle_classes,omitempty"`
	VehicleManufacturers []VehicleManufacturers `json:"vehicle_manufacturers,omitempty"`
	Locations            []Locations            `json:"locations,omitempty"`
	Routes               []Routes               `json:"routes,omitempty"`
	VehicleTyreState     []VehicleTyreState     `json:"vehicle_tyre_state,omitempty"`
	GameMode             []GameMode             `json:"game_mode,omitempty"`
	StageResultStatus    []StageResultStatus    `json:"stage_result_status,omitempty"`
}

type Data struct {
	Build int `json:"build,omitempty"`
	Major int `json:"major,omitempty"`
	Minor int `json:"minor,omitempty"`
}

type IDsVersions struct {
	Schema int  `json:"schema,omitempty"`
	Data   Data `json:"data,omitempty"`
}

type Vehicles struct {
	ID           int    `json:"id,omitempty"`
	Class        int    `json:"class,omitempty"`
	Manufacturer int    `json:"manufacturer,omitempty"`
	Name         string `json:"name,omitempty"`
	Builder      bool   `json:"builder,omitempty"`
}
type VehicleClasses struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
type VehicleManufacturers struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
type Locations struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
type Routes struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
type VehicleTyreState struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
type GameMode struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
type StageResultStatus struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
