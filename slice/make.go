package main

import "fmt"

func main() {
	res := make(chan int, 2)
	fmt.Println(res)
}
