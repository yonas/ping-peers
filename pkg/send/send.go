package send

import (
 "net"
 "fmt"
 "time"

 log "broadcast/pkg/logging"
 "github.com/vmihailenco/msgpack/v5"
)

type Msg struct {
    Tag string
    Text string
}

func Broadcast(tag string) {
  pc, err := net.ListenPacket("udp4", ":8829")
  if err != nil {
    panic(err)
  }
  defer pc.Close()

  addr, err := net.ResolveUDPAddr("udp4", "192.168.1.255:8829")
  if err != nil {
    panic(err)
  }

  pc2, err2 := net.ListenPacket("udp4", ":8828")
  if err2 != nil {
    panic(err2)
  }
  defer pc.Close()

  done := make(chan bool)
  timeout := time.After(3 * time.Second)

  go sendBroadcast(tag, addr, pc, pc2, done)

  for {
    select {
      case <-done:
        done = make(chan bool)
        go sendBroadcast(tag, addr, pc, pc2, done)

      case <-timeout:
        log.Debug("Timeout")
        return
    }

    time.Sleep(time.Second * 1)
  }
}

func sendBroadcast(tag string, addr net.Addr, pc net.PacketConn, pc2 net.PacketConn, done chan bool) { 
  b, err := msgpack.Marshal(&Msg{Tag: tag, Text: "hello world"})
  if err != nil {
      panic(err)
  }

  _,err = pc.WriteTo(b, addr)
  if err != nil {
    panic(err)
  }

  buf := make([]byte, 1024)
  var addr2 net.Addr

  _, addr2, err = pc2.ReadFrom(buf)
  if err != nil {
    panic(err)
  }

  var msg Msg
  err = msgpack.Unmarshal(buf, &msg)
  if err != nil {
    panic(err)
  }

  fmt.Println(fmt.Sprintf("%s replied: %s #%s", addr2, msg.Text, msg.Tag))
  done <- true
}
