package main

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"os"
	"strings"
	"time"
)

type Subscriber struct {
	redis *redis.Client
	pubsub *redis.PubSub
	chanName string
	ch       <-chan *redis.Message
}

// Lookup environment variable HOST for hostname parameter
func lookupHostnameEnv() (string, error) {
	hostname, ok := os.LookupEnv("HOST")
	if ok != true {
		return "", errors.New("HOST environment variable not set.")
	}
	return hostname, nil
}

func CreateSubscriber(chanName string) (*Subscriber, error) {
	var sb strings.Builder
	hostname, err := lookupHostnameEnv()
	if err != nil {
		fmt.Println(err)
		return &Subscriber{}, err
	}
	sb.WriteString(hostname)
	sb.WriteString(":6379")
	hostname = sb.String()
	var options = redis.Options{
		Addr:     hostname,
		Password: "",
		DB:       0,
		// OnConnect: func(*Conn) error
	}
	redis := redis.NewClient(&options)
	time.Sleep(time.Second)
	pubsub := redis.Subscribe(chanName)
	ch := pubsub.Channel()
	sub := Subscriber{
		redis: redis,
		pubsub: pubsub,
		chanName: chanName,
		ch: ch,
	}
	return &sub, nil
}

func (s *Subscriber) GetChannel() <-chan *redis.Message {
	return s.ch
}
