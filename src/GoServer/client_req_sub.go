package main

import (
	"MyTest/BaseDatapb"
	"MyTest/EmployeeTestpb"
	"fmt"
	"github.com/golang/protobuf/proto"
	zmq "github.com/pebbe/zmq4"
	"os"
	"sync"
	"time"
)

var (
	Zctx0 *zmq.Context
	//Zctx1 *zmq.Context
	S0   *zmq.Socket
	S1   *zmq.Socket
	Wait = sync.WaitGroup{}
)

func S0Foo() {
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

		S0.SendBytes(sendData, 0)
		msg, _ := S0.Recv(0)
		fmt.Println(msg)

		time.Sleep(2000 * time.Millisecond)
	}

	Wait.Done()
}

func S1Foo() {
	for {
		msg, _ := S1.RecvMessage(0)
		theData := BaseDatapb.LogData{}
		proto.Unmarshal([]byte(msg[1]), &theData)
		fmt.Println(theData.GetType(), string(theData.GetMsg()))
	}

	Wait.Done()
}

func init() {
	Wait.Add(2)
	Zctx0, _ = zmq.NewContext()

	//S0
	S0, _ = Zctx0.NewSocket(zmq.REQ)
	err := S0.Connect("tcp://localhost:9527")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	} else {
		fmt.Println("success connected")
	}

	//S1
	S1, _ = Zctx0.NewSocket(zmq.SUB)
	err = S1.Connect("tcp://localhost:9526")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	} else {
		fmt.Println("success connected")
	}

	S1.SetSubscribe("")
	//S1.SetSubscribe("LogMsg")
}

func main() {
	go S0Foo()
	go S1Foo()

	defer S0.Close()
	defer S1.Close()

	Wait.Wait()
}
