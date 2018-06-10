package main

import (
	"fmt"
)

func main() {
	go func () {}()
	go fmt.Println("1")
	go func () {
		fmt.Println("2")
	}()
	fmt.Println("3")
}
