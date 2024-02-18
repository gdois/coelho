package main

import (
	"fmt"
	"os"
)

type Queue struct {
	messages []string
	filename string
}

func NewQueue(filename string) *Queue {
	q := &Queue{
		filename: filename,
	}
	q.loadFromFile()
	return q
}

func (q *Queue) Push(value string) {
	q.messages = append(q.messages, value)
	q.saveToFile()
}

func (q *Queue) Pop() string {
	if len(q.messages) == 0 {
		return ""
	}
	value := q.messages[0]
	q.messages = q.messages[1:]
	q.saveToFile()
	return value
}

func (q *Queue) Length() int {
	return len(q.messages)
}

func (q *Queue) loadFromFile() {
	file, err := os.Open(q.filename)
	defer file.Close()
	if err != nil {
		fmt.Println("Error loading from file:", err)
		return
	}
	var message string
	for {
		_, err := fmt.Fscanln(file, &message)
		if err != nil {
			break
		}
		q.messages = append(q.messages, message)
	}
}

func (q *Queue) Peek() string {
	if len(q.messages) == 0 {
		return ""
	}
	return q.messages[0]
}

func (q *Queue) saveToFile() {
	file, err := os.Create(q.filename)
	defer file.Close()
	if err != nil {
		fmt.Println("Error saving to file:", err)
		return
	}
	for _, message := range q.messages {
		fmt.Fprintln(file, message)
	}
}
