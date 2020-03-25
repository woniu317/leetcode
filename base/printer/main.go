package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	//fmt.Println(runtime.NumCPU())
	numCpu := runtime.NumCPU()
	ch := make(chan int, numCpu)
	step := 1000000
	start := time.Now().UnixNano()
	for i := 0; i < numCpu; i++ {
		go sum(i*step, (i+1)*step-1, ch)
	}
	sum := 0
	for i := 0; i < numCpu; i++ {
		c := <-ch
		sum += c
	}
	end := time.Now().UnixNano()
	fmt.Println(sum, ",cost:", end-start)
	//time.NewTicker(1 * time.Second)
	//for range time.NewTicker(1 * time.Second).C {
	//}

	//task := make(chan int)
	//go PrinterM(task)
	//for i := 1; i <= 10; i++ {
	//	task <- i
	//}
	//task <- 0
	//<-task
}

func sum(i, j int, ch chan int) {
	sum := 0
	for i < j {
		sum += i
		i++
	}
	ch <- sum
}

func PrinterM(task chan int) {
	for ch := range task {
		if ch == 0 {
			break
		}
		fmt.Println(ch)
	}
	task <- 0
}
