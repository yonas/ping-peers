package agent

import (
  "fmt"
  "net"
  "strings"

  log "broadcast/pkg/logging"

  "github.com/vmihailenco/msgpack/v5"
)

func Listen(tags string, response string) {
  pc,err := net.ListenPacket("udp4", ":8829")
  if err != nil {
    panic(err)
  }
  defer pc.Close()

  pc2, err2 := net.ListenPacket("udp4", ":8828")
  if err2 != nil {
    panic(err2)
  }
  defer pc2.Close()

  type Msg struct {
      Tag string
      Text string
  }

  for {
    buf := make([]byte, 1024)
    _,addr,err := pc.ReadFrom(buf)
    if err != nil {
      panic(err)
    }

    var msg Msg
    err = msgpack.Unmarshal(buf, &msg)
    if err != nil {
        panic(err)
    }

    log.Debug(fmt.Sprintf("%s sent this: %s #%s", addr, msg.Text, msg.Tag))

    if !strings.Contains(tags, msg.Tag) {
      continue
    }

    // Response
    addr2,err2 := net.ResolveUDPAddr("udp4", "192.168.1.255:8828")
    if err2 != nil {
      panic(err2)
    }

    b, err := msgpack.Marshal(&Msg{Tag: tags, Text: response})
    if err != nil {
      panic(err)
    }

    _,err2 = pc.WriteTo(b, addr2)
    if err2 != nil {
      panic(err2)
    }
  }
}
