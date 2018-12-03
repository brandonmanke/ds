package main

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"os"
	"strings"
	"time"
)

// Publisher wrapper around go-redis redis client
type Publisher struct {
	redis    *redis.Client
	pubsub   *redis.PubSub
	chanName string
}

// Lookup environment variable HOST for hostname parameter
func lookupHostnameEnv() (string, error) {
	hostname, ok := os.LookupEnv("HOST")
	if ok != true {
		return "", errors.New("HOST environment variable not set")
	}
	return hostname, nil
}

// CreatePublisher Creates and returns a *Publisher
func CreatePublisher(ch string) (*Publisher, error) {
	var sb strings.Builder
	hostname, err := lookupHostnameEnv()
	if err != nil {
		fmt.Println(err)
		return &Publisher{}, err
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
	time.Sleep(time.Second) // sleep one second to wait for redis service to start
	pubsub := redis.Subscribe(ch)
	p := Publisher{
		redis:    redis,
		pubsub:   pubsub,
		chanName: ch,
	}
	return &p, nil
}

//func _createPublisher(r *redis.Client, ch string) *Publisher {

//channel := pubsub.Channel() //pubsub.ch

//return &p
//}

// PublishMessage takes interface{} and publishes it to channel
func (p *Publisher) PublishMessage(msg interface{}) error {
	err := p.redis.Publish(p.chanName, msg).Err()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// PublishMessages publishes multiple strings to channel
func (p *Publisher) PublishMessages(msgs ...string) error {
	for _, m := range msgs {
		err := p.redis.Publish(p.chanName, m).Err()
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	return nil
}

// Close pubsub channel
func (p *Publisher) Close() {
	p.pubsub.Close()
}
