package main

import (
	"fmt"
)

func main() {
	fmt.Println("Subscriber Test!")
	sub, err := CreateSubscriber("test.channel")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	ch := sub.GetChannel()
	for msg := range ch {
		fmt.Println(msg.Channel, msg.Payload)
	}
}