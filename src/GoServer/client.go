package main

import (
	"fmt"
	zmq "github.com/pebbe/zmq4"
)

func main() {
	zctx, _ := zmq.NewContext()

	// Socket to talk to server
	fmt.Printf("Connecting to the server...\n")
	s, _ := zctx.NewSocket(zmq.SUB)
	defer s.Close()
	s.Connect("tcp://localhost:9526")

	s.SetSubscribe("AA")
	s.SetSubscribe("BB")

	for {
		msg, _ := s.RecvMessage(0)
		for idx, str := range msg {
			fmt.Println(idx, str)
		}
	}
}
