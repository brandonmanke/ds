package main

import (
	"os"
	"fmt"
	"time"
	"sync"
	"errors"
	"strings"
	"github.com/go-redis/redis"
)	

func lookupHostnameEnv() (string, error) {
	hostname, ok := os.LookupEnv("HOST")
	if ok != true {
		return "", errors.New("HOST environment variable not set.")
	}
	return hostname, nil
}

func redisConnect(options *redis.Options) {
	redis := redis.NewClient(options)
	pong, err := redis.Ping().Result()
	if err != nil {
		fmt.Println("ERROR:")
		fmt.Println(err)
	}
	fmt.Printf("pong: %s\n", pong)
}

func main() {
	var sb strings.Builder
	hostname, err := lookupHostnameEnv()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
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

	var wg sync.WaitGroup
	wg.Add(1)

	time.AfterFunc(2 * time.Second, func() {
		redisConnect(&options)
		wg.Done()
	})

	wg.Wait() // blocks until wg counter is 0
}
