package packet

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"golang.org/x/sys/windows"
)

type ChannelTable map[string]*Channel

type Packet struct {
	Packet4CC                 [4]byte `json:"packet_4cc"`
	PacketUID                 uint64  `json:"packet_uid"`
	ShiftlightsFraction       float32 `json:"shiftlights_fraction"`
	ShiftlightsRpmStart       float32 `json:"shiftlights_rpm_start"`
	ShiftlightsRpmEnd         float32 `json:"shiftlights_rpm_end"`
	ShiftlightsRpmValid       bool    `json:"shiftlights_rpm_valid"`
	VehicleGearIndex          uint8   `json:"vehicle_gear_index"`
	VehicleGearIndexNeutral   uint8   `json:"vehicle_gear_index_neutral"`
	VehicleGearIndexReverse   uint8   `json:"vehicle_gear_index_reverse"`
	VehicleGearMaximum        uint8   `json:"vehicle_gear_maximum"`
	VehicleSpeed              float32 `json:"vehicle_speed"`
	VehicleTransmissionSpeed  float32 `json:"vehicle_transmission_speed"`
	VehiclePositionX          float32 `json:"VehiclePositionX"`
	VehiclePositionY          float32 `json:"vehicle_position_y"`
	VehiclePositionZ          float32 `json:"vehicle_position_z"`
	VehicleVelocityX          float32 `json:"vehicle_velocity_x"`
	VehicleVelocityY          float32 `json:"vehicle_velocity_y"`
	VehicleVelocityZ          float32 `json:"vehicle_velocity_z"`
	VehicleAccelerationX      float32 `json:"vehicle_acceleration_x"`
	VehicleAccelerationY      float32 `json:"vehicle_acceleration_y"`
	VehicleAccelerationZ      float32 `json:"vehicle_acceleration_z"`
	VehicleLeftDirectionX     float32 `json:"vehicle_left_direction_x"`
	VehicleLeftDirectionY     float32 `json:"vehicle_left_direction_y"`
	VehicleLeftDirectionZ     float32 `json:"vehicle_left_direction_z"`
	VehicleForwardDirectionX  float32 `json:"vehicle_forward_direction_x"`
	VehicleForwardDirectionY  float32 `json:"vehicle_forward_direction_y"`
	VehicleForwardDirectionZ  float32 `json:"vehicle_forward_direction_z"`
	VehicleUpDirectionX       float32 `json:"vehicle_up_direction_x"`
	VehicleUpDirectionY       float32 `json:"vehicle_up_direction_y"`
	VehicleUpDirectionZ       float32 `json:"vehicle_up_direction_z"`
	VehicleHubPositionBl      float32 `json:"vehicle_hub_position_bl"`
	VehicleHubPositionBr      float32 `json:"vehicle_hub_position_br"`
	VehicleHubPositionFl      float32 `json:"vehicle_hub_position_fl"`
	VehicleHubPositionFr      float32 `json:"vehicle_hub_position_fr"`
	VehicleHubVelocityBl      float32 `json:"vehicle_hub_velocity_bl"`
	VehicleHubVelocityBr      float32 `json:"vehicle_hub_velocity_br"`
	VehicleHubVelocityFl      float32 `json:"vehicle_hub_velocity_fl"`
	VehicleHubVelocityFr      float32 `json:"vehicle_hub_velocity_fr"`
	VehicleCpForwardSpeedBl   float32 `json:"vehicle_cp_forward_speed_bl"`
	VehicleCpForwardSpeedBr   float32 `json:"vehicle_cp_forward_speed_br"`
	VehicleCpForwardSpeedFl   float32 `json:"vehicle_cp_forward_speed_fl"`
	VehicleCpForwardSpeedFr   float32 `json:"vehicle_cp_forward_speed_fr"`
	VehicleBrakeTemperatureBl float32 `json:"vehicle_brake_temperature_bl"`
	VehicleBrakeTemperatureBr float32 `json:"vehicle_brake_temperature_br"`
	VehicleBrakeTemperatureFl float32 `json:"vehicle_brake_temperature_fl"`
	VehicleBrakeTemperatureFr float32 `json:"vehicle_brake_temperature_fr"`
	VehicleEngineRpmMax       float32 `json:"vehicle_engine_rpm_max"`
	VehicleEngineRpmIdle      float32 `json:"vehicle_engine_rpm_idle"`
	VehicleEngineRpmCurrent   float32 `json:"vehicle_engine_rpm_current"`
	VehicleThrottle           float32 `json:"vehicle_throttle"`
	VehicleBrake              float32 `json:"vehicle_brake"`
	VehicleClutch             float32 `json:"vehicle_clutch"`
	VehicleSteering           float32 `json:"vehicle_steering"`
	VehicleHandbrake          float32 `json:"vehicle_handbrake"`
	GameTotalTime             float32 `json:"game_total_time"`
	GameDeltaTime             float32 `json:"game_delta_time"`
	GameFrameCount            uint64  `json:"game_frame_count"`
	StageCurrentTime          float32 `json:"stage_current_time"`
	StagePreviousSplitTime    float32 `json:"stage_previous_split_time"`
	StageResultTime           float32 `json:"stage_result_time"`
	StageResultTimePenalty    float32 `json:"stage_result_time_penalty"`
	StageResultStatus         uint8   `json:"stage_result_status"`
	StageCurrentDistance      float64 `json:"stage_current_distance"`
	StageLength               float64 `json:"stage_length"`
	StageProgress             float32 `json:"stage_progress"`
	VehicleTyreStateBl        uint8   `json:"vehicle_tyre_state_bl"`
	VehicleTyreStateBr        uint8   `json:"vehicle_tyre_state_br"`
	VehicleTyreStateFl        uint8   `json:"vehicle_tyre_state_fl"`
	VehicleTyreStateFr        uint8   `json:"vehicle_tyre_state_fr"`
	StageShakedown            bool    `json:"stage_shakedown"`
	GameMode                  uint8   `json:"game_mode"`
	VehicleID                 uint16  `json:"vehicle_id"`
	VehicleClassID            uint16  `json:"vehicle_class_id"`
	VehicleManufacturerID     uint16  `json:"vehicle_manufacturer_id"`
	LocationID                uint16  `json:"location_id"`
	RouteID                   uint16  `json:"route_id"`
	VehicleClusterAbs         bool    `json:"vehicle_cluster_abs"`
}

var (
	WrcRoot              string
	ChannelDicts         = ChannelTable{}
	definitions          = &Definitions{}
	templates            = []string{}
	packetSize           = -1
	gameMode             = map[uint8]string{}
	locations            = map[uint16]string{}
	routes               = map[uint16]string{}
	vehicles             = map[uint16]string{}
	vehicleClasses       = map[uint16]string{}
	vehicleManufacturers = map[uint16]string{}
	vehicleTyreStates    = map[uint8]string{}
	stageResultStatus    = map[uint8]string{}
	endian               = binary.LittleEndian
)

func init() {
	doc, err := windows.KnownFolderPath(windows.FOLDERID_Documents, 0)
	if err != nil {
		log.Fatal(err)
	}
	WrcRoot = filepath.Join(doc, "My Games", "WRC")
	v, ok := os.LookupEnv("EASPORTSWRC_DOC_ROOT")
	if ok {
		WrcRoot = v
	}
	if _, err := os.Stat(WrcRoot); err != nil {
		log.Fatal(err)
	}
	ib, err := ReadFileUTF16(filepath.Join(WrcRoot, "telemetry", "readme", "ids.json"))
	if err != nil {
		log.Fatal(err)
	}
	idjson := &IDs{}
	if err := json.Unmarshal(ib, &idjson); err != nil {
		log.Fatal(err)
	}
	for _, v := range idjson.GameMode {
		gameMode[uint8(v.ID)] = v.Name
	}
	for _, v := range idjson.Locations {
		locations[uint16(v.ID)] = v.Name
	}
	for _, v := range idjson.Routes {
		routes[uint16(v.ID)] = v.Name
	}
	for _, v := range idjson.Vehicles {
		vehicles[uint16(v.ID)] = v.Name
	}
	for _, v := range idjson.VehicleClasses {
		vehicleClasses[uint16(v.ID)] = v.Name
	}
	for _, v := range idjson.VehicleManufacturers {
		vehicleManufacturers[uint16(v.ID)] = v.Name
	}
	for _, v := range idjson.VehicleTyreState {
		vehicleTyreStates[uint8(v.ID)] = v.Name
	}
	for _, v := range idjson.StageResultStatus {
		stageResultStatus[uint8(v.ID)] = v.Name
	}
	cb, err := os.ReadFile(filepath.Join(WrcRoot, "telemetry", "readme", "channels.json"))
	if err != nil {
		log.Fatal(err)
	}
	var chdefs *ChannelsDef
	if err := json.Unmarshal(cb, &chdefs); err != nil {
		log.Fatal(err)
	}
	for _, ch := range chdefs.Channels {
		ChannelDicts[ch.ID] = ch
	}
	conf, err := os.ReadFile(filepath.Join(WrcRoot, "telemetry", "config.json"))
	if err != nil {
		log.Fatal(err)
	}
	var config *Config
	if err := json.Unmarshal(conf, &config); err != nil {
		log.Fatal(err)
	}
	name := config.UDP.Packets[0].Structure
	fpath := filepath.Join(WrcRoot, "telemetry", "udp", name+".json")
	switch name {
	case "wrc", "wrc_experimental":
		fpath = filepath.Join(WrcRoot, "telemetry", "readme", "udp", name+".json")
	}
	pb, err := os.ReadFile(fpath)
	if err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(pb, &definitions); err != nil {
		log.Fatal(err)
	}
	templates = definitions.Packets[0].Channels
	sz := 0
	for _, key := range templates {
		channel, ok := ChannelDicts[key]
		if !ok {
			log.Fatal(fmt.Errorf("channel %s not found", key))
		}
		switch channel.Type {
		default:
			log.Fatal(fmt.Errorf("type %s not found", channel.Type))
		case "boolean":
			sz += 1
		case "float32":
			sz += 4
		case "float64":
			sz += 8
		case "fourcc":
			sz += 4
		case "uint8":
			sz += 1
		case "uint16":
			sz += 2
		case "uint64":
			sz += 8
		}
	}
	packetSize = sz
}

func New() *Packet {
	return &Packet{}
}

func (p *Packet) Length() int {
	return packetSize
}

func (p *Packet) String() string {
	return fmt.Sprintf("%v", *p)
}

func (p *Packet) Fields() []string {
	return templates
}

func (p *Packet) MarshalBinary() ([]byte, error) {
	writer := bytes.NewBuffer(nil)
	for _, v := range templates {
		switch v {
		default:
			return nil, fmt.Errorf("unknown field %s", v)
		case "packet_4cc":
			binary.Write(writer, endian, p.Packet4CC)
		case "packet_uid":
			binary.Write(writer, endian, p.PacketUID)
		case "shiftlights_fraction":
			binary.Write(writer, endian, p.ShiftlightsFraction)
		case "shiftlights_rpm_start":
			binary.Write(writer, endian, p.ShiftlightsRpmStart)
		case "shiftlights_rpm_end":
			binary.Write(writer, endian, p.ShiftlightsRpmEnd)
		case "shiftlights_rpm_valid":
			if p.ShiftlightsRpmValid {
				writer.Write([]byte{1})
			} else {
				writer.Write([]byte{0})
			}
		case "vehicle_gear_index":
			binary.Write(writer, endian, p.VehicleGearIndex)
		case "vehicle_gear_index_neutral":
			binary.Write(writer, endian, p.VehicleGearIndexNeutral)
		case "vehicle_gear_index_reverse":
			binary.Write(writer, endian, p.VehicleGearIndexReverse)
		case "vehicle_gear_maximum":
			binary.Write(writer, endian, p.VehicleGearMaximum)
		case "vehicle_speed":
			binary.Write(writer, endian, p.VehicleSpeed)
		case "vehicle_transmission_speed":
			binary.Write(writer, endian, p.VehicleTransmissionSpeed)
		case "vehicle_position_x":
			binary.Write(writer, endian, p.VehiclePositionX)
		case "vehicle_position_y":
			binary.Write(writer, endian, p.VehiclePositionY)
		case "vehicle_position_z":
			binary.Write(writer, endian, p.VehiclePositionZ)
		case "vehicle_velocity_x":
			binary.Write(writer, endian, p.VehicleVelocityX)
		case "vehicle_velocity_y":
			binary.Write(writer, endian, p.VehicleVelocityY)
		case "vehicle_velocity_z":
			binary.Write(writer, endian, p.VehicleVelocityZ)
		case "vehicle_acceleration_x":
			binary.Write(writer, endian, p.VehicleAccelerationX)
		case "vehicle_acceleration_y":
			binary.Write(writer, endian, p.VehicleAccelerationY)
		case "vehicle_acceleration_z":
			binary.Write(writer, endian, p.VehicleAccelerationZ)
		case "vehicle_left_direction_x":
			binary.Write(writer, endian, p.VehicleLeftDirectionX)
		case "vehicle_left_direction_y":
			binary.Write(writer, endian, p.VehicleLeftDirectionY)
		case "vehicle_left_direction_z":
			binary.Write(writer, endian, p.VehicleLeftDirectionZ)
		case "vehicle_forward_direction_x":
			binary.Write(writer, endian, p.VehicleForwardDirectionX)
		case "vehicle_forward_direction_y":
			binary.Write(writer, endian, p.VehicleForwardDirectionY)
		case "vehicle_forward_direction_z":
			binary.Write(writer, endian, p.VehicleForwardDirectionZ)
		case "vehicle_up_direction_x":
			binary.Write(writer, endian, p.VehicleUpDirectionX)
		case "vehicle_up_direction_y":
			binary.Write(writer, endian, p.VehicleUpDirectionY)
		case "vehicle_up_direction_z":
			binary.Write(writer, endian, p.VehicleUpDirectionZ)
		case "vehicle_hub_position_bl":
			binary.Write(writer, endian, p.VehicleHubPositionBl)
		case "vehicle_hub_position_br":
			binary.Write(writer, endian, p.VehicleHubPositionBr)
		case "vehicle_hub_position_fl":
			binary.Write(writer, endian, p.VehicleHubPositionFl)
		case "vehicle_hub_position_fr":
			binary.Write(writer, endian, p.VehicleHubPositionFr)
		case "vehicle_hub_velocity_bl":
			binary.Write(writer, endian, p.VehicleHubVelocityBl)
		case "vehicle_hub_velocity_br":
			binary.Write(writer, endian, p.VehicleHubVelocityBr)
		case "vehicle_hub_velocity_fl":
			binary.Write(writer, endian, p.VehicleHubVelocityFl)
		case "vehicle_hub_velocity_fr":
			binary.Write(writer, endian, p.VehicleHubVelocityFr)
		case "vehicle_cp_forward_speed_bl":
			binary.Write(writer, endian, p.VehicleCpForwardSpeedBl)
		case "vehicle_cp_forward_speed_br":
			binary.Write(writer, endian, p.VehicleCpForwardSpeedBr)
		case "vehicle_cp_forward_speed_fl":
			binary.Write(writer, endian, p.VehicleCpForwardSpeedFl)
		case "vehicle_cp_forward_speed_fr":
			binary.Write(writer, endian, p.VehicleCpForwardSpeedFr)
		case "vehicle_brake_temperature_bl":
			binary.Write(writer, endian, p.VehicleBrakeTemperatureBl)
		case "vehicle_brake_temperature_br":
			binary.Write(writer, endian, p.VehicleBrakeTemperatureBr)
		case "vehicle_brake_temperature_fl":
			binary.Write(writer, endian, p.VehicleBrakeTemperatureFl)
		case "vehicle_brake_temperature_fr":
			binary.Write(writer, endian, p.VehicleBrakeTemperatureFr)
		case "vehicle_engine_rpm_max":
			binary.Write(writer, endian, p.VehicleEngineRpmMax)
		case "vehicle_engine_rpm_idle":
			binary.Write(writer, endian, p.VehicleEngineRpmIdle)
		case "vehicle_engine_rpm_current":
			binary.Write(writer, endian, p.VehicleEngineRpmCurrent)
		case "vehicle_throttle":
			binary.Write(writer, endian, p.VehicleThrottle)
		case "vehicle_brake":
			binary.Write(writer, endian, p.VehicleBrake)
		case "vehicle_clutch":
			binary.Write(writer, endian, p.VehicleClutch)
		case "vehicle_steering":
			binary.Write(writer, endian, p.VehicleSteering)
		case "vehicle_handbrake":
			binary.Write(writer, endian, p.VehicleHandbrake)
		case "game_total_time":
			binary.Write(writer, endian, p.GameTotalTime)
		case "game_delta_time":
			binary.Write(writer, endian, p.GameDeltaTime)
		case "game_frame_count":
			binary.Write(writer, endian, p.GameFrameCount)
		case "stage_current_time":
			binary.Write(writer, endian, p.StageCurrentTime)
		case "stage_previous_split_time":
			binary.Write(writer, endian, p.StagePreviousSplitTime)
		case "stage_result_time":
			binary.Write(writer, endian, p.StageResultTime)
		case "stage_result_time_penalty":
			binary.Write(writer, endian, p.StageResultTimePenalty)
		case "stage_result_status":
			binary.Write(writer, endian, p.StageResultStatus)
		case "stage_current_distance":
			binary.Write(writer, endian, p.StageCurrentDistance)
		case "stage_length":
			binary.Write(writer, endian, p.StageLength)
		case "stage_progress":
			binary.Write(writer, endian, p.StageProgress)
		case "vehicle_tyre_state_bl":
			binary.Write(writer, endian, p.VehicleTyreStateBl)
		case "vehicle_tyre_state_br":
			binary.Write(writer, endian, p.VehicleTyreStateBr)
		case "vehicle_tyre_state_fl":
			binary.Write(writer, endian, p.VehicleTyreStateFl)
		case "vehicle_tyre_state_fr":
			binary.Write(writer, endian, p.VehicleTyreStateFr)
		case "stage_shakedown":
			if p.StageShakedown {
				writer.Write([]byte{1})
			} else {
				writer.Write([]byte{0})
			}
		case "game_mode":
			binary.Write(writer, endian, p.GameMode)
		case "vehicle_id":
			binary.Write(writer, endian, p.VehicleID)
		case "vehicle_class_id":
			binary.Write(writer, endian, p.VehicleClassID)
		case "vehicle_manufacturer_id":
			binary.Write(writer, endian, p.VehicleManufacturerID)
		case "location_id":
			binary.Write(writer, endian, p.LocationID)
		case "route_id":
			binary.Write(writer, endian, p.RouteID)
		case "vehicle_cluster_abs":
			if p.VehicleClusterAbs {
				writer.Write([]byte{1})
			} else {
				writer.Write([]byte{0})
			}
		}
	}
	if writer.Len() != packetSize {
		return nil, fmt.Errorf("invalid packet size %d expected: %d", writer.Len(), packetSize)
	}
	return writer.Bytes(), nil
}

func (p *Packet) UnmarshalBinary(b []byte) error {
	if len(b) != packetSize {
		return fmt.Errorf("invalid packet size %d expected: %d", len(b), packetSize)
	}
	reader := bytes.NewReader(b)
	buf := []byte{0}
	for _, v := range templates {
		var err error
		var n int
		switch v {
		default:
			return fmt.Errorf("unknown field %s", v)
		case "packet_4cc":
			err = binary.Read(reader, binary.LittleEndian, &p.Packet4CC)
		case "packet_uid":
			err = binary.Read(reader, binary.LittleEndian, &p.PacketUID)
		case "shiftlights_fraction":
			err = binary.Read(reader, binary.LittleEndian, &p.ShiftlightsFraction)
		case "shiftlights_rpm_start":
			err = binary.Read(reader, binary.LittleEndian, &p.ShiftlightsRpmStart)
		case "shiftlights_rpm_end":
			err = binary.Read(reader, binary.LittleEndian, &p.ShiftlightsRpmEnd)
		case "shiftlights_rpm_valid":
			n, err = reader.Read(buf)
			if n > 0 {
				p.ShiftlightsRpmValid = buf[0] != 0
			}
		case "vehicle_gear_index":
			err = binary.Read(reader, binary.LittleEndian, &p.VehicleGearIndex)
		case "vehicle_gear_index_neutral":
			err = binary.Read(reader, binary.LittleEndian, &p.VehicleGearIndexNeutral)
		case "vehicle_gear_index_reverse":
			err = binary.Read(reader, binary.LittleEndian, &p.VehicleGearIndexReverse)
		case "vehicle_gear_maximum":
			err = binary.Read(reader, binary.LittleEndian, &p.VehicleGearMaximum)
		case "vehicle_speed":
			err = binary.Read(reader, binary.LittleEndian, &p.VehicleSpeed)
		case "vehicle_transmission_speed":
			err = binary.Read(reader, binary.LittleEndian, &p.VehicleTransmissionSpeed)
		case "vehicle_position_x":
			err = binary.Read(reader, binary.LittleEndian, &p.VehiclePositionX)
		case "vehicle_position_y":
			err = binary.Read(reader, binary.LittleEndian, &p.VehiclePositionY)
		case "vehicle_position_z":
			err = binary.Read(reader, binary.LittleEndian, &p.VehiclePositionZ)
		case "vehicle_velocity_x":
			err = binary.Read(reader, binary.LittleEndian, &p.VehicleVelocityX)
		case "vehicle_velocity_y":
			err = binary.Read(reader, binary.LittleEndian, &p.VehicleVelocityY)
		case "vehicle_velocity_z":
			err = binary.Read(reader, binary.LittleEndian, &p.VehicleVelocityZ)
		case "vehicle_acceleration_x":
			err = binary.Read(reader, binary.LittleEndian, &p.VehicleAccelerationX)
		case "vehicle_acceleration_y":
			err = binary.Read(reader, binary.LittleEndian, &p.VehicleAccelerationY)
		case "vehicle_acceleration_z":
			err = binary.Read(reader, binary.LittleEndian, &p.VehicleAccelerationZ)
		case "vehicle_left_direction_x":
			err = binary.Read(reader, binary.LittleEndian, &p.VehicleLeftDirectionX)
		case "vehicle_left_direction_y":
			err = binary.Read(reader, binary.LittleEndian, &p.VehicleLeftDirectionY)
		case "vehicle_left_direction_z":
			err = binary.Read(reader, binary.LittleEndian, &p.VehicleLeftDirectionZ)
		case "vehicle_forward_direction_x":
			err = binary.Read(reader, binary.LittleEndian, &p.VehicleForwardDirectionX)
		case "vehicle_forward_direction_y":
			err = binary.Read(reader, binary.LittleEndian, &p.VehicleForwardDirectionY)
		case "vehicle_forward_direction_z":
			err = binary.Read(reader, binary.LittleEndian, &p.VehicleForwardDirectionZ)
		case "vehicle_up_direction_x":
			err = binary.Read(reader, binary.LittleEndian, &p.VehicleUpDirectionX)
		case "vehicle_up_direction_y":
			err = binary.Read(reader, binary.LittleEndian, &p.VehicleUpDirectionY)
		case "vehicle_up_direction_z":
			err = binary.Read(reader, binary.LittleEndian, &p.VehicleUpDirectionZ)
		case "vehicle_hub_position_bl":
			err = binary.Read(reader, binary.LittleEndian, &p.VehicleHubPositionBl)
		case "vehicle_hub_position_br":
			err = binary.Read(reader, binary.LittleEndian, &p.VehicleHubPositionBr)
		case "vehicle_hub_position_fl":
			err = binary.Read(reader, binary.LittleEndian, &p.VehicleHubPositionFl)
		case "vehicle_hub_position_fr":
			err = binary.Read(reader, binary.LittleEndian, &p.VehicleHubPositionFr)
		case "vehicle_hub_velocity_bl":
			err = binary.Read(reader, binary.LittleEndian, &p.VehicleHubVelocityBl)
		case "vehicle_hub_velocity_br":
			err = binary.Read(reader, binary.LittleEndian, &p.VehicleHubVelocityBr)
		case "vehicle_hub_velocity_fl":
			err = binary.Read(reader, binary.LittleEndian, &p.VehicleHubVelocityFl)
		case "vehicle_hub_velocity_fr":
			err = binary.Read(reader, binary.LittleEndian, &p.VehicleHubVelocityFr)
		case "vehicle_cp_forward_speed_bl":
			err = binary.Read(reader, binary.LittleEndian, &p.VehicleCpForwardSpeedBl)
		case "vehicle_cp_forward_speed_br":
			err = binary.Read(reader, binary.LittleEndian, &p.VehicleCpForwardSpeedBr)
		case "vehicle_cp_forward_speed_fl":
			err = binary.Read(reader, binary.LittleEndian, &p.VehicleCpForwardSpeedFl)
		case "vehicle_cp_forward_speed_fr":
			err = binary.Read(reader, binary.LittleEndian, &p.VehicleCpForwardSpeedFr)
		case "vehicle_brake_temperature_bl":
			err = binary.Read(reader, binary.LittleEndian, &p.VehicleBrakeTemperatureBl)
		case "vehicle_brake_temperature_br":
			err = binary.Read(reader, binary.LittleEndian, &p.VehicleBrakeTemperatureBr)
		case "vehicle_brake_temperature_fl":
			err = binary.Read(reader, binary.LittleEndian, &p.VehicleBrakeTemperatureFl)
		case "vehicle_brake_temperature_fr":
			err = binary.Read(reader, binary.LittleEndian, &p.VehicleBrakeTemperatureFr)
		case "vehicle_engine_rpm_max":
			err = binary.Read(reader, binary.LittleEndian, &p.VehicleEngineRpmMax)
		case "vehicle_engine_rpm_idle":
			err = binary.Read(reader, binary.LittleEndian, &p.VehicleEngineRpmIdle)
		case "vehicle_engine_rpm_current":
			err = binary.Read(reader, binary.LittleEndian, &p.VehicleEngineRpmCurrent)
		case "vehicle_throttle":
			err = binary.Read(reader, binary.LittleEndian, &p.VehicleThrottle)
		case "vehicle_brake":
			err = binary.Read(reader, binary.LittleEndian, &p.VehicleBrake)
		case "vehicle_clutch":
			err = binary.Read(reader, binary.LittleEndian, &p.VehicleClutch)
		case "vehicle_steering":
			err = binary.Read(reader, binary.LittleEndian, &p.VehicleSteering)
		case "vehicle_handbrake":
			err = binary.Read(reader, binary.LittleEndian, &p.VehicleHandbrake)
		case "game_total_time":
			err = binary.Read(reader, binary.LittleEndian, &p.GameTotalTime)
		case "game_delta_time":
			err = binary.Read(reader, binary.LittleEndian, &p.GameDeltaTime)
		case "game_frame_count":
			err = binary.Read(reader, binary.LittleEndian, &p.GameFrameCount)
		case "stage_current_time":
			err = binary.Read(reader, binary.LittleEndian, &p.StageCurrentTime)
		case "stage_previous_split_time":
			err = binary.Read(reader, binary.LittleEndian, &p.StagePreviousSplitTime)
		case "stage_result_time":
			err = binary.Read(reader, binary.LittleEndian, &p.StageResultTime)
		case "stage_result_time_penalty":
			err = binary.Read(reader, binary.LittleEndian, &p.StageResultTimePenalty)
		case "stage_result_status":
			err = binary.Read(reader, binary.LittleEndian, &p.StageResultStatus)
		case "stage_current_distance":
			err = binary.Read(reader, binary.LittleEndian, &p.StageCurrentDistance)
		case "stage_length":
			err = binary.Read(reader, binary.LittleEndian, &p.StageLength)
		case "stage_progress":
			err = binary.Read(reader, binary.LittleEndian, &p.StageProgress)
		case "vehicle_tyre_state_bl":
			err = binary.Read(reader, binary.LittleEndian, &p.VehicleTyreStateBl)
		case "vehicle_tyre_state_br":
			err = binary.Read(reader, binary.LittleEndian, &p.VehicleTyreStateBr)
		case "vehicle_tyre_state_fl":
			err = binary.Read(reader, binary.LittleEndian, &p.VehicleTyreStateFl)
		case "vehicle_tyre_state_fr":
			err = binary.Read(reader, binary.LittleEndian, &p.VehicleTyreStateFr)
		case "stage_shakedown":
			n, err = reader.Read(buf)
			if n > 0 {
				p.StageShakedown = buf[0] != 0
			}
		case "game_mode":
			err = binary.Read(reader, binary.LittleEndian, &p.GameMode)
		case "vehicle_id":
			err = binary.Read(reader, binary.LittleEndian, &p.VehicleID)
		case "vehicle_class_id":
			err = binary.Read(reader, binary.LittleEndian, &p.VehicleClassID)
		case "vehicle_manufacturer_id":
			err = binary.Read(reader, binary.LittleEndian, &p.VehicleManufacturerID)
		case "location_id":
			err = binary.Read(reader, binary.LittleEndian, &p.LocationID)
		case "route_id":
			err = binary.Read(reader, binary.LittleEndian, &p.RouteID)
		case "vehicle_cluster_abs":
			n, err = reader.Read(buf)
			if n > 0 {
				p.VehicleClusterAbs = buf[0] != 0
			}
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *Packet) GameModeString() string {
	s, ok := gameMode[p.GameMode]
	if !ok {
		return "unknown"
	}
	return s
}

func (p *Packet) Location() string {
	s, ok := locations[p.LocationID]
	if !ok {
		return "unknown"
	}
	return s
}

func (p *Packet) Route() string {
	s, ok := routes[p.RouteID]
	if !ok {
		return "unknown"
	}
	return s
}

func (p *Packet) Vehicle() string {
	s, ok := vehicles[p.VehicleID]
	if !ok {
		return "unknown"
	}
	return s
}

func (p *Packet) VehicleClass() string {
	s, ok := vehicleClasses[p.VehicleClassID]
	if !ok {
		return "unknown"
	}
	return s
}

func (p *Packet) VehicleManufacturer() string {
	s, ok := vehicleManufacturers[p.VehicleManufacturerID]
	if !ok {
		return "unknown"
	}
	return s
}

type Position string

const (
	ForwardLeft   Position = "_fl"
	ForwardRight  Position = "_fr"
	BackwordLeft  Position = "_bl"
	BackwordRight Position = "_br"
)

func (p *Packet) VehicleTyreState(pos Position) string {
	v := uint8(0)
	switch pos {
	default:
		return "unknown"
	case ForwardLeft:
		v = p.VehicleTyreStateFl
	case ForwardRight:
		v = p.VehicleTyreStateFr
	case BackwordLeft:
		v = p.VehicleTyreStateBl
	case BackwordRight:
		v = p.VehicleTyreStateBr
	}
	s, ok := vehicleTyreStates[v]
	if !ok {
		return "unknown"
	}
	return s
}

func (p *Packet) StageResultStatusString() string {
	s, ok := stageResultStatus[p.StageResultStatus]
	if !ok {
		return "unknown"
	}
	return s
}
