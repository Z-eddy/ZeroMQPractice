package main

import (
	"MyTest/TrCtlpb"
	"fmt"
	"github.com/golang/protobuf/proto"
	zmq "github.com/pebbe/zmq4"
	"os"
)

func main() {
	zctx, _ := zmq.NewContext()

	// Socket to talk to server
	fmt.Printf("Connecting to the server...\n")
	s1, _ := zctx.NewSocket(zmq.SUB)
	defer s1.Close()
	err := s1.Connect("tcp://localhost:9526")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	} else {
		fmt.Println("success connected")
	}

	//s1.SetSubscribe("")
	s1.SetSubscribe("TrDevMeasData")
	//s1.SetSubscribe("LogMsg")

	for {
		msg, _ := s1.RecvMessage(0)
		theData := TrCtlpb.TrDevMeasData{}
		proto.Unmarshal([]byte(msg[1]), &theData)
		fmt.Println(theData.GetBarcode(), theData.GetMeasType(), theData.GetDatas())
	}
}
