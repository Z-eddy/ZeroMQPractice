package main

import (
	"MyTest/BaseDatapb"
	"fmt"
	"github.com/golang/protobuf/proto"
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
	s.SetSubscribe("LogMsg")

	for {
		msg, _ := s.RecvMessage(0)
		theData := BaseDatapb.LogData{}
		proto.Unmarshal([]byte(msg[1]), &theData)
		fmt.Println(theData.GetType(), string(theData.GetMsg()))
	}
}
