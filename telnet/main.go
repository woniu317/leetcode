package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func processTelnetComand(command string, exitChan chan int) bool {
	switch command {
	case "@close":
		fmt.Println("Session close")
		return false
	case "@shutdown":
		fmt.Println("Server shutdown")
		exitChan <- 0
		return false
	default:
		return true
	}
}

func handleSession(conn net.Conn, exitChan chan int) {
	fmt.Println("Session started")
	reader := bufio.NewReader(conn)
	for {
		str, err := reader.ReadString('\n')
		if err == nil {
			str = strings.TrimSpace(str)
			if !processTelnetComand(str, exitChan) {
				conn.Close()
				break
			}
			conn.Write([]byte(str + "\r\n"))
		} else {
			fmt.Println("Session closed")
			conn.Close()
			break
		}
	}
}

func server(address string, exitChan chan int) {
	l, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Listen ", address)
	defer l.Close()
	for {
		if conn, err := l.Accept(); err == nil {
			go handleSession(conn, exitChan)
		} else {
			fmt.Println(err)
		}
	}
}

func main() {
	address := "127.0.0.1:7001"
	exitChan := make(chan int)
	go server(address, exitChan)
	code := <-exitChan
	os.Exit(code)
}
