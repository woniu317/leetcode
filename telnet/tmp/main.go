package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

const ExitCode = 1

func main() {
	address := "127.0.0.1:7001"
	exitChan := make(chan int)
	go server(address, exitChan)
	code := <-exitChan
	os.Exit(code)
}

func server(address string, exitCode chan int) {
	l, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println(err.Error())
		exitCode <- ExitCode
		return
	}
	fmt.Println("listen: " + address)
	defer l.Close()

	for {
		if conn, err := l.Accept(); err != nil {
			fmt.Println(err.Error())
			continue
		} else {
			fmt.Println("new session")
			go handleSession(conn, exitCode)
		}

	}
}

func handleSession(conn net.Conn, exitCode chan int) {
	fmt.Println("handle session started:")
	reader := bufio.NewReader(conn)
	for {
		if str, err := reader.ReadString('\n'); err != nil {
			fmt.Println("Session closed")
			conn.Close()
			break
		} else {
			str = strings.TrimSpace(str)
			if !processTelnetCommand(str, exitCode) {
				conn.Close()
				break
			}
			conn.Write([]byte(str + "\r\n"))
		}
	}
}

func processTelnetCommand(command string, exitCode chan int) bool {
	if strings.HasPrefix(command, "@close") {
		fmt.Println("Session close")
		return false
	} else if strings.HasPrefix(command, "@shutdown") {
		fmt.Println("Server shutdown")
		exitCode <- ExitCode
		return false
	}
	fmt.Println(command)
	return true
}
