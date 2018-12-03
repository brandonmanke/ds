/*Super basic architecture overview
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

// Reads config.json file at file path and returns it as a map
func readConfig(filePath string) (map[string]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var config map[string]string
	json.Unmarshal(bytes, &config)
	return config, nil
}

func main() {
	pub, err := CreatePublisher("test.channel")
	if err != nil {
		panic(err)
	}
	defer pub.Close()

	config, err := readConfig("./config.json")
	if err != nil {
		panic(err)
	}

	weatherAPI := CreateWeatherAPI(config["weather-api-key"])
	res, err := weatherAPI.GetForecast()
	if err != nil {
		panic(err)
	}

	serializedWeatherRes, err := SerializeJSON(res)
	if err != nil {
		panic(err)
	}

	newsAPI := CreateNYTimesAPI(config["news-api-key"])
	res, err = newsAPI.GetHome()
	if err != nil {
		panic(err)
	}

	serializedNewsRes, err := SerializeJSON(res)
	if err != nil {
		panic(err)
	}

	// TODO remove wait groups and use channels instead
	// to know when threads finish publishing
	// We need to refactor this anyways since we are going
	// to be polling a bunch of different APIs
	//ch := make

	var wg sync.WaitGroup
	wg.Add(2)

	// publish our weather and news data
	pub.PublishMessage(serializedWeatherRes)
	pub.PublishMessage(serializedNewsRes)

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
