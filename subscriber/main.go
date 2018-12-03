package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"

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

func subscribe(channelName string, mt int, socket *websocket.Conn) {
	fmt.Printf("Subscribing to %s.\n", channelName)
	sub, err := CreateSubscriber(channelName)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	ch := sub.GetChannel()
	// print all messages delivered to channel (blocking)
	for msg := range ch {
		var sb strings.Builder
		sb.WriteString(msg.Channel)
		sb.WriteString(": ")
		sb.WriteString(msg.Payload)
		socket.WriteMessage(mt, []byte(sb.String()))
		fmt.Println("Emiting message:", msg.Channel, msg.Payload)
	}
}

func unsubscribe(channelName string, socket *websocket.Conn) error {
	fmt.Printf("Unsubscribing to %s.\n", channelName)
	redis, err := GetRedis()
	if err != nil {
		log.Println(err)
		panic(err)
	}
	// unsure is this actually does much
	// we need to somehow get a reference
	// to the redis instance and unsub
	err = redis.Subscribe(channelName).Unsubscribe(channelName)
	if err != nil {
		socket.WriteMessage(1, []byte(err.Error()))
		return err
	}
	return nil
}

func listenForMessages(socket *websocket.Conn) {
	// parse messages for subscribe:  channel
	// parse messages for unsubscribe: channel
	// subscribe and emit accordingly,
	// if client is subscribed to that channel
	for {
		// Possible MT values:
		//TextMessage = 1
		//BinaryMessage = 2
		//CloseMessage = 8
		//PingMessage = 9
		//PongMessage = 10

		mt, err := parseMessage(socket)
		if mt == 8 {
			log.Println("Close message received. Closing Socket.")
			break
		}
		if err != nil {
			log.Println("Parse message error:", err)
			socket.WriteMessage(mt, []byte(err.Error()))
			break
		}
	}
}

// Temp not sure if needed
func parseMessage(socket *websocket.Conn) (int, error) {
	mt, message, err := socket.ReadMessage()
	if err != nil {
		log.Println("read error:", err)
		return mt, err
	}
	str := string(message)
	arr := strings.Split(str, " ")
	if len(arr) != 2 {
		return mt, errors.New("Sentence must follow format: [subscribe/unsubscribe] channel")
	}
	switch arr[0] {
	case "subscribe":
		go subscribe(arr[1], mt, socket)
	case "unsubscribe":
		go unsubscribe(arr[1], socket)
	default:
		var sb strings.Builder
		sb.WriteString("Unable to understand message: ")
		sb.WriteString(arr[0])
		return mt, errors.New(sb.String())
	}
	return mt, nil
}

func handleSocketConn(wr http.ResponseWriter, req *http.Request) {
	// Never do this in production (allows CORS for dev)
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	socket, err := upgrader.Upgrade(wr, req, nil)
	if err != nil {
		log.Println(err)
	}
	defer socket.Close()

	go listenForMessages(socket)

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
