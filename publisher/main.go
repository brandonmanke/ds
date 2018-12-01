/* Super basic architecture overview
 *       ----------------
 *       |  Publisher   |
 *       ----------------
 *              | (publish)
 *              v
 *          ---------
 *          | Redis |
 *          ---------
 *            |   ^
 * (Receive)  v   | (Post messages & Subscribe)
 *       ----------------
 *       |  Subscriber  |
 *       ----------------
 *               ^
 *               | (WebSockets)
 *               v
 *        --------------
 *        |   Client   |
 *        --------------
 */
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
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

	file, err := os.Open("./config.json")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	var config map[string]string
	json.Unmarshal(bytes, &config)
	fmt.Println(config)
	fmt.Println("config key")
	fmt.Println(config["weather-api-key"])

	weatherAPI := CreateWeatherAPI(config["weather-api-key"])
	res, err := weatherAPI.GetForecast()
	if err != nil {
		panic(err)
	}
	fmt.Println("HTTP RESPONSE:")
	fmt.Println(res)

	// TODO remove wait groups and use channels instead
	// to know when threads finish publishing
	// We need to refactor this anyways since we are going
	// to be polling a bunch of different APIs
	//ch := make

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		testMessages(pub)
		wg.Done()
	}()

	go func() {
		testMessages(pub)
		wg.Done()
	}()

	wg.Wait()
}
