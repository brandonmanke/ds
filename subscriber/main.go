package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)

const port = ":8080"
var addr = flag.String("addr", port, "http service address")
var upgrader = websocket.Upgrader{}

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

func handleSocketConn(wr http.ResponseWriter, req *http.Request) {
	// Never do this in production
	upgrader.CheckOrigin = func(r *http.Request) bool { 
		return true 
	}
	socket, err := upgrader.Upgrade(wr, req, nil)
	if err != nil {
		log.Println(err)
	}
	defer socket.Close()

	for {
		mt, message, err := socket.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = socket.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func main() {
	go testSub()

	log.SetFlags(0)
	log.Println("Serving at localhost:8080...")
	http.HandleFunc("/ws", handleSocketConn)
	log.Fatal(http.ListenAndServe(*addr, nil))
}