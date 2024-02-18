package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func consumer() {
	address := os.Getenv("ServerAddress") + ":" + os.Getenv("ServerPort")
	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer listener.Close()

	fmt.Printf("Server is listening in %s\n", address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			if err == io.EOF {
				fmt.Println("Connection closed")
				return
			}
			fmt.Println("Error:", err)
			continue
		}

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	queue := NewQueue("queue.txt")
	defer conn.Close()

	buffer := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Printf("Received: %s\n", buffer[:n])

		if string(buffer[:n]) == "read" {
			message := queue.Peek()
			fmt.Println(message)
			queue.Pop()
		} else {
			queue.Push(string(buffer[:n]))
		}

		fmt.Printf("Queue size %d\n", queue.Length())
	}
}
