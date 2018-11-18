package main

import ( 
	"fmt"
	"sync"
)

func testMessages(pub *Publisher) {
	err := pub.PublishMessages("Hello", "World", "!!!!", "more messages", "omg")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	//time.Sleep(500 * time.Millisecond)
}

func main() {
	pub, err := CreatePublisher("test.channel")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	// TODO remove wait groups and use channels instead 
	// to know when threads finish publishing
	//ch := make

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		testMessages(pub)
	}()

	go func() {
		testMessages(pub)
	}()
	
	//wg.Done()
	wg.Wait()
}