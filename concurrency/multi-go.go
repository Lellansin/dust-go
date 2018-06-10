package main

import (
	"fmt"
	"time"
	// "sync"
)

// func run(wg *sync.WaitGroup) {
func run() {
	c := make (chan int)
	go func () {
		time.Sleep(time.Second)
		fmt.Println("111", 111)
		// wg.Done()
		c <- 0
	}()
	<- c
}

func main() {
	// var wg sync.WaitGroup
	// start := time.Now()
	for i := 0; i < 5; i++ {
		// wg.Add(1)
		// go run(&wg)
		go run()
	}
	// wg.Wait()
	// delta := time.Now().Sub(start)
	for {
		time.Sleep(time.Second)
		// fmt.Println("Hello world", delta)
		// fmt.Println("Hello world", delta)
	}
}