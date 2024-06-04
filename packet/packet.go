package packet

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/sys/windows"
)

type ChannelTable map[string]*Channel

type KeyValue struct {
	Key   string
	Value any
}

type Packet struct{}

var (
	WrcRoot      string
	IdDicts      = &IDs{}
	ChannelDicts = ChannelTable{}
	packetConfig = &PacketsDef{}
	values       = []*KeyValue{}
	keys         = map[string]*KeyValue{}
	packetSize   = -1
)

func init() {
	doc, err := windows.KnownFolderPath(windows.FOLDERID_Documents, 0)
	if err != nil {
		log.Fatal(err)
	}
	WrcRoot = os.ExpandEnv(filepath.Join(doc, "My Games", "WRC", "telemetry"))
	if _, err := os.Stat(WrcRoot); err != nil {
		log.Fatal(err)
	}
	ib, err := ReadFileUTF16(filepath.Join(WrcRoot, "readme", "ids.json"))
	if err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(ib, &IdDicts); err != nil {
		log.Fatal(err)
	}
	cb, err := os.ReadFile(filepath.Join(WrcRoot, "readme", "channels.json"))
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
	pb, err := os.ReadFile(filepath.Join(WrcRoot, "readme", "udp", "wrc.json"))
	if err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(pb, &packetConfig); err != nil {
		log.Fatal(err)
	}
	template := packetConfig.Packets[0].Channels
	sz := 0
	for _, key := range template {
		channel, ok := ChannelDicts[key]
		if !ok {
			log.Fatal(fmt.Errorf("channel %s not found", key))
		}
		var value any
		switch channel.Type {
		default:
			log.Fatal(fmt.Errorf("type %s not found", channel.Type))
		case "boolean":
			value = false
			sz += 1
		case "float32":
			value = float32(0.0)
			sz += 4
		case "float64":
			value = 0.0
			sz += 8
		case "fourcc":
			value = "____"
			sz += 4
		case "uint8":
			value = uint8(0)
			sz += 1
		case "uint16":
			value = uint16(0)
			sz += 2
		case "uint64":
			value = uint64(0)
			sz += 8
		}
		kv := &KeyValue{
			Key:   key,
			Value: value,
		}
		values = append(values, kv)
		keys[key] = kv
	}
	packetSize = sz
}

func New() *Packet {
	return &Packet{}
}

func (p *Packet) String() string {
	res := strings.Builder{}
	first := true
	for _, v := range values {
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
			res.WriteString(fmt.Sprintf("%s:%t", v.Key, v.Value))
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
	for _, v := range values {
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
	if len(res) != packetSize {
		return nil, fmt.Errorf("invalid packet size %d", len(res))
	}
	return res, nil
}

func (p *Packet) UnmarshalBinary(b []byte) error {
	if len(b) != packetSize {
		return fmt.Errorf("invalid packet size %d", len(b))
	}
	reader := bytes.NewReader(b)
	buf := make([]byte, 8)
	for _, v := range values {
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
	v, ok := keys[key]
	if !ok {
		return nil, fmt.Errorf("key %s not found", key)
	}
	return v.Value, nil
}

func (p *Packet) Set(key string, value any) error {
	v, ok := keys[key]
	if !ok {
		return fmt.Errorf("key %s not found", key)
	}
	v.Value = value
	return nil
}
