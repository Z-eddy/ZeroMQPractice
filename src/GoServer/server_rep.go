package main

import (
	"MyTest/EmployeeTestpb"
	"MyTest/TrZeroMQMsgpb"
	"fmt"
	"github.com/golang/protobuf/proto"
	zmq "github.com/pebbe/zmq4"
)

func main() {
	zctx, _ := zmq.NewContext()

	s, _ := zctx.NewSocket(zmq.REP)
	s.Bind("tcp://*:9527")

	for {
		// Wait for next request from client
		msg, _ := s.Recv(0)
		//log.Printf("Recieved %s\n", msg)

		// Do some 'work'
		//time.Sleep(time.Second * 1)

		topData := &TrZeroMQMsgpb.TopData{}
		proto.Unmarshal([]byte(msg), topData)

		if topData.GetType() == "Employee" {
			theData := &EmployeeTestpb.Employee{}
			proto.Unmarshal(topData.GetRawData(), theData)

			fmt.Println(theData.GetId(), theData.GetAge(), theData.GetName())
		}

		// Send reply back to client
		s.Send("World", 0)
	}
}
