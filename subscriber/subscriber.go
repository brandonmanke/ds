package main

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"os"
	"strings"
	"time"
)

// Subscriber wrapper around go-redis redis client
type Subscriber struct {
	redis    *redis.Client
	pubsub   *redis.PubSub
	chanName string
	ch       <-chan *redis.Message
	subbed   bool
}

// Lookup environment variable HOST for hostname parameter
func lookupHostnameEnv() (string, error) {
	hostname, ok := os.LookupEnv("HOST")
	if ok != true {
		return "", errors.New("HOST environment variable not set")
	}
	return hostname, nil
}

// GetRedis return redis client instance
func GetRedis() (*redis.Client, error) {
	var sb strings.Builder
	hostname, err := lookupHostnameEnv()
	if err != nil {
		fmt.Println(err)
		return nil, err
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
	return redis, nil
}

// CreateSubscriber creates subscriber object and returns *Subscriber to it
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
		redis:    redis,
		pubsub:   pubsub,
		chanName: chanName,
		ch:       ch,
		subbed:   true,
	}
	return &sub, nil
}

// GetChannel returns channel type that current subscriber is using
func (s *Subscriber) GetChannel() <-chan *redis.Message {
	return s.ch
}

// CloseChan close pubsub chan
func (s *Subscriber) CloseChan() {
	s.pubsub.Close()
}
