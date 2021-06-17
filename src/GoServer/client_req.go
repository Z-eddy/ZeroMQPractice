package main

import (
	"MyTest/BaseDatapb"
	"MyTest/EmployeeTestpb"
	"fmt"
	"github.com/golang/protobuf/proto"
	zmq "github.com/pebbe/zmq4"
	"time"
)

func main() {
	zctx, _ := zmq.NewContext()

	// Socket to talk to server
	fmt.Printf("Connecting to the server...\n")
	s, _ := zctx.NewSocket(zmq.REQ)
	s.Connect("tcp://localhost:9527")

	for {
		theData := EmployeeTestpb.Employee{
			Id:   12,
			Age:  78,
			Name: "WangYiYun",
		}
		rawData, _ := proto.Marshal(&theData)

		topData := BaseDatapb.TopData{
			Type:    "Employee",
			RawData: rawData,
		}

		sendData, _ := proto.Marshal(&topData)

		s.SendBytes(sendData, 0)
		msg, _ := s.Recv(0)
		fmt.Println(msg)

		time.Sleep(2000 * time.Millisecond)
	}
}
