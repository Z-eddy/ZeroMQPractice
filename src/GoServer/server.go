package main

import (
	"MyTest/EmployeeTestpb"
	"fmt"
	"github.com/golang/protobuf/proto"
	"time"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	zctx, _ := zmq.NewContext()

	s, _ := zctx.NewSocket(zmq.REP)
	s.Bind("tcp://*:5555")

	for {
		// Wait for next request from client
		msg, _ := s.Recv(0)
		//log.Printf("Recieved %s\n", msg)

		// Do some 'work'
		time.Sleep(time.Second * 1)

		theData := &EmployeeTestpb.Employee{}
		proto.Unmarshal([]byte(msg), theData)

		fmt.Println(theData.GetId(), theData.GetAge(), theData.GetName())

		// Send reply back to client
		s.Send("World", 0)
	}
}
