package main

import (
	"encoding/json"
	"fmt"
	"net"
)

const (
	ServerAddress = "localhost"
	ServerPort    = "8080"
)

var messages = [3]string{"teste", "teste1", "teste3"}

func producer() {
	conn, err := net.Dial("tcp", ServerAddress+":"+ServerPort)

	if err != nil {
		fmt.Println("Connection failed")
		return
	}

	for _, message := range messages {
		jsonMessage, err := json.Marshal(message)
		if err != nil {
			fmt.Println("Error encoding message:", err)
			continue
		}

		fmt.Printf("Sent message: %s\n", jsonMessage)
		if err != nil {
			fmt.Println("Error sending message:", err)
			continue
		}
	}

	defer conn.Close()
}

func main() {
	producer()
}
