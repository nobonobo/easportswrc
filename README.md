# easportswrc

## environment var

`EASPORTSWRC_DOC_ROOT`:

  WRC document root folder path

  Default: `%USERPROFILE%\Documents\My Games\WRC`

## usage

```go
import "github.com/nobonobo/easportswrc/packet"

func main() {
  pkt := packet.New()
  if err := pkt.UnmarshalBinary([]byte{...}); err != nil {
    panic(err)
  }
  pkt.GameMode = 2
  b, err := pkt.MarshalBinary()
  if err := nil {
    panic(err)
  }
  fmt.Printf("%x", b)
}
```
