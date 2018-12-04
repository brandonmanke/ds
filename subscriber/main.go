package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

const port = ":8080"

var addr = flag.String("addr", port, "http service address")
var upgrader = websocket.Upgrader{}

// WebSocket wrapper around gorilla websocket library
type WebSocket struct {
	socket      *websocket.Conn
	subscribers map[string]*Subscriber
	sync.Mutex
}

// CreateWS Creates WebSocket type
func CreateWS(socket *websocket.Conn, subs map[string]*Subscriber) *WebSocket {
	return &WebSocket{
		socket:      socket,
		subscribers: subs,
	}
}

func testSub() {
	fmt.Println("Subscriber Test!")
	sub, err := CreateSubscriber("test.channel")
	if err != nil {
		log.Println(err)
		panic(err)
	}
	defer sub.CloseChan()

	ch := sub.GetChannel()
	// print all messages delivered to channel (blocking)
	for msg := range ch {
		fmt.Println(msg.Channel, msg.Payload)
	}
}

func createSubMap(channels ...string) map[string]*Subscriber {
	m := make(map[string]*Subscriber)
	for _, ch := range channels {
		sub, err := CreateSubscriber(ch)
		if err != nil {
			log.Println(err)
			panic(err)
		}
		m[ch] = sub
	}
	return m
}

func (ws *WebSocket) subscribe(channelName string, mt int) error {
	ws.Lock()
	defer ws.Unlock()
	fmt.Printf("Subscribing to %s.\n", channelName)
	ws.socket.WriteMessage(mt, []byte("Subbing: "+channelName))
	if ws.subscribers[channelName].subbed == false {
		ws.subscribers[channelName].subbed = true
	}

	err := ws.subscribers[channelName].pubsub.Subscribe(channelName)
	if err != nil {
		ws.subscribers[channelName].subbed = false
		fmt.Println("subscribe error:", err)
		return err
	}

	//ch := ws.subscribers[channelName].ch
	// print all messages delivered to channel (blocking)
	/*select {
	case msg := <-ch:
		// as soon as we receive a message we decrement our wg counter
		// so this function stops blocking and returns
		var sb strings.Builder
		sb.WriteString(msg.Channel)
		sb.WriteString(": ")
		sb.WriteString(msg.Payload)
		fmt.Println("Emiting message:", msg.Channel, msg.Payload)
		err := ws.socket.WriteMessage(mt, []byte(sb.String()))
		if err != nil {
			ws.subscribers[channelName].subbed = false
			fmt.Println("error writing message:", err)
			break
		}
		if ws.subscribers[channelName].subbed == false {
			break
		}
	default:
		break
	}*/
	return nil
}

func (ws *WebSocket) unsubscribe(channelName string, mt int) error {
	if channelName == "poll" {
		return nil
	}
	ws.Lock()
	defer ws.Unlock()
	fmt.Printf("Unsubscribing to %s.\n", channelName)
	ws.socket.WriteMessage(mt, []byte("Unsubbing: "+channelName))
	err := ws.subscribers[channelName].pubsub.Unsubscribe(channelName)
	if err != nil {
		return err
	}
	ws.subscribers[channelName].subbed = false
	return nil
}

func (ws *WebSocket) listenForMessages() {
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

		mt, err := ws.parseMessage() //ws.socket
		//fmt.Println(mt, err)
		//if mt == 8 {
		//	log.Println("Close message received. Closing Socket.")
		//	break
		//}
		if err != nil {
			ws.Lock()
			log.Println("Parse message error:", err)
			ws.socket.WriteMessage(mt, []byte(err.Error()))
			ws.Unlock()
			break
		}

		select {
		case testmsg := <-ws.subscribers["test.channel"].GetChannel():
			ws.Lock()
			fmt.Println("Sending test.message")
			ws.socket.WriteMessage(mt, []byte(testmsg.String()))
			ws.Unlock()
		case weathermsg := <-ws.subscribers["weather"].GetChannel():
			ws.Lock()
			fmt.Println("Sending weather msg")
			ws.socket.WriteMessage(mt, []byte(weathermsg.Channel+": "+weathermsg.Payload))
			ws.Unlock()
		case newsmsg := <-ws.subscribers["news"].GetChannel():
			ws.Lock()
			fmt.Println("Sending news msg")
			ws.socket.WriteMessage(mt, []byte(newsmsg.Channel+": "+newsmsg.Payload))
			ws.Unlock()
		default:
			time.Sleep(1 * time.Millisecond)
			break
		}
	}
}

// Temp not sure if needed
func (ws *WebSocket) parseMessage() (int, error) {
	mt, message, err := ws.socket.ReadMessage()
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
		go ws.subscribe(arr[1], mt)
	case "unsubscribe":
		go ws.unsubscribe(arr[1], mt)
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
	m := createSubMap("test.channel", "weather", "news", "poll")
	ws := CreateWS(socket, m)
	if err != nil {
		log.Println(err)
	}
	defer socket.Close()

	ws.listenForMessages()
}

func main() {
	//go testSub()

	log.SetFlags(0)
	log.Println("Serving at localhost:8080...")
	http.HandleFunc("/ws", handleSocketConn)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
