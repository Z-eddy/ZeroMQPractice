package main

import (
	"fmt"
	"sync"
)

var (
	myWait = sync.WaitGroup{}
	mutex  sync.Mutex
	count  = 10
)

func foo() {
	//锁住
	mutex.Lock()

	//执行内容
	for i := 0; i != count; i++ {
		fmt.Println(i)
		myWait.Done()
	}

	//释放
	mutex.Unlock()
}

func main() {
	myWait.Add(count)
	go foo()

	myWait.Add(count)
	go foo()

	myWait.Wait()
}
