package main

import (
	"fmt"
	"log"
	"net/http"
	socketio "github.com/googollee/go-socket.io"
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

	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	server.On("connection", func(socket socketio.Socket) {
		log.Println("connected!")
		socket.On("disconnection", func() {
			log.Println("disconnected..")
		})
	})

	server.On("error", func(s socketio.Socket, err error) {
		log.Println("error:", err)
	})

	http.Handle("/ws", server)
	log.Println("Serving at localhost:8080...")
	http.ListenAndServe(":8080", nil)
}