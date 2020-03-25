package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	l, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println(err)
		return
	}
	go broadcaster()
	for {
		if conn, err := l.Accept(); err == nil {
			go handleConn(conn)
		} else {

		}
	}
}

type client chan<- string

var entering = make(chan client)
var leaving = make(chan client)
var messages = make(chan string)

func handleConn(conn net.Conn) {
	fmt.Println("Session started")
	ch := make(chan string)
	go clientWriter(conn, ch)
	who := conn.RemoteAddr().String()
	ch <- "欢迎" + who
	messages <- who + "上线"
	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ":" + input.Text()
	}
	leaving <- ch
	messages <- who + "下线"
	conn.Close()
}

func clientWriter(conn net.Conn, ch chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

func broadcaster() {
	clients := map[client]bool{}
	for {
		select {
		case msg := <-messages:
			fmt.Println("broad started")
			for cli := range clients {
				cli <- msg
			}
		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}
