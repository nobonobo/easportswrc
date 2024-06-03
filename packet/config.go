package packet

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"os"
	"path/filepath"

	"golang.org/x/sys/windows"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

type ChannelTable map[string]*Channel

var (
	WrcRoot      string
	IdDicts      = &IDs{}
	ChannelDicts = ChannelTable{}
	packetConfig = &PacketsDef{}
)

func ReadFileUTF16(filename string) ([]byte, error) {

	// Read the file into a []byte:
	raw, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Make an tranformer that converts MS-Win default to UTF8:
	win16be := unicode.UTF16(unicode.BigEndian, unicode.IgnoreBOM)
	// Make a transformer that is like win16be, but abides by BOM:
	utf16bom := unicode.BOMOverride(win16be.NewDecoder())

	// Make a Reader that uses utf16bom:
	unicodeReader := transform.NewReader(bytes.NewReader(raw), utf16bom)

	// decode and print:
	decoded, err := io.ReadAll(unicodeReader)
	return decoded, err
}

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
}
