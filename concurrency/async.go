package main

import (
	"fmt"
	"time"
	"sync"

)

type Async struct {}

func (this *Async) Waterfall(tasks ...interface{}) {
	var wg sync.WaitGroup
	for i := 0; i < len(tasks); i++ {
		wg.Add(1)
		go tasks[i](wg)
	}
	wg.Wait()
	fmt.Println("hi")
}

func main() {
	var a *Async
	a.Waterfall(func (wg sync.WaitGroup) {
		time.Sleep(time.Second)
		fmt.Println("1")
		wg.Done()
	})
	
	fmt.Println("Hello world")
}
