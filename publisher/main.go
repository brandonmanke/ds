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
	"time"
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

func setInterval(d time.Duration, f func()) {
	for range time.Tick(d) {
		f()
	}
}

func publishWeather(channel string, config map[string]string) error {
	pub, err := CreatePublisher(channel)
	weatherAPI := CreateWeatherAPI(config["weather-api-key"])
	res, err := weatherAPI.GetForecast()
	if err != nil {
		return err
	}

	serializedWeatherRes, err := SerializeJSON(res)
	if err != nil {
		return err
	}
	pub.PublishMessage(serializedWeatherRes)
	return nil
}

func publishNews(channel string, config map[string]string) error {
	pub, err := CreatePublisher(channel)
	newsAPI := CreateNYTimesAPI(config["news-api-key"])
	res, err := newsAPI.GetHome()
	if err != nil {
		return err
	}

	serializedNewsRes, err := SerializeJSON(res)
	if err != nil {
		return err
	}

	pub.PublishMessage(serializedNewsRes)
	return nil
}

func publishTest(channel, msg string) error {
	pub, err := CreatePublisher(channel)
	if err != nil {
		return err
	}
	defer pub.Close()
	pub.PublishMessage(msg)
	return nil
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

	// TODO remove wait groups and use channels instead
	// easier to know when threads finish publishing
	//ch := make

	var wg sync.WaitGroup
	wg.Add(5)

	go func() {
		testMessages(pub)
		wg.Done()
	}()

	go func() {
		testMessages(pub)
		wg.Done()
	}()

	go func() {
		setInterval(15*time.Second, func() {
			fmt.Println("publishing to test.channel - every 15 seconds")
			publishTest("test.channel", "every 15 seconds")
		})
		wg.Done()
	}()

	go func() {
		setInterval(25*time.Second, func() {
			fmt.Println("publishing to weather channel - every 25 seconds")
			publishWeather("weather", config)
		})
		wg.Done()
	}()

	go func() {
		setInterval(30*time.Second, func() {
			fmt.Println("publishing to news channel - every 30 seconds")
			publishNews("news", config)
		})
		//wg.Done()
	}()

	wg.Wait()
}
