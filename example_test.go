package main

import (
	"fmt"

	"github.com/nobonobo/easportswrc/packet"
)

func ExampleNew() {
	pkt := packet.New()
	fmt.Println(pkt)
	// Output:
	// packet_4cc:____, packet_uid:0, shiftlights_fraction:0.000000, shiftlights_rpm_start:0.000000, shiftlights_rpm_end:0.000000, shiftlights_rpm_valid:false, vehicle_gear_index:0, vehicle_gear_index_neutral:0, vehicle_gear_index_reverse:0, vehicle_gear_maximum:0, vehicle_speed:0.000000, vehicle_transmission_speed:0.000000, vehicle_position_x:0.000000, vehicle_position_y:0.000000, vehicle_position_z:0.000000, vehicle_velocity_x:0.000000, vehicle_velocity_y:0.000000, vehicle_velocity_z:0.000000, vehicle_acceleration_x:0.000000, vehicle_acceleration_y:0.000000, vehicle_acceleration_z:0.000000, vehicle_left_direction_x:0.000000, vehicle_left_direction_y:0.000000, vehicle_left_direction_z:0.000000, vehicle_forward_direction_x:0.000000, vehicle_forward_direction_y:0.000000, vehicle_forward_direction_z:0.000000, vehicle_up_direction_x:0.000000, vehicle_up_direction_y:0.000000, vehicle_up_direction_z:0.000000, vehicle_hub_position_bl:0.000000, vehicle_hub_position_br:0.000000, vehicle_hub_position_fl:0.000000, vehicle_hub_position_fr:0.000000, vehicle_hub_velocity_bl:0.000000, vehicle_hub_velocity_br:0.000000, vehicle_hub_velocity_fl:0.000000, vehicle_hub_velocity_fr:0.000000, vehicle_cp_forward_speed_bl:0.000000, vehicle_cp_forward_speed_br:0.000000, vehicle_cp_forward_speed_fl:0.000000, vehicle_cp_forward_speed_fr:0.000000, vehicle_brake_temperature_bl:0.000000, vehicle_brake_temperature_br:0.000000, vehicle_brake_temperature_fl:0.000000, vehicle_brake_temperature_fr:0.000000, vehicle_engine_rpm_max:0.000000, vehicle_engine_rpm_idle:0.000000, vehicle_engine_rpm_current:0.000000, vehicle_throttle:0.000000, vehicle_brake:0.000000, vehicle_clutch:0.000000, vehicle_steering:0.000000, vehicle_handbrake:0.000000, game_total_time:0.000000, game_delta_time:0.000000, game_frame_count:0, stage_current_time:0.000000, stage_previous_split_time:0.000000, stage_result_time:0.000000, stage_result_time_penalty:0.000000, stage_result_status:0, stage_current_distance:0.000000, stage_length:0.000000, stage_progress:0.000000, vehicle_tyre_state_bl:0, vehicle_tyre_state_br:0, vehicle_tyre_state_fl:0, vehicle_tyre_state_fr:0, stage_shakedown:false, game_mode:0, vehicle_id:0, vehicle_class_id:0, vehicle_manufacturer_id:0, location_id:0, route_id:0, vehicle_cluster_abs:false
}

func ExamplePacket_MarshalBinary() {
	pkt := packet.New()
	pkt.Set("packet_4cc", "ABCD")
	pkt.Set("packet_uid", uint64(123456))
	fmt.Println(pkt.MarshalBinary())
	// Output:
	// [65 66 67 68 64 226 1 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0] <nil>
}

func ExamplePacket_UnmarshalBinary() {
	pkt := packet.New()
	b := []byte{
		65, 66, 67, 68, 64, 226, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	}
	if err := pkt.UnmarshalBinary(b); err != nil {
		panic(err)
	}
	fmt.Println(pkt)
	// Output:
	// packet_4cc:ABCD, packet_uid:123456, shiftlights_fraction:0.000000, shiftlights_rpm_start:0.000000, shiftlights_rpm_end:0.000000, shiftlights_rpm_valid:false, vehicle_gear_index:0, vehicle_gear_index_neutral:0, vehicle_gear_index_reverse:0, vehicle_gear_maximum:0, vehicle_speed:0.000000, vehicle_transmission_speed:0.000000, vehicle_position_x:0.000000, vehicle_position_y:0.000000, vehicle_position_z:0.000000, vehicle_velocity_x:0.000000, vehicle_velocity_y:0.000000, vehicle_velocity_z:0.000000, vehicle_acceleration_x:0.000000, vehicle_acceleration_y:0.000000, vehicle_acceleration_z:0.000000, vehicle_left_direction_x:0.000000, vehicle_left_direction_y:0.000000, vehicle_left_direction_z:0.000000, vehicle_forward_direction_x:0.000000, vehicle_forward_direction_y:0.000000, vehicle_forward_direction_z:0.000000, vehicle_up_direction_x:0.000000, vehicle_up_direction_y:0.000000, vehicle_up_direction_z:0.000000, vehicle_hub_position_bl:0.000000, vehicle_hub_position_br:0.000000, vehicle_hub_position_fl:0.000000, vehicle_hub_position_fr:0.000000, vehicle_hub_velocity_bl:0.000000, vehicle_hub_velocity_br:0.000000, vehicle_hub_velocity_fl:0.000000, vehicle_hub_velocity_fr:0.000000, vehicle_cp_forward_speed_bl:0.000000, vehicle_cp_forward_speed_br:0.000000, vehicle_cp_forward_speed_fl:0.000000, vehicle_cp_forward_speed_fr:0.000000, vehicle_brake_temperature_bl:0.000000, vehicle_brake_temperature_br:0.000000, vehicle_brake_temperature_fl:0.000000, vehicle_brake_temperature_fr:0.000000, vehicle_engine_rpm_max:0.000000, vehicle_engine_rpm_idle:0.000000, vehicle_engine_rpm_current:0.000000, vehicle_throttle:0.000000, vehicle_brake:0.000000, vehicle_clutch:0.000000, vehicle_steering:0.000000, vehicle_handbrake:0.000000, game_total_time:0.000000, game_delta_time:0.000000, game_frame_count:0, stage_current_time:0.000000, stage_previous_split_time:0.000000, stage_result_time:0.000000, stage_result_time_penalty:0.000000, stage_result_status:0, stage_current_distance:0.000000, stage_length:0.000000, stage_progress:0.000000, vehicle_tyre_state_bl:0, vehicle_tyre_state_br:0, vehicle_tyre_state_fl:0, vehicle_tyre_state_fr:0, stage_shakedown:false, game_mode:0, vehicle_id:0, vehicle_class_id:0, vehicle_manufacturer_id:0, location_id:0, route_id:0, vehicle_cluster_abs:false
}
