package main

import (
	"fmt"
	"log"
	"net/http"
)

func testSub() {
	fmt.Println("Subscriber Test!")
	sub, err := CreateSubscriber("test.channel")
	if err != nil {
		log.Println(err)
		panic(err)
	}

	ch := sub.GetChannel()
	// print all messages delivered to channel (blocking)
	for msg := range ch {
		fmt.Println(msg.Channel, msg.Payload)
	}
}

func main() {

	go testSub()

	log.Println("Serving at localhost:8080...")
	http.ListenAndServe(":8080", nil)
}