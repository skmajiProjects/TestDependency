package main

import (
	"log"

	"github.com/skmajiProjects/TestDependency-Message/messageGo/messageProto"
)

func main() {
	msg := &messageProto.TestMessage{
		MessageId:   "id123",
		MessageName: "message123",
	}
	log.Println("message Data=", msg)
}
