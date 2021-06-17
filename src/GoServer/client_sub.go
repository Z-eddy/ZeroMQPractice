package main

import (
	"MyTest/BaseDatapb"
	"fmt"
	"github.com/golang/protobuf/proto"
	zmq "github.com/pebbe/zmq4"
	"os"
)

func main() {
	zctx, _ := zmq.NewContext()

	// Socket to talk to server
	fmt.Printf("Connecting to the server...\n")
	s, _ := zctx.NewSocket(zmq.SUB)
	defer s.Close()
	err:=s.Connect("tcp://localhost:9526")
	if err!=nil{
		fmt.Println(err)
		os.Exit(0)
	}else{
		fmt.Println("success connected")
	}

	s.SetSubscribe("")
	//s.SetSubscribe("LogMsg")

	for {
		msg, _ := s.RecvMessage(0)
		theData := BaseDatapb.LogData{}
		proto.Unmarshal([]byte(msg[1]), &theData)
		fmt.Println(theData.GetType(), string(theData.GetMsg()))
	}
}
