package packet

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
	"strings"
)

type KeyValue struct {
	Key   string
	Value any
}

type Packet struct {
	values []*KeyValue
	keys   map[string]*KeyValue
}

func Make() (*Packet, error) {
	template := packetConfig.Packets[0].Channels
	p := &Packet{
		values: []*KeyValue{},
		keys:   make(map[string]*KeyValue),
	}
	for _, key := range template {
		channel, ok := ChannelDicts[key]
		if !ok {
			return nil, fmt.Errorf("channel %s not found", key)
		}
		var value any
		switch channel.Type {
		default:
			return nil, fmt.Errorf("type %s not found", channel.Type)
		case "boolean":
			value = false
		case "float32":
			value = float32(0.0)
		case "float64":
			value = 0.0
		case "fourcc":
			value = "____"
		case "uint8":
			value = uint8(0)
		case "uint16":
			value = uint16(0)
		case "uint64":
			value = uint64(0)
		}
		kv := &KeyValue{
			Key:   key,
			Value: value,
		}
		p.values = append(p.values, kv)
		p.keys[key] = kv
	}
	return p, nil
}

func (p *Packet) String() string {
	res := strings.Builder{}
	first := true
	for _, v := range p.values {
		ch, ok := ChannelDicts[v.Key]
		if !ok {
			return fmt.Sprintf("unknown channel %s", v.Key)
		}
		if first {
			first = false
		} else {
			res.WriteString(", ")
		}
		switch ch.Type {
		default:
			return fmt.Sprintf("unknown type %s", ch.Type)
		case "boolean":
			res.WriteString(fmt.Sprintf("%s:", v.Key))
			if v.Value.(bool) {
				res.WriteString("1")
			} else {
				res.WriteString("0")
			}
		case "float32":
			res.WriteString(fmt.Sprintf("%s:%f", v.Key, v.Value.(float32)))
		case "float64":
			res.WriteString(fmt.Sprintf("%s:%f", v.Key, v.Value.(float64)))
		case "fourcc":
			res.WriteString(fmt.Sprintf("%s:%s", v.Key, v.Value.(string)))
		case "uint8":
			res.WriteString(fmt.Sprintf("%s:%d", v.Key, v.Value.(uint8)))
		case "uint16":
			res.WriteString(fmt.Sprintf("%s:%d", v.Key, v.Value.(uint16)))
		case "uint64":
			res.WriteString(fmt.Sprintf("%s:%d", v.Key, v.Value.(uint64)))
		}
	}
	return res.String()
}

func (p *Packet) MarshalBinary() ([]byte, error) {
	res := []byte{}
	for _, v := range p.values {
		ch, ok := ChannelDicts[v.Key]
		if !ok {
			return nil, fmt.Errorf("unknown channel %s", v.Key)
		}
		switch ch.Type {
		default:
			return nil, fmt.Errorf("unknown type %s", ch.Type)
		case "boolean":
			if v.Value.(bool) {
				res = append(res, 1)
			} else {
				res = append(res, 0)
			}
		case "float32":
			res = binary.LittleEndian.AppendUint32(
				res,
				math.Float32bits(v.Value.(float32)),
			)
		case "float64":
			res = binary.LittleEndian.AppendUint64(
				res,
				math.Float64bits(v.Value.(float64)),
			)
		case "fourcc":
			res = append(res, v.Value.(string)[0:4]...)
		case "uint8":
			res = append(res, v.Value.(uint8))
		case "uint16":
			res = binary.LittleEndian.AppendUint16(
				res,
				v.Value.(uint16),
			)
		case "uint64":
			res = binary.LittleEndian.AppendUint64(
				res,
				v.Value.(uint64),
			)
		}
	}
	return nil, nil
}

func (p *Packet) UnmarshalBinary(b []byte) error {
	reader := bytes.NewReader(b)
	buf := make([]byte, 8)
	for _, v := range p.values {
		ch, ok := ChannelDicts[v.Key]
		if !ok {
			return fmt.Errorf("unknown channel %s", v.Key)
		}
		switch ch.Type {
		default:
			return fmt.Errorf("unknown type %s", ch.Type)
		case "boolean":
			if _, err := reader.Read(buf[0:1]); err != nil {
				return err
			}
			v.Value = buf[0] != 0
		case "float32":
			val := uint32(0)
			if err := binary.Read(reader, binary.LittleEndian, &val); err != nil {
				return err
			}
			v.Value = math.Float32frombits(val)
		case "float64":
			val := uint64(0)
			if err := binary.Read(reader, binary.LittleEndian, &val); err != nil {
				return err
			}
			v.Value = math.Float64frombits(val)
		case "fourcc":
			if _, err := reader.Read(buf[0:4]); err != nil {
				return err
			}
			v.Value = string(buf[0:4])
		case "uint8":
			if _, err := reader.Read(buf[0:1]); err != nil {
				return err
			}
			v.Value = buf[0]
		case "uint16":
			val := uint16(0)
			if err := binary.Read(reader, binary.LittleEndian, &val); err != nil {
				return err
			}
			v.Value = val
		case "uint64":
			val := uint64(0)
			if err := binary.Read(reader, binary.LittleEndian, &val); err != nil {
				return err
			}
			v.Value = val
		}
	}
	return nil
}

func (p *Packet) Get(key string) (any, error) {
	v, ok := p.keys[key]
	if !ok {
		return nil, fmt.Errorf("key %s not found", key)
	}
	return v.Value, nil
}

func (p *Packet) Set(key string, value any) error {
	v, ok := p.keys[key]
	if !ok {
		return fmt.Errorf("key %s not found", key)
	}
	v.Value = value
	return nil
}
