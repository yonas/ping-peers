package sender

import (
 "net"
 "fmt"

 "github.com/vmihailenco/msgpack/v5"
)

var j string

func send(tag string) {
  pc, err := net.ListenPacket("udp4", ":8829")
  if err != nil {
    panic(err)
  }
  defer pc.Close()

  pc2,err2 := net.ListenPacket("udp4", ":8828")
  if err2 != nil {
    panic(err2)
  }
  defer pc2.Close()

  addr,err := net.ResolveUDPAddr("udp4", "192.168.1.255:8829")
  if err != nil {
    panic(err)
  }

  type Msg struct {
      Tag string
      Text string
  }

  b, err := msgpack.Marshal(&Msg{Tag: tag, Text: "hello world"})
  if err != nil {
      panic(err)
  }

  _,err = pc.WriteTo(b, addr)
  if err != nil {
    panic(err)
  }

  buf := make([]byte, 1024)
  _,addr2,err2 := pc2.ReadFrom(buf)
  if err2 != nil {
    panic(err2)
  }
  
  var msg2 Msg
  err = msgpack.Unmarshal(buf, &msg2)
  if err != nil {
      panic(err)
  }

  fmt.Println(fmt.Sprintf("%s replied: %s #%s", addr2, msg2.Text, msg2.Tag))
}
