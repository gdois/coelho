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

var messages [100]string

func generateMessages() {
	for i := 0; i < len(messages); i++ {
		messages[i] = fmt.Sprintf("teste%d", i+1)
	}
}

func producer() {
	conn, err := net.Dial("tcp", ServerAddress+":"+ServerPort)
	if err != nil {
		fmt.Println("Connection failed")
		return
	}
	defer conn.Close()

	for _, message := range messages {
		jsonMessage, err := json.Marshal(message)
		if err != nil {
			fmt.Println("Error encoding message:", err)
			continue
		}

		jsonMessage = append(jsonMessage, '\n')

		_, err = conn.Write(jsonMessage)
		if err != nil {
			fmt.Println("Error sending message:", err)
			continue
		}

		fmt.Printf("Sent message: %s\n", jsonMessage)
	}
}

func main() {
	generateMessages()
	producer()
}
