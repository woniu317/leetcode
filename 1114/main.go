package main

import (
	"fmt"
	"time"
)

var second = make(chan bool, 1)
var third = make(chan bool, 1)

func main() {
	go firstP()
	go secondP()
	time.Sleep(3 * time.Second)
	second <- true
	time.Sleep(3 * time.Second)
	time.Sleep(3 * time.Second)
}

func firstP() {
	fmt.Println("first")
	second <- true
}

func secondP() {
	select {
	case <-second:
		fmt.Println("second")
	case <-time.After(1 * time.Second):
		fmt.Println("default...")
	}
}
